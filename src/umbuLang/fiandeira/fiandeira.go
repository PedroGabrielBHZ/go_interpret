package fiandeira

import (
	"umbuLang/arbusteiro"
	"umbuLang/catador"
	"umbuLang/toco"
)

type Fiandeira struct {
	c *catador.Lexer

	tocoAtual toco.Toco
	tocoEspia toco.Toco
}

func New(c *catador.Lexer) *Fiandeira {
	f := &Fiandeira{c: c}

	f.tocoSeguinte()
	f.tocoSeguinte()

	return f
}

func (f *Fiandeira) tocoSeguinte() {
	f.tocoAtual = f.tocoEspia
	f.tocoEspia = f.c.NextToco()
}

func (f *Fiandeira) FiaCauso() *arbusteiro.Causo {
	return nil
}
