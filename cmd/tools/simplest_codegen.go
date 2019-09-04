package main

import (
	"bytes"
	"fmt"
	"io"
)

type Declaration interface {
	WriteDeclaration(b io.Writer)
}

type Function struct {
	Name         string
	ArgTypes     []string
	ArgNames     []string
	AcceptorType string
	AcceptorName string
	RetTypes     []string
	Body         []string
}

type Interface struct {
	Name    string
	Methods []*Function
}

type Struct struct {
	Interface
	Abbrev     string
	FieldTypes []string
	FieldNames []string
}

func (s *Struct) PopulateAcceptors() {
	for i, f := range s.Methods {
		if f.AcceptorType == "" {
			s.Methods[i].AcceptorType = s.Name
		}
		if f.AcceptorName == "" {
			s.Methods[i].AcceptorName = s.Abbrev
		}
	}
}

func (s *Struct) WriteDefinition(b io.Writer) {
	fmt.Fprintf(b, "\ntype %s struct {\n", s.Name)
	for i, n := range s.FieldNames {
		t := s.FieldTypes[i]
		if n == "" {
			fmt.Fprintf(b, "\t%s\n", t)
		} else {
			fmt.Fprintf(b, "\t%s %s\n", n, t)
		}
	}
	fmt.Fprintf(b, "}\n")
}

func (s *Struct) WriteMethods(b io.Writer) {
	s.PopulateAcceptors()
	for _, m := range s.Methods {
		m.WriteDeclaration(b)
	}
}

func (s *Struct) WriteDeclaration(b io.Writer) {
	s.WriteDefinition(b)
	s.WriteMethods(b)
}

func (i *Interface) WriteDeclaration(b io.Writer) {
	fmt.Fprintf(b, "\ntype %s interface {\n", i.Name)
	for _, f := range i.Methods {
		fmt.Fprintf(b, "\t")
		f.WriteSignatureForInterface(b)
	}
	fmt.Fprintf(b, "}\n")
}

func (i *Interface) Definition() string {
	b := bytes.NewBufferString("")
	i.WriteDeclaration(b)
	return b.String()
}

func (i *Interface) Method(name string) *Function {
	for _, m := range i.Methods {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (i *Interface) Implement(name string, abbrev string) *Struct {
	ret := &Struct{
		Interface: Interface{
			Name:    name,
			Methods: make([]*Function, 0, len(i.Methods)),
		},
		Abbrev: abbrev,
	}
	for _, f := range i.Methods {
		newF := *f
		ret.Methods = append(ret.Methods, &newF)
	}
	ret.PopulateAcceptors()
	return ret
}

func (f *Function) writeMainPartOfSignature(b io.Writer) {
	fmt.Fprintf(b, "%s(", f.Name)
	for i, _ := range f.ArgTypes {
		if i != 0 {
			fmt.Fprintf(b, ", ")
		}
		fmt.Fprintf(b, "%s %s", f.ArgNames[i], f.ArgTypes[i])
	}
	fmt.Fprintf(b, ")")
	switch len(f.RetTypes) {
	case 0:
	case 1:
		fmt.Fprintf(b, " %s", f.RetTypes[0])
	default:
		fmt.Fprintf(b, "(")
		for i, t := range f.RetTypes {
			if i != 0 {
				fmt.Fprintf(b, ", ")
			}
			fmt.Fprintf(b, t)
		}
		fmt.Fprintf(b, ")")
	}

}

func (f *Function) WriteDeclaration(b io.Writer) {
	fmt.Fprintf(b, "func ")
	if f.AcceptorName != "" {
		fmt.Fprintf(b, "(%s %s) ", f.AcceptorName, f.AcceptorType)
	}
	f.writeMainPartOfSignature(b)
	fmt.Fprintf(b, " {\n")
	for _, l := range f.Body {
		fmt.Fprintf(b, "\t%s\n", l)
	}
	fmt.Fprintf(b, "}\n")
}

func (f *Function) Declaration() string {
	b := bytes.NewBufferString("")
	f.WriteDeclaration(b)
	return b.String()
}

func (f *Function) WriteSignatureForInterface(b io.Writer) {
	f.writeMainPartOfSignature(b)
	fmt.Fprintf(b, "\n")
}

func (f *Function) SignatureForInterface() string {
	b := bytes.NewBufferString("")
	f.WriteSignatureForInterface(b)
	return b.String()
}