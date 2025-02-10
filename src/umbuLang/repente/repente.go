package repente

import (
	"bufio"
	"fmt"
	"io"
	"umbuLang/catador"
	"umbuLang/toco"
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
		l := catador.New(line)

		for tok := l.NextToco(); tok.Type != toco.ESTIO; tok = l.NextToco() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
