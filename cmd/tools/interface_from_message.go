package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"k8s.io/apimachinery/pkg/util/sets"
)

type GenerationOptions struct {
	PkgPath    string
	Roots      []reflect.Type
	OmitFields map[string]sets.String
	SliceKeys  map[string]sets.String
}

type AbstractionContext struct {
	GenerationOptions
	Imports  map[string]string       // package path to name imported as
	Abstract map[reflect.Type]string // concrete message type to interface name for it
}

func (o GenerationOptions) NewAbstractionContext() *AbstractionContext {
	ret := &AbstractionContext{
		GenerationOptions: o,
		Imports:           map[string]string{},
		Abstract:          map[reflect.Type]string{},
	}
	for _, t := range o.Roots {
		interfacesToWrite(t, o.PkgPath, ret.Imports, ret.Abstract)
	}
	return ret
}

func (c *AbstractionContext) Type(t reflect.Type) string {
	return c.stringifyType(t, false, true)
}

func (c *AbstractionContext) TypeForGetter(t reflect.Type) string {
	return c.stringifyType(t, true, true)
}

func (c *AbstractionContext) TypeOriginal(t reflect.Type) string {
	return c.stringifyType(t, false, false)
}

func (c *AbstractionContext) TypeForKey(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr:
		return c.Type(t.Elem())
	default:
		return c.Type(t)
	}
}

func (c *AbstractionContext) TypeForWrapper(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Struct:
		pkgName := c.findPkgName(t.PkgPath())
		prefix := strings.Title(pkgName)
		return prefix + t.Name()
	case reflect.Slice:
		pkgName := c.findPkgName(t.Elem().PkgPath())
		prefix := strings.Title(pkgName)
		return prefix + t.Elem().Name() + "Slice"
	}
	panic("Got to the end and can't get a typedef name")
	return ""
}

func (c *AbstractionContext) WriteImports() string {
	b := bytes.NewBufferString("")
	fmt.Fprintf(b, "import(\n")
	for path, name := range c.Imports {
		if strings.HasSuffix(path, name) {
			fmt.Fprintf(b, "\t\"%s\"\n", path)
		} else {
			fmt.Fprintf(b, "\t%s \"%s\"\n", name, path)
		}
	}
	fmt.Fprintf(b, ")\n")
	return b.String()
}

func interfacesToWrite(t reflect.Type, pkgPath string, imports map[string]string, tm map[reflect.Type]string) {
	if t.PkgPath() != pkgPath {
		return
	}
	tm[t] = t.Name()
	if t.Kind() != reflect.Struct {
		return
	}
	for _, f := range TopLevelFields(t, imports) {
		switch f.Type.Kind() {
		case reflect.Map:
			interfacesToWrite(f.Type.Key(), pkgPath, imports, tm)
			fallthrough
		case reflect.Ptr:
			fallthrough
		case reflect.Slice:
			if f.Type.Elem().PkgPath() == pkgPath {
				tm[f.Type] = f.Type.Elem().Name() + "Slice"
			}
			interfacesToWrite(f.Type.Elem(), pkgPath, imports, tm)
		case reflect.Interface:
			// pass
		case reflect.Struct:
			interfacesToWrite(f.Type, pkgPath, imports, tm)
		}
	}
}

func (c *AbstractionContext) MakeSliceInterface(t reflect.Type) *Interface {
	vars := c.getSliceVariables(t)
	elemType := c.Type(t.Elem())
	methods := []*Function{
		&Function{Name: "Iter", RetTypes: []string{"chan " + elemType}},
		&Function{Name: "Index", ArgNames: vars.KeyArgNames, ArgTypes: vars.KeyArgTypes, RetTypes: []string{"int"}},
		&Function{Name: "Get", ArgNames: []string{"i"}, ArgTypes: []string{"int"}, RetTypes: []string{vars.ElemType}},
		&Function{Name: "Find", ArgNames: vars.KeyArgNames, ArgTypes: vars.KeyArgTypes, RetTypes: []string{elemType, "bool"}},
		&Function{Name: "Remove", ArgNames: vars.KeyArgNames, ArgTypes: vars.KeyArgTypes},
		&Function{Name: "Upsert", ArgNames: vars.ArgNames, ArgTypes: vars.ArgTypes},
	}
	// MatchesF for each key field
	for i, f := range vars.KeyFields {
		methods = append(methods, &Function{
			Name:     "Matches" + f.Name,
			ArgNames: []string{f.Arg},
			ArgTypes: []string{vars.KeyArgTypes[i]},
			RetTypes: []string{"[]" + elemType},
		})
	}

	return &Interface{
		Name:    t.Elem().Name() + "Slice",
		Methods: methods,
	}
}

