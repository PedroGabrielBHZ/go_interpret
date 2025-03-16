package fiandeira

import (
	"umbuLang/arbusteiro"
	"umbuLang/lexer"
	"umbuLang/token"
)

type Fiandeira struct {
	c *lexer.Lexer

	TokenAtual token.Token
	TokenEspia token.Token
}

func New(c *lexer.Lexer) *Fiandeira {
	f := &Fiandeira{c: c}

	f.TokenSeguinte()
	f.TokenSeguinte()

	return f
}

func (f *Fiandeira) TokenSeguinte() {
	f.TokenAtual = f.TokenEspia
	f.TokenEspia = f.c.NextToken()
}

func (f *Fiandeira) FiaCauso() *arbusteiro.Causo {
	return nil
}
