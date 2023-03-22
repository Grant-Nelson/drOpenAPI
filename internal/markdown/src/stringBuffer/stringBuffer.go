package stringBuffer

import (
	"bytes"
	"fmt"

	"github.com/Grant-Nelson/drOpenAPI/internal/markdown"
)

// stringBufferImp is the implementation of the StringBuffer interface.
type stringBufferImp struct {
	buf *bytes.Buffer
}

// New creates a new StringBuffer instance.
func New() markdown.StringBuffer {
	return &stringBufferImp{
		buf: &bytes.Buffer{},
	}
}

func (imp *stringBufferImp) Empty() bool {
	return imp.buf.Len() > 0
}

func (imp *stringBufferImp) Write(msg string, args ...any) markdown.StringBuffer {
	_, err := imp.buf.WriteString(fmt.Sprintf(msg, args...))
	if err != nil {
		panic(err)
	}
	return imp
}

func (imp *stringBufferImp) String() string {
	return imp.buf.String()
}
