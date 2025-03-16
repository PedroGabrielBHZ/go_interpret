package arbusteiro

import "umbuLang/token"

type Caroço interface {
	TokenLiteral() string
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

func (c *Causo) TokenLiteral() string {
	if len(c.Sentenças) > 0 {
		return c.Sentenças[0].TokenLiteral()
	} else {
		return ""
	}
}

type BoteSentença struct {
	token.Token
	Nome  *Identificador
	Valor Calculação
}

func (bs *BoteSentença) caroçoSentença()      {}
func (bs *BoteSentença) TokenLiteral() string { return bs.Token.Literal }

type Identificador struct {
	Token token.Token
	Valor string
}

func (i *Identificador) caroçoCalculação()    {}
func (i *Identificador) TokenLiteral() string { return i.Token.Literal }
