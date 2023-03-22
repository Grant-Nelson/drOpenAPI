package factory

import (
	"github.com/Grant-Nelson/drOpenAPI/internal/markdown"
	"github.com/Grant-Nelson/drOpenAPI/internal/markdown/src/class"
	srcMD "github.com/Grant-Nelson/drOpenAPI/internal/markdown/src/markdown"
	"github.com/Grant-Nelson/drOpenAPI/internal/markdown/src/mermaid"
	"github.com/Grant-Nelson/drOpenAPI/internal/markdown/src/stringBuffer"
	"github.com/Grant-Nelson/drOpenAPI/internal/markdown/src/text"
)

// factory is the implementation of the Factory interface.
type factory struct{}

// New creates a new instance of the markdown factory.
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