func TemplBody(templ *template.Template, ctx interface{}) []string {
	w := bytes.NewBufferString("")
	err := templ.Execute(w, ctx)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(w.String(), "\n")
	ret := []string{}
	for _, l := range lines {
		if strings.TrimSpace(l) != "" {
			ret = append(ret, l)
		}
	}
	return ret
}

type fieldArgMatch struct {
	Name     string
	Arg      string
	Type     string
	Optional bool
}

type sliceVariables struct {
	ElemType    string
	ElemWrapper string
	KeyFields   []fieldArgMatch
	AllFields   []fieldArgMatch
	ArgNames    []string
	ArgTypes    []string
	JoinedArgs  string
	KeyArgNames []string
	KeyArgTypes []string
	JoinedKeys  string
}

func (c *AbstractionContext) getSliceVariables(t reflect.Type) sliceVariables {
	ret := sliceVariables{}
	keys := c.SliceKeys[t.Elem().Name()]
	fields := TopLevelFields(t.Elem(), c.Imports)
	for _, f := range fields {
		n := strings.ToLower(f.Name)
		ret.ArgNames = append(ret.ArgNames, n)
		typ := c.Type(f.Type)
		ret.ArgTypes = append(ret.ArgTypes, typ)
		ret.AllFields = append(ret.AllFields, fieldArgMatch{f.Name, n, typ, f.Type.Kind() == reflect.Ptr})
		if len(keys) == 0 || keys.Has(f.Name) {
			kTyp := c.TypeForKey(f.Type)
			ret.KeyArgNames = append(ret.KeyArgNames, strings.ToLower(f.Name))
			ret.KeyArgTypes = append(ret.KeyArgTypes, c.TypeForKey(f.Type))
			ret.KeyFields = append(ret.KeyFields, fieldArgMatch{f.Name, n, kTyp, f.Type.Kind() == reflect.Ptr})
		}
	}
	ret.JoinedArgs = strings.Join(ret.ArgNames, ", ")
	ret.JoinedKeys = strings.Join(ret.KeyArgNames, ", ")
	ret.ElemType = c.Type(t.Elem())
	ret.ElemWrapper = c.TypeForWrapper(t.Elem())

	return ret
}

var indexTempl = template.Must(template.New("Index").Parse(`
for i, elt := range s.Elts {
    {{range .KeyFields}}
    {{if .Optional}}
    var v {{.Type}}
    if elt.{{.Name}} != nil { v = *elt.{{.Name}} }
    if v != {{.Arg}} { continue }
    {{else}}
    if elt.{{.Name}} != {{.Arg}} { continue }
    {{end}}
    {{end}}
    return i
}
return -1`))

var iterTempl = template.Must(template.New("Iter").Parse(`
ret := make(chan {{.ElemType}}, len(s.Elts))
for _, elt := range s.Elts {
    ret <- {{.ElemWrapper}}{&elt}
}
close(ret)
return ret`))

var getTempl = template.Must(template.New("Get").Parse(`
return {{.ElemWrapper}}{&s.Elts[i]}`))

var findTempl = template.Must(template.New("Find").Parse(`
i := s.Index({{.JoinedKeys}})
if i < 0 {
    return {{.ElemWrapper}}{nil}, false
}
return Get(i), true
`))

func (c *AbstractionContext) MakeSliceImpl(t reflect.Type) *Struct {
	vars := c.getSliceVariables(t)
	i := c.MakeSliceInterface(t) // re-does work, maybe refactor later.
	s := i.Implement(c.TypeForWrapper(t), "s")
	// Some things we'll want to fill in
	origType := c.TypeOriginal(t)
	s.FieldTypes = []string{origType}
	s.FieldNames = []string{"Elts"}

	iter := s.Method("Iter")
	iter.Body = TemplBody(iterTempl, vars)

	index := s.Method("Index")
	index.Body = TemplBody(indexTempl, vars)

	get := s.Method("Get")
	get.Body = TemplBody(getTempl, vars)

	find := s.Method("Find")
	find.Body = TemplBody(findTempl, vars)

	return s
}

func (c *AbstractionContext) MakeInterface(t reflect.Type) *Interface {
	omit := c.OmitFields[t.Name()]
	methods := []*Function{}
	for _, f := range TopLevelFields(t, c.Imports) {
		if omit.Has(f.Name) {
			continue
		}
		methods = append(methods, c.makeGetter(f))
		switch f.Type.Kind() {
		case reflect.Struct:
		default:
			methods = append(methods, c.makeSetter(f))
		}
	}
	return &Interface{
		Name:    t.Name(),
		Methods: methods,
	}
}

