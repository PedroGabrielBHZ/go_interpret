package ast

import "umbuLang/toco"

type Caroço interface {
	TocoLiteral() string
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

func (c *Causo) TocoLiteral() string {
	if len(c.Sentenças) > 0 {
		return c.Sentenças[0].TocoLiteral()
	} else {
		return ""
	}
}

type BoteSentença struct {
	toco.Toco
	Nome  *Identificador
	Valor Calculação
}

func (bs *BoteSentença) caroçoSentença()     {}
func (bs *BoteSentença) TocoLiteral() string { return bs.Toco.Literal }

type Identificador struct {
	Toco  toco.Toco
	Valor string
}

func (i *Identificador) caroçoCalculação()   {}
func (i *Identificador) TocoLiteral() string { return i.Toco.Literal }
