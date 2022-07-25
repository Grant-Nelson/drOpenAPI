package factory

import (
	"github.com/grant-nelson/DrOpenAPI/internal/markdown"
	"github.com/grant-nelson/DrOpenAPI/internal/markdown/src/class"
	srcMD "github.com/grant-nelson/DrOpenAPI/internal/markdown/src/markdown"
	"github.com/grant-nelson/DrOpenAPI/internal/markdown/src/mermaid"
	"github.com/grant-nelson/DrOpenAPI/internal/markdown/src/stringBuffer"
	"github.com/grant-nelson/DrOpenAPI/internal/markdown/src/text"
)

type factory struct{}

func New() markdown.Factory {
	return &factory{}
}

func (f *factory) StringBuffer() markdown.StringBuffer {
	return stringBuffer.New()
}

func (f *factory) Markdown(title string) markdown.Markdown {
	return srcMD.New(f, title)
}

func (f *factory) Text() markdown.Text {
	return text.New(f)
}

func (f *factory) Mermaid() markdown.Mermaid {
	return mermaid.New(f)
}

func (f *factory) Class(name string) markdown.Class {
	return class.New(f, name)
}
