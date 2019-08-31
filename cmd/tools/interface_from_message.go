package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

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
	return c.stringifyType(t, false)
}

func (c *AbstractionContext) TypeForGetter(t reflect.Type) string {
	return c.stringifyType(t, true)
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
	keys := c.SliceKeys[t.Elem().Name()]
	fields := TopLevelFields(t.Elem(), c.Imports)
	var keyFields []reflect.StructField
	keyArgNames := []string{}
	keyArgTypes := []string{}
	argNames := []string{}
	argTypes := []string{}
	keyFields = []reflect.StructField{}
	for _, f := range fields {
		argNames = append(argNames, strings.ToLower(f.Name))
		argTypes = append(argTypes, c.Type(f.Type))
		if len(keys) == 0 || keys.Has(f.Name) {
			keyArgNames = append(keyArgNames, strings.ToLower(f.Name))
			keyArgTypes = append(keyArgTypes, c.Type(f.Type))
			keyFields = append(keyFields, f)
		}
	}
	elemType := c.Type(t.Elem())
	methods := []*Function{
		&Function{Name: "Iter", RetTypes: []string{"chan " + elemType}},
		&Function{Name: "Index", ArgNames: keyArgNames, ArgTypes: keyArgTypes, RetTypes: []string{"int"}},
		&Function{Name: "Get", ArgNames: keyArgNames, ArgTypes: keyArgTypes, RetTypes: []string{elemType}},
		&Function{Name: "Remove", ArgNames: keyArgNames, ArgTypes: keyArgTypes},
		&Function{Name: "Set", ArgNames: argNames, ArgTypes: argTypes},
	}
	// MatchesF for each key field
	for i, f := range keyFields {
		methods = append(methods, &Function{
			Name:     "Matches" + f.Name,
			ArgNames: []string{keyArgNames[i]},
			ArgTypes: []string{keyArgTypes[i]},
			RetTypes: []string{"[]" + elemType},
		})
	}

	return &Interface{
		Name:    t.Elem().Name() + "Slice",
		Methods: methods,
	}
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
		Name:       c.TypeForWrapper(t),
		Abbrev:     "r",
		FieldNames: []string{""},
		FieldTypes: []string{fmt.Sprintf("*%s.%s", c.findPkgName(t.PkgPath()), t.Name())}, // avoid tm mapping
		Methods:    methods,
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

func (c *AbstractionContext) stringifyType(t reflect.Type, structStar bool) string {
	w := bytes.NewBufferString("")
	s, ok := c.Abstract[t]
	if ok {
		w.Write([]byte(s))
		return w.String()
	}
	pkgPath := t.PkgPath()
	switch t.Kind() {
	case reflect.Ptr:
		fmt.Fprintf(w, "*%s", c.stringifyType(t.Elem(), false))
	case reflect.Slice:
		fmt.Fprintf(w, "[]%s", c.stringifyType(t.Elem(), false))
	case reflect.Map:
		fmt.Fprintf(w, "map[%s]%s", c.stringifyType(t.Key(), false), c.stringifyType(t.Elem(), false))
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
