package text

import (
	"github.com/grant-nelson/DrOpenAPI/internal/markdown"
)

// textImp is the implementation of the Text interface.
type textImp struct {
	buf markdown.StringBuffer
}

// New creates a new Text instance.
func New(factory markdown.Factory) markdown.Text {
	return &textImp{
		buf: factory.StringBuffer(),
	}
}

// convertNameToRef will convert a section/subsection name into
// a local link reference to that section/subsection.
func convertNameToRef(name string) string {
	href := make([]rune, 0, len(name))
	hasHyphen := false
	for _, c := range name {
		switch {
		case c >= 'a' && c <= 'z':
			href = append(href, c)
			hasHyphen = false
		case c >= 'A' && c <= 'Z':
			href = append(href, c+('a'-'A'))
			hasHyphen = false
		case c == ' ' || c == '-':
			if !hasHyphen {
				hasHyphen = true
				href = append(href, '-')
			}
		}
	}
	return `#` + string(href)
}

func (imp *textImp) Bold(msg string, args ...interface{}) markdown.Text {
	if len(msg) > 0 {
		imp.buf.Write(`**`+msg+`**`, args...)
	}
	return imp
}

func (imp *textImp) Code(msg string, args ...interface{}) markdown.Text {
	if len(msg) > 0 {
		imp.buf.Write("`"+msg+"`", args...)
	}
	return imp
}

func (imp *textImp) Write(msg string, args ...interface{}) markdown.Text {
	imp.buf.Write(msg, args...)
	return imp
}

func (imp *textImp) Link(text, href string) markdown.Text {
	if len(text) > 0 {
		if len(href) > 0 {
			imp.buf.Write(`[%s](%s)`, text, href)
		} else {
			imp.buf.Write(text)
		}
	}
	return imp
}

func (imp *textImp) Ref(text, name string) markdown.Text {
	return imp.Link(text, convertNameToRef(name))
}

func (imp *textImp) LineBreak() markdown.Text {
	imp.buf.Write("  \n")
	return imp
}

func (imp *textImp) String() string {
	return imp.buf.String()
}
