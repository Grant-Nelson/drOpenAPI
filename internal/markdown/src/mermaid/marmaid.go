package mermaid

import (
	"fmt"
	"strings"

	"github.com/Grant-Nelson/DrOpenAPI/internal/markdown"
)

// mermaidImp is the implementation of the Mermaid interface.
type mermaidImp struct {
	factory markdown.Factory
	order   []markdown.Class
	classes map[string]markdown.Class
}

// New creates a new Mermaid instance.
func New(factory markdown.Factory) markdown.Mermaid {
	return &mermaidImp{
		factory: factory,
		order:   []markdown.Class{},
		classes: map[string]markdown.Class{},
	}
}

func (imp *mermaidImp) Has(name string) bool {
	_, has := imp.classes[name]
	return has
}

func (imp *mermaidImp) newClass(name string) markdown.Class {
	c := imp.factory.Class(name)
	imp.classes[name] = c
	imp.order = append(imp.order, c)
	return c
}

func (imp *mermaidImp) Class(name string) markdown.Class {
	if c, has := imp.classes[name]; has {
		return c
	}
	return imp.newClass(name)
}

func (imp *mermaidImp) Interface(name string) markdown.Class {
	if c, has := imp.classes[name]; has {
		return c
	}
	c := imp.newClass(name)
	c.AddEntry(`<<Interface>>`)
	return c
}

func (imp *mermaidImp) Enum(name string, values ...string) bool {
	if _, has := imp.classes[name]; has {
		return false
	}

	c := imp.newClass(name)
	c.AddEntry(`<<enumeration>>`)
	for _, value := range values {
		c.AddEntry(value)
	}
	return true
}

func (imp *mermaidImp) String() string {
	if len(imp.order) == 0 {
		return ``
	}
	parts := make([]string, len(imp.order))
	for i, class := range imp.order {
		parts[i] = strings.TrimSpace(class.String())
	}
	return fmt.Sprintf("```mermaid\nclassDiagram\ndirection LR\n\n%s\n```\n\n",
		strings.Join(parts, "\n\n"))
}
