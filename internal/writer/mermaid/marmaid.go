package mermaid

import (
	"bytes"
	"fmt"
)

type (
	Mermaid interface {
		Has(name string) bool
		Class(name string) Class
		Interface(name string) Class
		Enum(name string, values ...string) Class
		String() string
	}

	Class interface {
		AddEntry(entry string) Class
		AddMember(name, typeName string) Class
		ConnectTo(name string) Class
		String() string
	}

	mermaidImp struct {
		order   []*classImp
		classes map[string]*classImp
	}

	classImp struct {
		name   string
		body   *bytes.Buffer
		hasCon map[string]bool
		cons   *bytes.Buffer
	}
)

func New() Mermaid {
	return &mermaidImp{
		order:   []*classImp{},
		classes: map[string]*classImp{},
	}
}

func write(buf *bytes.Buffer, msg string, args ...interface{}) {
	_, err := buf.WriteString(fmt.Sprintf(msg, args...))
	if err != nil {
		panic(err)
	}
}

func (imp *mermaidImp) Has(name string) bool {
	_, has := imp.classes[name]
	return has
}

func (imp *mermaidImp) Class(name string) Class {
	if c, has := imp.classes[name]; has {
		return c
	}

	c := &classImp{
		name:   name,
		body:   &bytes.Buffer{},
		hasCon: map[string]bool{},
		cons:   &bytes.Buffer{},
	}
	imp.classes[name] = c
	imp.order = append(imp.order, c)
	return c
}

func (imp *mermaidImp) Interface(name string) Class {
	c := imp.Class(name)
	c.AddEntry(`<<Interface>>`)
	return c
}

func (imp *mermaidImp) Enum(name string, values ...string) Class {
	c := imp.Class(name)
	c.AddEntry(`<<enumeration>>`)
	for _, value := range values {
		c.AddEntry(value)
	}
	return c
}

func (imp *mermaidImp) String() string {
	buf := &bytes.Buffer{}
	write(buf, "classDiagram\ndirection LR\n\n")
	for _, class := range imp.order {
		write(buf, class.String())
	}
	return buf.String()
}

func (imp *classImp) AddEntry(entry string) Class {
	write(imp.body, "  %s\n", entry)
	return imp
}

func (imp *classImp) AddMember(name, typeName string) Class {
	write(imp.body, "  %s %s\n", typeName, name)
	return imp
}

func (imp *classImp) ConnectTo(name string) Class {
	if !imp.hasCon[name] {
		imp.hasCon[name] = true
		write(imp.cons, "%s --> %s\n", imp.name, name)
	}
	return imp
}

func (imp *classImp) String() string {
	buf := &bytes.Buffer{}
	write(buf, "class %s {\n%s}\n", imp.name, imp.body.String())
	if imp.cons.Len() > 0 {
		write(buf, "%s\n", imp.cons.String())
	}
	return buf.String()
}
