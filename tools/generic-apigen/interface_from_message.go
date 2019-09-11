// Copyright © 2019 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"k8s.io/apimachinery/pkg/util/sets"
)

type GenerationOptions struct {
	PkgPath             string
	InterfacePkg        string
	MyPkg               string
	Roots               []reflect.Type
	OmitFields          map[string]sets.String
	OmitImplementations map[string]sets.String
	SliceKeys           map[string]sets.String
}

type AbstractionContext struct {
	GenerationOptions
	Imports  map[string]string       // package path to name imported as
	Abstract map[reflect.Type]string // concrete message type to interface name for it
}

func (o GenerationOptions) NewAbstractionContext(imports map[string]string) *AbstractionContext {
	if imports == nil {
		imports = map[string]string{}
	}
	ret := &AbstractionContext{
		GenerationOptions: o,
		Imports:           imports,
		Abstract:          map[reflect.Type]string{},
	}
	for _, t := range o.Roots {
		interfacesToWrite(t, o.PkgPath, ret.Imports, ret.Abstract)
	}
	return ret
}

func (c *AbstractionContext) sortedTypesToGenerate() []reflect.Type {
	ret := []reflect.Type{}
	for t := range c.Abstract {
		ret = append(ret, t)
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].Name() < ret[j].Name() })
	return ret
}

func (c *AbstractionContext) WriteInterfaceFile(w io.Writer) {
	myPkgShort := c.findPkgName(c.MyPkg)
	declarations := []Declaration{}
	for _, t := range c.sortedTypesToGenerate() {
		switch t.Kind() {
		case reflect.Struct:
			declarations = append(declarations, c.MakeInterface(t))
		case reflect.Slice:
			declarations = append(declarations, c.MakeSliceInterface(t))
		}
	}
	sort.Slice(declarations, func(i, j int) bool { return declarations[i].GetName() < declarations[j].GetName() })
	fmt.Fprintf(w, "package %s\n\n%s\n", myPkgShort, c.WriteImports())
	for _, d := range declarations {
		d.WriteDeclaration(w)
	}

}

