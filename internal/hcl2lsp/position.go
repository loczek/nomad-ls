package hcl2lsp

import (
	"unicode/utf8"

	"github.com/hashicorp/hcl/v2"
	"go.lsp.dev/protocol"
)

func Position(pos protocol.Position, src []byte) hcl.Pos {
	runes := []rune(string(src))

	var runeIndex uint
	var line uint
	var bytesCount uint

	for line < uint(pos.Line) && runeIndex < uint(len(runes)) {
		if runes[runeIndex] == '\n' {
			line += 1
		}
		bytesCount += uint(utf8.RuneLen(runes[runeIndex]))
		runeIndex += 1
	}

	var j uint

	for j < uint(pos.Character) && runeIndex < uint(len(runes)) {
		bytesCount += uint(utf8.RuneLen(runes[runeIndex]))
		runeIndex += 1
		j += 1
	}

	return hcl.Pos{
		Line:   int(pos.Line),
		Column: int(pos.Character),
		Byte:   int(bytesCount),
	}
}
