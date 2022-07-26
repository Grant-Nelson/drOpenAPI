package text

import (
	"fmt"
	"strings"
	"testing"
)

func checkConvert(t *testing.T, name, expRef string) {
	ref := convertNameToRef(name)
	if ref != expRef {
		t.Error(strings.Join([]string{
			`Unexpected result from convertNameToRef:`,
			fmt.Sprintf(`  Input:    %q`, name),
			fmt.Sprintf(`  Result:   %q`, ref),
			fmt.Sprintf(`  Expected: %q`, expRef),
		}, "\n"))
	}
}

func Test_convertNameToRef(t *testing.T) {
	checkConvert(t, ``, `#`)
	checkConvert(t, `Main`, `#main`)
	checkConvert(t, `Hello World`, `#hello-world`)
	checkConvert(t, `Rich-Text`, `#rich-text`)
	checkConvert(t, `Rich--Text`, `#rich-text`)
	checkConvert(t, `Rich - Text`, `#rich-text`)
	checkConvert(t, `Oh #$%*@, Man!`, `#oh-man`)
	checkConvert(t, `-A-`, `#-a-`)
}
