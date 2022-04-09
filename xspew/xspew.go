package xspew

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"strings"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	spew.Config.Indent = strings.Repeat(" ", 4)
}

func MiddleStringLine(text string, width int, r rune) string {
	padLen := width - utf8.RuneCountInString(text) - 2 // <blank><text><blank>
	PadLeft, padRight := padLen/2, padLen-padLen/2
	return strings.Repeat(string(r), PadLeft) + " " + text + " " + strings.Repeat(string(r), padRight) + "\n"
}

func Dump(a ...interface{}) {
	var buf bytes.Buffer

	if bi, ok := debug.ReadBuildInfo(); ok {
		buf.WriteString(MiddleStringLine(bi.Path, 120, '~'))
	}

	spew.Fdump(&buf, a...)
	fmt.Println(buf.String())
}
