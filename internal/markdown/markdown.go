package markdown

type (
	Factory interface {
		StringBuffer() StringBuffer
		Markdown(title string) Markdown
		Text() Text
		Mermaid() Mermaid
		Class(name string) Class
	}

	Stringer interface {
		String() string
	}

	StringBuffer interface {
		Stringer
		Empty() bool
		Write(msg string, args ...interface{}) StringBuffer
	}

	Markdown interface {
		Stringer
		Section(name string)
		Subsection(name string)
		Par() Text
		Mermaid() Mermaid
	}

	Text interface {
		Stringer
		Bold(msg string, args ...interface{}) Text
		Code(msg string, args ...interface{}) Text
		Write(msg string, args ...interface{}) Text
		Link(text, href string) Text
		Ref(text, name string) Text
		LineBreak() Text
	}

	Mermaid interface {
		Stringer
		Has(name string) bool
		Class(name string) Class
		Interface(name string) Class
		Enum(name string, values ...string)
	}

	Class interface {
		Stringer
		AddEntry(entry string)
		AddMember(name, typeName string)
		ConnectTo(name string)
	}
)
