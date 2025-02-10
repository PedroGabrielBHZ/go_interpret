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

// New creates a new Lexer instance with the provided input string.
// It initializes the Lexer with the given input and returns a pointer to it.
//
// Parameters:
//
//	input - the string input to be tokenized by the lexer.
//
// Returns:
//
//	A pointer to a Lexer instance initialized with the provided input.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character from the input and advances the lexer positions.
// If the end of the input is reached, it sets the current character to rune(0).
// It updates the current character, the current position, and the read position.
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

// NextToco retrieves the next token from the input and advances the lexer.
// It skips any whitespace and then checks the current character to determine
// the type of token to generate. It handles multi-character tokens for equality
// and inequality, single-character tokens for various operators and delimiters,
// and identifiers and integers. If the character is not recognized, it returns
// an MALFEITO token. The function returns the generated token.
func (l *Lexer) NextToco() token.Toco {
	var tok token.Toco

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.readTwoCharToco('=', token.EQ, token.ASSIGN)
	case '!':
		tok = l.readTwoCharToco('=', token.NOT_EQ, token.BANG)
	case ';', '(', ')', ',', '<', '>', '/', '*', '+', '-', '{', '}':
		tok = newToco(l.lookupSingleCharTocoTipo(l.ch), l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.ESTIO
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
			tok = newToco(token.MALFEITO, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readTwoCharToco reads a token that can either be a single character token or a two-character token.
// It checks if the next character matches the expectedNextChar. If it does, it returns a token of type twoCharType
// with the literal value being the combination of the current character and the next character.
// If the next character does not match the expectedNextChar, it returns a single character token of type singleCharType.
//
// Parameters:
// - expectedNextChar: The rune that is expected to follow the current character to form a two-character token.
// - twoCharType: The token type to return if the next character matches the expectedNextChar.
// - singleCharType: The token type to return if the next character does not match the expectedNextChar.
//
// Returns:
// - A token of type twoCharType if the next character matches the expectedNextChar, otherwise a token of type singleCharType.
func (l *Lexer) readTwoCharToco(expectedNextChar rune, twoCharType, singleCharType token.TocoTipo) token.Toco {
	if l.peekChar() == expectedNextChar {
		ch := l.ch
		l.readChar()
		return token.Toco{Type: twoCharType, Literal: string(ch) + string(l.ch)}
	}
	return newToco(singleCharType, l.ch)
}

func (l *Lexer) lookupSingleCharTocoTipo(ch rune) token.TocoTipo {
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
		return token.MALFEITO
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

func newToco(tokenType token.TocoTipo, ch rune) token.Toco {
	return token.Toco{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