func (c *AbstractionContext) MakeImplementation(t reflect.Type) *Struct {
	methods := []*Function{}
	omit := c.OmitFields[t.Name()]
	for _, f := range TopLevelFields(t, c.Imports) {
		if omit.Has(f.Name) {
			continue
		}
		methods = append(methods, c.makeGetterImpl(t, f))
		switch f.Type.Kind() {
		case reflect.Struct:
		default:
			methods = append(methods, c.makeSetterImpl(t, f))
		}
	}
	return &Struct{
		Interface: Interface{
			Name:    c.TypeForWrapper(t),
			Methods: methods,
		},
		Abbrev:     "r",
		FieldNames: []string{""},
		FieldTypes: []string{fmt.Sprintf("*%s.%s", c.findPkgName(t.PkgPath()), t.Name())}, // avoid tm mapping
	}
}

func (c *AbstractionContext) makeGetter(f reflect.StructField) *Function {
	return &Function{
		Name:     "Get" + f.Name,
		RetTypes: []string{c.TypeForGetter(f.Type)},
	}
}

func (c *AbstractionContext) makeSetter(f reflect.StructField) *Function {
	return &Function{
		Name:     "Set" + f.Name,
		ArgNames: []string{"o"},
		ArgTypes: []string{c.Type(f.Type)},
	}
}

func (c *AbstractionContext) makeGetterImpl(t reflect.Type, f reflect.StructField) *Function {
	_, ok := c.Abstract[f.Type]
	var body string
	if ok {
		body = fmt.Sprintf("return %s{&r.%s}", c.TypeForWrapper(f.Type), f.Name)
	} else if f.Type.Kind() == reflect.Struct {
		body = fmt.Sprintf("return &r.%s", f.Name)
	} else {
		body = fmt.Sprintf("return r.%s", f.Name)
	}

	return &Function{
		Name:         "Get" + f.Name,
		AcceptorName: "r",
		AcceptorType: c.TypeForWrapper(t),
		RetTypes:     []string{c.TypeForGetter(f.Type)},
		Body:         []string{body},
	}

}

func (c *AbstractionContext) makeSetterImpl(t reflect.Type, f reflect.StructField) *Function {
	return &Function{
		Name:         "Set" + f.Name,
		AcceptorName: "r",
		AcceptorType: c.TypeForWrapper(t),
		ArgNames:     []string{"o"},
		ArgTypes:     []string{c.Type(f.Type)},
		Body:         []string{fmt.Sprintf("r.%s = o", f.Name)},
	}
}

func (c *AbstractionContext) stringifyType(t reflect.Type, structStar bool, lookupAbstraction bool) string {
	w := bytes.NewBufferString("")
	if lookupAbstraction {
		s, ok := c.Abstract[t]
		if ok {
			w.Write([]byte(s))
			return w.String()
		}
	}
	pkgPath := t.PkgPath()
	switch t.Kind() {
	case reflect.Ptr:
		fmt.Fprintf(w, "*%s", c.stringifyType(t.Elem(), false, lookupAbstraction))
	case reflect.Slice:
		fmt.Fprintf(w, "[]%s", c.stringifyType(t.Elem(), false, lookupAbstraction))
	case reflect.Map:
		fmt.Fprintf(w, "map[%s]%s", c.stringifyType(t.Key(), false, lookupAbstraction), c.stringifyType(t.Elem(), false, lookupAbstraction))
	case reflect.Struct:
		// When we write structs as return values at least, we want to make them pointer-valued so
		// we can do the usual thing of writing-through.
		if structStar {
			fmt.Fprintf(w, "*")
		}
		fallthrough
	default:
		if pkgPath != "" {
			fmt.Fprintf(w, "%s.%s", c.findPkgName(pkgPath), t.Name())
		} else {
			w.Write([]byte(t.Name()))
		}
	}
	return w.String()
}

func TopLevelFields(t reflect.Type, imports map[string]string) []reflect.StructField {
	ret := []reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Kind() != reflect.Struct {
			ret = append(ret, f)
		} else if f.Name == f.Type.Name() {
			ret = append(ret, TopLevelFields(f.Type, imports)...)
		} else {
			ret = append(ret, f)
		}
	}
	return ret
}

func (c *AbstractionContext) findPkgName(pkgPath string) string {
	pkgName, ok := c.Imports[pkgPath]
	if ok {
		return pkgName
	}
	reversed := map[string]string{}
	for p, n := range c.Imports {
		reversed[n] = p
	}

	splitPath := strings.Split(pkgPath, "/")
	candidate := ""
	for i := len(splitPath) - 1; i >= 0; i-- {
		candidate = splitPath[i] + candidate
		_, ok := reversed[candidate]
		if !ok {
			c.Imports[pkgPath] = candidate
			if candidate == "" {
				panic(fmt.Sprintf("Empty package name for %v", pkgPath))
			}
			return candidate
		}
	}
	panic("Could not find name to import as")
	return ""
}
