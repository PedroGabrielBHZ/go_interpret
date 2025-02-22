package repente

import (
	"bufio"
	"fmt"
	"io"
	"umbuLang/ciscador"
	"umbuLang/pontilha"
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
		l := ciscador.New(line)

		for tok := l.NextPontilha(); tok.Type != pontilha.ESTIO; tok = l.NextPontilha() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
