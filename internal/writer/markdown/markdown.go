package markdown

import (
	"bytes"
	"fmt"
	"strings"
)

type (
	Markdown interface {
		SetTitle(title string) Markdown
		AddSection(name string) Markdown
		AddSubsection(name string) Markdown
		Plain(msg string, args ...interface{}) Markdown
		Bold(msg string, args ...interface{}) Markdown
		Link(name, href string) Markdown
		NewPar() Markdown
		LineBreak() Markdown
		Mermaid(mermaid string) Markdown
		String() string
	}

	markdownImp struct {
		title         string
		hadNormalText bool
		index         *bytes.Buffer
		body          *bytes.Buffer
	}
)

func New() Markdown {
	imp := &markdownImp{
		index: &bytes.Buffer{},
		body:  &bytes.Buffer{},
	}
	return imp
}

func write(buf *bytes.Buffer, msg string, args ...interface{}) {
	_, err := buf.WriteString(fmt.Sprintf(msg, args...))
	if err != nil {
		panic(err)
	}
}

func nameToLink(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), ` `, `-`)
}

func (imp *markdownImp) SetTitle(title string) Markdown {
	imp.title = title
	return imp
}

func (imp *markdownImp) AddSection(name string) Markdown {
	if imp.hadNormalText {
		write(imp.body, "\n")
	}
	write(imp.body, "## %s\n\n", name)
	write(imp.index, "- [%s](#%s)\n", name, nameToLink(name))
	imp.hadNormalText = false
	return imp
}

func (imp *markdownImp) AddSubsection(name string) Markdown {
	if imp.hadNormalText {
		write(imp.body, "\n")
	}
	write(imp.body, "### %s\n\n", name)
	write(imp.index, "  - [%s](#%s)\n", name, nameToLink(name))
	imp.hadNormalText = false
	return imp
}

func (imp *markdownImp) Plain(msg string, args ...interface{}) Markdown {
	write(imp.body, msg, args...)
	imp.hadNormalText = true
	return imp
}

func (imp *markdownImp) Bold(msg string, args ...interface{}) Markdown {
	write(imp.body, `**`+msg+`**`, args...)
	imp.hadNormalText = true
	return imp
}

func (imp *markdownImp) Link(name, href string) Markdown {
	write(imp.body, `[%s](%s)`, name, href)
	imp.hadNormalText = true
	return imp
}

func (imp *markdownImp) NewPar() Markdown {
	write(imp.body, "\n\n")
	imp.hadNormalText = true
	return imp
}

func (imp *markdownImp) LineBreak() Markdown {
	write(imp.body, "  \n")
	imp.hadNormalText = true
	return imp
}

func (imp *markdownImp) Mermaid(mermaid string) Markdown {
	if imp.hadNormalText {
		write(imp.body, "\n")
	}
	write(imp.body, "```mermaid\n%s\n```\n\n", mermaid)
	imp.hadNormalText = false
	return imp
}

func (imp *markdownImp) String() string {
	buf := &bytes.Buffer{}
	if len(imp.title) > 0 {
		write(buf, "# %s\n\n", imp.title)
	}
	write(buf, "%s\n%s", imp.index.String(), imp.body.String())
	return buf.String()
}
