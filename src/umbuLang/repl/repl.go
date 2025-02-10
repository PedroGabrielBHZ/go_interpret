package repl

import (
	"bufio"
	"fmt"
	"io"
	"umbuLang/lexer"
	"umbuLang/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToco(); tok.Type != token.ESTIO; tok = l.NextToco() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
