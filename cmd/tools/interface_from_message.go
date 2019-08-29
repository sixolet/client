package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

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

func InterfacesToWrite(t reflect.Type, pkgPath string, imports map[string]string) map[reflect.Type]string {
	ret := map[reflect.Type]string{}
	interfacesToWrite(t, pkgPath, imports, ret)
	return ret
}

func interfacesToWrite(t reflect.Type, pkgPath string, imports map[string]string, tm map[reflect.Type]string) {
	if t.PkgPath() != pkgPath {
		return
	}
	tm[t] = t.Name()
	for _, f := range TopLevelFields(t, imports) {
		switch f.Type.Kind() {
		case reflect.Map:
			interfacesToWrite(f.Type.Key(), pkgPath, imports, tm)
			fallthrough
		case reflect.Ptr:
			fallthrough
		case reflect.Slice:
			interfacesToWrite(f.Type.Elem(), pkgPath, imports, tm)
		case reflect.Interface:
			// pass
		case reflect.Struct:
			interfacesToWrite(f.Type, pkgPath, imports, tm)
		}
	}
}

func WriteInterface(t reflect.Type, imports map[string]string, tm map[reflect.Type]string) string {
	b := bytes.NewBufferString("")
	fmt.Fprintf(b, "type %s interface {\n", t.Name())
	for _, f := range TopLevelFields(t, imports) {
		writeGetter(b, f, imports, tm)
	}
	fmt.Fprintf(b, "}\n")
	return b.String()
}

func writeGetter(w io.Writer, f reflect.StructField, imports map[string]string, tm map[reflect.Type]string) {
	fmt.Fprintf(w, "\t// Get %s\n\tGet%s() ", f.Name, f.Name)
	writeType(w, f.Type, imports, tm)
	w.Write([]byte("\n"))
}

func writeType(w io.Writer, t reflect.Type, imports map[string]string, tm map[reflect.Type]string) {
	s, ok := tm[t]
	if ok {
		w.Write([]byte(s))
		return
	}
	pkgPath := t.PkgPath()
	switch t.Kind() {
	case reflect.Ptr:
		w.Write([]byte("*"))
		writeType(w, t.Elem(), imports, tm)
	case reflect.Slice:
		fmt.Fprintf(w, "[]")
		writeType(w, t.Elem(), imports, tm)
	case reflect.Map:
		fmt.Fprintf(w, "map[")
		writeType(w, t.Key(), imports, tm)
		fmt.Fprintf(w, "]")
		writeType(w, t.Elem(), imports, tm)
	default:
		if pkgPath != "" {
			mapped, ok := tm[t]
			if ok {
				fmt.Fprintf(w, mapped)
			} else {
				fmt.Fprintf(w, "%s.%s", findPkgName(imports, pkgPath), t.Name())
			}
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
			return candidate
		}
	}
	panic("Could not find name to import as")
	return ""
}