func (c *AbstractionContext) WriteImplementationFile(w io.Writer) {
	myPkgShort := c.findPkgName(c.MyPkg)
	declarations := []Declaration{}
	for _, t := range c.sortedTypesToGenerate() {
		switch t.Kind() {
		case reflect.Struct:
			declarations = append(declarations, c.MakeImplementation(t))
		case reflect.Slice:
			declarations = append(declarations, c.MakeSliceImpl(t))
		}
	}
	sort.Slice(declarations, func(i, j int) bool { return declarations[i].GetName() < declarations[j].GetName() })
	fmt.Fprintf(w, "package %s\n\n%s\n", myPkgShort, c.WriteImports())
	for _, d := range declarations {
		d.WriteDeclaration(w)
	}

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

func (c *AbstractionContext) TypeForWrapper(t reflect.Type) string {
	pkgName := c.findPkgName(c.MyPkg)
	switch t.Kind() {
	case reflect.Struct:
		prefix := strings.Title(pkgName)
		return prefix + t.Name()
	case reflect.Slice:
		prefix := strings.Title(pkgName)
		return prefix + t.Elem().Name() + "Slice"
	}
	fmt.Fprintf(os.Stderr, "The type is %v\n", t)
	panic("Got to the end and can't get a typedef name")
	return ""
}

func (c *AbstractionContext) WriteImports() string {
	b := bytes.NewBufferString("")
	fmt.Fprintf(b, "import(\n")
	for path, name := range c.Imports {
		if path == c.MyPkg {
			continue
		}
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
	methods := []*Function{
		{Name: "Iter", RetTypes: []string{"chan " + vars.ElemType}},
		{Name: "Index", ArgNames: vars.KeyArgNames, ArgTypes: vars.KeyArgTypes, RetTypes: []string{"int"}},
		{Name: "Get", ArgNames: []string{"i"}, ArgTypes: []string{"int"}, RetTypes: []string{vars.ElemType}},
		{Name: "Find", ArgNames: vars.KeyArgNames, ArgTypes: vars.KeyArgTypes, RetTypes: []string{vars.ElemType, "bool"}},
		{
			Name:     "Filter",
			ArgNames: []string{"predicate"},
			ArgTypes: []string{fmt.Sprintf("func (e %s) bool", vars.ElemType)},
			RetTypes: []string{vars.InterfaceType},
		},
		{Name: "Upsert", ArgNames: vars.ArgNames, ArgTypes: vars.ArgTypes, RetTypes: []string{vars.ElemType}},
	}
	return &Interface{
		Name:    vars.InterfaceType,
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
	ElemType      string
	ElemWrapper   string
	OrigElemType  string
	KeyFields     []fieldArgMatch
	AllFields     []fieldArgMatch
	ArgNames      []string
	ArgTypes      []string
	JoinedArgs    string
	KeyArgNames   []string
	KeyArgTypes   []string
	JoinedKeys    string
	WrapperType   string
	OrigType      string
	InterfaceType string
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
		optional := f.Type.Kind() == reflect.Ptr
		ret.AllFields = append(ret.AllFields, fieldArgMatch{f.Name, n, typ, optional})
		if len(keys) == 0 || keys.Has(f.Name) {
			ret.KeyArgNames = append(ret.KeyArgNames, strings.ToLower(f.Name))
			ret.KeyArgTypes = append(ret.KeyArgTypes, typ)
			if optional {
				// when we loop through these we need to know the non-ptr type
				ret.KeyFields = append(ret.KeyFields, fieldArgMatch{f.Name, n, c.Type(f.Type.Elem()), true})
			} else {
				ret.KeyFields = append(ret.KeyFields, fieldArgMatch{f.Name, n, typ, false})
			}
		}
	}
	ret.JoinedArgs = strings.Join(ret.ArgNames, ", ")
	ret.JoinedKeys = strings.Join(ret.KeyArgNames, ", ")
	ret.InterfaceType = c.Type(t)
	ret.ElemType = c.Type(t.Elem())
	ret.ElemWrapper = c.TypeForWrapper(t.Elem())
	ret.WrapperType = c.TypeForWrapper(t)
	ret.OrigType = c.TypeOriginal(t)
	ret.OrigElemType = c.TypeOriginal(t.Elem())
	return ret
}

var indexTempl = template.Must(template.New("Index").Parse(`
for i, elt := range s.Elts {
    {{range .KeyFields}}
    {{if .Optional}}
    var v {{.Type}}
    if elt.{{.Name}} != nil {
        v = *elt.{{.Name}}
    } else if {{.Arg}} != nil {
        continue
    }
    if v != *{{.Arg}} { continue }
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
return s.Get(i), true
`))

var filterTempl = template.Must(template.New("Filter").Parse(`
ret := {{.OrigType}}{}
for _, elt := range s.Elts {
    if predicate({{.ElemWrapper}}{&elt}) {
        ret = append(ret, elt)
    }
}
return {{.WrapperType}}{ret}
`))

var upsertTempl = template.Must(template.New("Upsert").Parse(`
ins := {{.OrigElemType}}{ {{.JoinedArgs}} }
idx := s.Index({{.JoinedKeys}})
if idx >= 0 {
    s.Elts[idx] = ins
} else {
    idx = len(s.Elts)
    s.Elts = append(s.Elts, ins)
}
return {{.ElemWrapper}}{&s.Elts[idx]}
`))

func (c *AbstractionContext) MakeSliceImpl(t reflect.Type) *Struct {
	vars := c.getSliceVariables(t)

	i := c.MakeSliceInterface(t) // re-does work, maybe refactor later.
	s := i.Implement(vars.WrapperType, "s")

	omit := c.OmitImplementations[t.Elem().Name()+"Slice"]
	filteredMethods := []*Function{}
	for _, m := range s.Methods {
		if !omit.Has(m.Name) {
			filteredMethods = append(filteredMethods, m)
		}
	}
	s.Methods = filteredMethods
	// Some things we'll want to fill in
	s.FieldTypes = []string{vars.OrigType}
	s.FieldNames = []string{"Elts"}

	if iter := s.Method("Iter"); iter != nil {
		iter.Body = TemplBody(iterTempl, vars)
	}

	if index := s.Method("Index"); index != nil {
		index.Body = TemplBody(indexTempl, vars)
	}

	if get := s.Method("Get"); get != nil {
		get.Body = TemplBody(getTempl, vars)
	}

	if find := s.Method("Find"); find != nil {
		find.Body = TemplBody(findTempl, vars)
	}

	if filter := s.Method("Filter"); filter != nil {
		filter.Body = TemplBody(filterTempl, vars)
	}

	if upsert := s.Method("Upsert"); upsert != nil {
		upsert.Body = TemplBody(upsertTempl, vars)
	}

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
	omit := c.OmitFields[t.Name()].Union(c.OmitImplementations[t.Name()])
	for _, f := range TopLevelFields(t, c.Imports) {
		if omit.Has(f.Name) {
			continue
		}
		if strings.HasPrefix(f.Name, "Deprecated") {
			continue
		}
		fmt.Fprintf(os.Stderr, "Generating %s on %s\n", f.Name, t.Name())
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
	var raw string
	if f.Type.Kind() == reflect.Struct {
		raw = "&r." + f.Name
	} else {
		raw = "r." + f.Name
	}
	if ok {
		body = fmt.Sprintf("return %s{%s}", c.TypeForWrapper(f.Type), raw)
	} else {
		body = fmt.Sprintf("return %s", raw)
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
	var body string
	_, ok := c.Abstract[f.Type]
	if ok {
		body = fmt.Sprintf("r.%s = o.(%s).Elts", f.Name, c.TypeForWrapper(f.Type))
	} else {
		body = fmt.Sprintf("r.%s = o", f.Name)
	}
	return &Function{
		Name:         "Set" + f.Name,
		AcceptorName: "r",
		AcceptorType: c.TypeForWrapper(t),
		ArgNames:     []string{"o"},
		ArgTypes:     []string{c.Type(f.Type)},
		Body:         []string{body},
	}
}

func (c *AbstractionContext) stringifyType(t reflect.Type, structStar bool, lookupAbstraction bool) string {
	w := bytes.NewBufferString("")
	if lookupAbstraction {
		s, ok := c.Abstract[t]
		if ok {
			if c.MyPkg == c.InterfacePkg {
				w.Write([]byte(s))
			} else {
				fmt.Fprintf(w, "%s.%s", c.findPkgName(c.InterfacePkg), s)
			}
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
