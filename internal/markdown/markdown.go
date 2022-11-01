package markdown

type (
	// Factory is used for creating new instances of markdown parts.
	Factory interface {

		// StringBuffer creates a new StringBuffer instance.
		StringBuffer() StringBuffer

		// Markdown creates a new root Markdown object with the given title.
		Markdown(title string) Markdown

		// Text creates a new object for storing markdown text.
		Text() Text

		// Mermaid creates a new Mermaid diagram.
		Mermaid() Mermaid

		// Class creates a new Class for a mermaid diagram.
		Class(name string) Class
	}

	// Stringer is an interface for any object which can return a string.
	Stringer interface {

		// String gets the string to write for this object.
		String() string
	}

	// StringBuffer is a buffer for storing and adding to a text string.
	StringBuffer interface {
		Stringer

		// Empty indicates if this buffer has not had any text written in it yet.
		Empty() bool

		// Write concatenates a formatted string to the end of the buffer.
		// This returns this instance of the string buffer to chain writes together.
		Write(msg string, args ...any) StringBuffer
	}

	// Markdown is a tool for writing markdown files programmatically.
	Markdown interface {
		Stringer

		// Section adds a section into the document and add the section to the index.
		Section(name string)

		// Subsection adds a subsection into the document and add the subsection
		// to the index. This should only be added after a section has already been added.
		Subsection(name string)

		// Par adds a new paragraph to the document under the current section/subsection.
		Par() Text

		// Mermaid add a mermaid diagram to the document under the current section/subsection.
		Mermaid() Mermaid
	}

	// Text is a buffer of markdown text usually for a paragraph.
	Text interface {
		Stringer

		// Bold writes bold text to the end of the buffer.
		// This returns this instance to chain text writing.
		Bold(msg string, args ...any) Text

		// Code writes code text to the end of the buffer.
		// This returns this instance to chain text writing.
		Code(msg string, args ...any) Text

		// Code writes plain text to the end of the buffer.
		// This returns this instance to chain text writing.
		Write(msg string, args ...any) Text

		// Link writes an external link to the end of the buffer.
		// This returns this instance to chain text writing.
		Link(text, href string) Text

		// Ref writes a link to a section or subsection to the buffer.
		// The given name is the name of section or subsection to reference.
		// This returns this instance to chain text writing.
		Ref(text, name string) Text

		// LineBreak writes a line break to the buffer.
		LineBreak() Text

		// HorizontalLine writes a horizontal line for dividing ideas.
		HorizontalLine() Text
	}

	// Mermaid is a tool for creating a mermaid class diagrams.
	Mermaid interface {
		Stringer

		// Has determines if a class, interface, or enum has been added to the diagram.
		Has(name string) bool

		// Class adds or gets a class with the given name.
		Class(name string) Class

		// Interface adds or gets a class with the given name.
		// If the class is new then it will be labelled as an interface.
		Interface(name string) Class

		// Enum add a class to represent an enumerator with he given name and values.
		// If a class by that name already exists then
		// it will not be changed and false will be returned.
		Enum(name string, values ...string) bool
	}

	// Class is a class for the mermaid diagram.
	Class interface {
		Stringer

		// AddEntry adds the given entry text to a line of the class body.
		AddEntry(entry string)

		// AddMember adds a member definition with the given name and type.
		AddMember(name, typeName string)

		// ConnectTo add a connection from this class to the class with the given name.
		ConnectTo(name string)
	}
)
