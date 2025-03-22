package parser

import (
	"umbuLang/ast"
	"umbuLang/lexer"
	"umbuLang/token"
)

type Parser struct {
	c *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(c *lexer.Lexer) *Parser {
	f := &Parser{c: c}

	f.nextToken()
	f.nextToken()

	return f
}

func (f *Parser) nextToken() {
	f.curToken = f.peekToken
	f.peekToken = f.c.NextToken()
}

func (f *Parser) ParseProgram() *ast.Program {
	return nil
}
