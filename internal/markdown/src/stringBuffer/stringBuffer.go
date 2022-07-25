package stringBuffer

import (
	"bytes"
	"fmt"

	"github.com/grant-nelson/DrOpenAPI/internal/markdown"
)

type stringBufferImp struct {
	buf *bytes.Buffer
}

func New() markdown.StringBuffer {
	return &stringBufferImp{
		buf: &bytes.Buffer{},
	}
}

func (imp *stringBufferImp) Empty() bool {
	return imp.buf.Len() > 0
}

func (imp *stringBufferImp) Write(msg string, args ...interface{}) markdown.StringBuffer {
	_, err := imp.buf.WriteString(fmt.Sprintf(msg, args...))
	if err != nil {
		panic(err)
	}
	return imp
}

func (imp *stringBufferImp) String() string {
	return imp.buf.String()
}
