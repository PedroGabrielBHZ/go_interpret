package fiandeira

import (
	"umbuLang/arbusteiro"
	"umbuLang/ciscador"
	"umbuLang/pontilha"
)

type Fiandeira struct {
	c *ciscador.Catador

	pontilhaAtual pontilha.Pontilha
	pontilhaEspia pontilha.Pontilha
}

func New(c *ciscador.Catador) *Fiandeira {
	f := &Fiandeira{c: c}

	f.pontilhaSeguinte()
	f.pontilhaSeguinte()

	return f
}

func (f *Fiandeira) pontilhaSeguinte() {
	f.pontilhaAtual = f.pontilhaEspia
	f.pontilhaEspia = f.c.NextPontilha()
}

func (f *Fiandeira) FiaCauso() *arbusteiro.Causo {
	return nil
}
