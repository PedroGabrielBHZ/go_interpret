package ast

import "umbuLang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (c *Program) TokenLiteral() string {
	if len(c.Statements) > 0 {
		return c.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identificador
	Value Expression
}

func (bs *LetStatement) statementNode()       {}
func (bs *LetStatement) TokenLiteral() string { return bs.Token.Literal }

type Identificador struct {
	Token token.Token
	Value string
}

func (i *Identificador) expressionNode()      {}
func (i *Identificador) TokenLiteral() string { return i.Token.Literal }
