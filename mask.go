package mask

import (
	"bufio"
	"os"
	"strings"

	"golang.org/x/term"
)

func Password(mask rune) string {
	var runes []rune
	builder := &strings.Builder{}
	previousTerminal, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func(fd int, terminal *term.State) {
		err := term.Restore(fd, terminal)
		if err != nil {
			panic(err)
		}
	}(int(os.Stdin.Fd()), previousTerminal)

	reader := bufio.NewReader(os.Stdin)
	for {
		if char, err := reader.ReadByte(); err == nil {
			runeChar := rune(char)
			if runeChar == '\r' {
				break
			} else if runeChar == '\b' || runeChar == 0x7f {
				if len(runes) != 0 {
					print("\b \b")
					runes = runes[:len(runes)-1]
				}
			} else {
				print(string(mask))
				runes = append(runes, runeChar)
			}
		}
	}
	println("")
	for _, r := range runes {
		builder.WriteString(string(r))
	}

	return builder.String()
}
