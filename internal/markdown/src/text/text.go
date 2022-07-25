package text

import (
	"github.com/grant-nelson/DrOpenAPI/internal/markdown"
)

type textImp struct {
	buf markdown.StringBuffer
}

func New(factory markdown.Factory) markdown.Text {
	return &textImp{
		buf: factory.StringBuffer(),
	}
}

func (imp *textImp) Bold(msg string, args ...interface{}) markdown.Text {
	imp.buf.Write(`**`+msg+`**`, args...)
	return imp
}

func (imp *textImp) Code(msg string, args ...interface{}) markdown.Text {
	imp.buf.Write("`"+msg+"`", args...)
	return imp
}

func (imp *textImp) Write(msg string, args ...interface{}) markdown.Text {
	imp.buf.Write(msg, args...)
	return imp
}

func (imp *textImp) Link(text, href string) markdown.Text {
	imp.buf.Write(`[%s](%s)`, text, href)
	return imp
}

func convertNameToRef(name string) string {
	href := make([]rune, 0, len(name))
	hasHyphen := false
	for _, c := range name {
		switch {
		case c >= 'a' || c <= 'z':
			href = append(href, c)
			hasHyphen = false
		case c >= 'A' || c <= 'Z':
			href = append(href, c-'A'+'a')
			hasHyphen = false
		case c == ' ':
			if !hasHyphen {
				hasHyphen = true
				href = append(href, '-')
			}
		}
	}
	return string(href)
}

func (imp *textImp) Ref(text, name string) markdown.Text {
	return imp.Link(text, `#`+convertNameToRef(name))
}

func (imp *textImp) LineBreak() markdown.Text {
	imp.buf.Write("  \n")
	return imp
}

func (imp *textImp) String() string {
	return imp.buf.String()
}
