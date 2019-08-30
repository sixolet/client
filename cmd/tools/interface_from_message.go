package main

import (
	"bytes"
	"fmt"
	"io"
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

func WriteImports(imports map[string]string) string {
	b := bytes.NewBufferString("")
	fmt.Fprintf(b, "import(\n")
	for path, name := range imports {
		if strings.HasSuffix(path, name) {
			fmt.Fprintf(b, "\t\"%s\"\n", path)
		} else {
			fmt.Fprintf(b, "\t%s \"%s\"\n", name, path)
		}
	}
	fmt.Fprintf(b, ")\n")
	return b.String()
}

func InterfacesToWrite(tt []reflect.Type, pkgPath string, imports map[string]string) map[reflect.Type]string {
	ret := map[reflect.Type]string{}
	for _, t := range tt {
		interfacesToWrite(t, pkgPath, imports, ret)
	}
	return ret
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

func WriteSliceInterface(t reflect.Type, imports map[string]string, tm map[reflect.Type]string, op GenerationOptions) string {
	keys := op.SliceKeys[t.Elem().Name()]
	fields := TopLevelFields(t.Elem(), imports)
	var keyFields []reflect.StructField
	if len(keys) > 0 {
		keyFields = []reflect.StructField{}
		for _, f := range fields {
			if keys.Has(f.Name) {
				keyFields = append(keyFields, f)
			}
		}
	} else {
		keyFields = fields
	}

	b := bytes.NewBufferString("")

	fmt.Fprintf(b, "\n\ntype %sSlice interface {\n", t.Elem().Name())

	// Iter()
	fmt.Fprintf(b, "\tIter() chan ")
	writeType(b, t.Elem(), imports, tm, false)
	fmt.Fprintf(b, "\n")
	// MatchesF for each key field
	for _, f := range keyFields {
		fmt.Fprintf(b, "\tMatches%s(%s ", f.Name, strings.ToLower(f.Name))
		writeType(b, f.Type, imports, tm, false)
		fmt.Fprintf(b, ") []")
		writeType(b, t.Elem(), imports, tm, false)
		fmt.Fprintf(b, "\n")
	}
	// Index(for all key fields) -1 is not present
	// Get(for all key fields)
	// Set(for all fields)
	// Remove(for all key fields)
	fmt.Fprintf(b, "}\n\n")
	return b.String()
}

func MakeInterface(t reflect.Type, imports map[string]string, tm map[reflect.Type]string, op GenerationOptions) Interface {
	omit := op.OmitFields[t.Name()]
	methods := []Function{}
	for _, f := range TopLevelFields(t, imports) {
		if omit.Has(f.Name) {
			continue
		}
		methods = append(methods, makeGetter(f, imports, tm))
		switch f.Type.Kind() {
		case reflect.Struct:
		default:
			methods = append(methods, makeSetter(f, imports, tm))
		}
	}
	return Interface{
		Name:    t.Name(),
		Methods: methods,
	}
}

func WriteInterface(t reflect.Type, imports map[string]string, tm map[reflect.Type]string, op GenerationOptions) string {
	b := bytes.NewBufferString("")
	i := MakeInterface(t, imports, tm, op)
	i.WriteDefinition(b)
	return b.String()
}

func typedefName(t reflect.Type, imports map[string]string) string {
	switch t.Kind() {
	case reflect.Struct:
		pkgName := findPkgName(imports, t.PkgPath())
		prefix := strings.Title(pkgName)
		return prefix + t.Name()
	case reflect.Slice:
		pkgName := findPkgName(imports, t.Elem().PkgPath())
		prefix := strings.Title(pkgName)
		return prefix + t.Name() + "Slice"
	}
	panic("Got to the end and can't get a typedef name")
	return ""
}

func MakeImplementation(t reflect.Type, imports map[string]string, tm map[reflect.Type]string, op GenerationOptions) Struct {
	methods := []Function{}
	omit := op.OmitFields[t.Name()]
	for _, f := range TopLevelFields(t, imports) {
		if omit.Has(f.Name) {
			continue
		}
		methods = append(methods, makeGetterImpl(t, f, imports, tm))
		switch f.Type.Kind() {
		case reflect.Struct:
		default:
			methods = append(methods, makeSetterImpl(t, f, imports, tm))
		}
	}
	return Struct{
		Name:       typedefName(t, imports),
		Abbrev:     "r",
		FieldNames: []string{""},
		FieldTypes: []string{fmt.Sprintf("*%s.%s", findPkgName(imports, t.PkgPath()), t.Name())}, // avoid tm mapping
		Methods:    methods,
	}
}

func WriteImplementation(t reflect.Type, imports map[string]string, tm map[reflect.Type]string, op GenerationOptions) string {
	b := bytes.NewBufferString("")
	impl := MakeImplementation(t, imports, tm, op)
	impl.WriteDeclaration(b)
	return b.String()
}

func stringifyType(t reflect.Type, imports map[string]string, tm map[reflect.Type]string, structToPtr bool) string {
	b := bytes.NewBufferString("")
	writeType(b, t, imports, tm, structToPtr)
	return b.String()
}

func makeGetter(f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) Function {
	return Function{
		Name:     "Get" + f.Name,
		RetTypes: []string{stringifyType(f.Type, imports, tm, true)},
	}
}

func writeGetter(w io.Writer, f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) {
	getter := makeGetter(f, imports, tm)
	fmt.Fprintf(w, "\t%s", getter.SignatureForInterface())
}

func makeSetter(f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) Function {
	return Function{
		Name:     "Set" + f.Name,
		ArgNames: []string{"o"},
		ArgTypes: []string{stringifyType(f.Type, imports, tm, false)},
	}
}

func writeSetter(w io.Writer, f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) {
	setter := makeSetter(f, imports, tm)
	fmt.Fprintf(w, "\t%s", setter.SignatureForInterface())
}

func makeGetterImpl(t reflect.Type, f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) Function {
	_, ok := tm[f.Type]
	var body string
	if ok {
		body = fmt.Sprintf("return %s{&r.%s}", typedefName(f.Type, imports), f.Name)
	} else if f.Type.Kind() == reflect.Struct {
		body = fmt.Sprintf("return &r.%s", f.Name)
	} else {
		body = fmt.Sprintf("return r.%s", f.Name)
	}

	return Function{
		Name:         "Get" + f.Name,
		AcceptorName: "r",
		AcceptorType: typedefName(t, imports),
		RetTypes:     []string{stringifyType(f.Type, imports, tm, true)},
		Body:         []string{body},
	}

}

func writeGetterImpl(w io.Writer, t reflect.Type, f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) {
	impl := makeGetterImpl(t, f, imports, tm)
	impl.WriteDeclaration(w)
}

func makeSetterImpl(t reflect.Type, f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) Function {
	return Function{
		Name:         "Set" + f.Name,
		AcceptorName: "r",
		AcceptorType: typedefName(t, imports),
		ArgNames:     []string{"o"},
		ArgTypes:     []string{stringifyType(f.Type, imports, tm, false)},
		Body:         []string{fmt.Sprintf("r.%s = o", f.Name)},
	}
}

func writeSetterImpl(w io.Writer, t reflect.Type, f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) {
	impl := makeSetterImpl(t, f, imports, tm)
	impl.WriteDeclaration(w)
}

func writeType(w io.Writer, t reflect.Type, imports map[string]string, tm map[reflect.Type]string, structStar bool) {
	s, ok := tm[t]
	if ok {
		w.Write([]byte(s))
		return
	}
	pkgPath := t.PkgPath()
	switch t.Kind() {
	case reflect.Ptr:
		w.Write([]byte("*"))
		writeType(w, t.Elem(), imports, tm, false)
	case reflect.Slice:
		fmt.Fprintf(w, "[]")
		writeType(w, t.Elem(), imports, tm, false)
	case reflect.Map:
		fmt.Fprintf(w, "map[")
		writeType(w, t.Key(), imports, tm, false)
		fmt.Fprintf(w, "]")
		writeType(w, t.Elem(), imports, tm, false)
	case reflect.Struct:
		// When we write structs as return values at least, we want to make them pointer-valued so
		// we can do the usual thing of writing-through.
		if structStar {
			fmt.Fprintf(w, "*")
		}
		fallthrough
	default:
		if pkgPath != "" {
			fmt.Fprintf(w, "%s.%s", findPkgName(imports, pkgPath), t.Name())
		} else {
			w.Write([]byte(t.Name()))
		}
	}
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

func findPkgName(imports map[string]string, pkgPath string) string {
	pkgName, ok := imports[pkgPath]
	if ok {
		return pkgName
	}
	reversed := map[string]string{}
	for p, n := range imports {
		reversed[n] = p
	}

	splitPath := strings.Split(pkgPath, "/")
	candidate := ""
	for i := len(splitPath) - 1; i >= 0; i-- {
		candidate = splitPath[i] + candidate
		_, ok := reversed[candidate]
		if !ok {
			imports[pkgPath] = candidate
			if candidate == "" {
				panic(fmt.Sprintf("Empty package name for %v", pkgPath))
			}
			return candidate
		}
	}
	panic("Could not find name to import as")
	return ""
}
