package lexer

import (
	"umbuLang/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = rune(0)
		l.position = l.readPosition
		l.readPosition += 1
	} else {
		var size int
		l.ch, size = utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.position = l.readPosition
		l.readPosition += size
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.readTwoCharToken('=', token.EQ, token.ASSIGN)
	case '!':
		tok = l.readTwoCharToken('=', token.NOT_EQ, token.BANG)
	case ';', '(', ')', ',', '<', '>', '/', '*', '+', '-', '{', '}':
		tok = newToken(l.lookupSingleCharTokenType(l.ch), l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readTwoCharToken(expectedNextChar rune, twoCharType, singleCharType token.TokenType) token.Token {
	if l.peekChar() == expectedNextChar {
		ch := l.ch
		l.readChar()
		return token.Token{Type: twoCharType, Literal: string(ch) + string(l.ch)}
	}
	return newToken(singleCharType, l.ch)
}

func (l *Lexer) lookupSingleCharTokenType(ch rune) token.TokenType {
	switch ch {
	case ';':
		return token.SEMICOLON
	case '(':
		return token.LPAREN
	case ')':
		return token.RPAREN
	case ',':
		return token.COMMA
	case '<':
		return token.LT
	case '>':
		return token.GT
	case '/':
		return token.SLASH
	case '*':
		return token.ASTERISK
	case '+':
		return token.PLUS
	case '-':
		return token.MINUS
	case '{':
		return token.LBRACE
	case '}':
		return token.RBRACE
	default:
		return token.ILLEGAL
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		ch, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		return ch
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func newToken(TokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: TokenType, Literal: string(ch)}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
