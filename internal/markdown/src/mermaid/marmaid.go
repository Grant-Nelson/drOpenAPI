package mermaid

import (
	"fmt"
	"strings"

	"github.com/grant-nelson/DrOpenAPI/internal/markdown"
)

type mermaidImp struct {
	factory markdown.Factory
	order   []markdown.Class
	classes map[string]markdown.Class
}

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

func (imp *mermaidImp) Enum(name string, values ...string) {
	if _, has := imp.classes[name]; !has {
		c := imp.newClass(name)
		c.AddEntry(`<<enumeration>>`)
		for _, value := range values {
			c.AddEntry(value)
		}
	}
}

func (imp *mermaidImp) String() string {
	parts := make([]string, len(imp.order))
	for i, class := range imp.order {
		parts[i] = strings.TrimSpace(class.String())
	}
	return fmt.Sprintf("```mermaid\nclassDiagram\ndirection LR\n\n%s\n```\n\n",
		strings.Join(parts, "\n\n"))
}
