package arbusteiro

import "umbuLang/pontilha"

type Caroço interface {
	PontilhaLiteral() string
}

type Sentença interface {
	Caroço
	caroçoSentença()
}

type Calculação interface {
	Caroço
	caroçoCalculação()
}

type Causo struct {
	Sentenças []Sentença
}

func (c *Causo) PontilhaLiteral() string {
	if len(c.Sentenças) > 0 {
		return c.Sentenças[0].PontilhaLiteral()
	} else {
		return ""
	}
}

type BoteSentença struct {
	pontilha.Pontilha
	Nome  *Identificador
	Valor Calculação
}

func (bs *BoteSentença) caroçoSentença()         {}
func (bs *BoteSentença) PontilhaLiteral() string { return bs.Pontilha.Literal }

type Identificador struct {
	Pontilha pontilha.Pontilha
	Valor    string
}

func (i *Identificador) caroçoCalculação()       {}
func (i *Identificador) PontilhaLiteral() string { return i.Pontilha.Literal }
