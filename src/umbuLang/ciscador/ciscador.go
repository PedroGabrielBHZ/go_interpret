package ciscador

import (
	"umbuLang/pontilha"
	"unicode"
	"unicode/utf8"
)

type Catador struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

// New creates a new Catador instance with the provided input string.
// It initializes the Catador with the given input and returns a pointer to it.
//
// Parameters:
//
//	input - the string input to be pontilhaized by the lexer.
//
// Returns:
//
//	A pointer to a Catador instance initialized with the provided input.
func New(input string) *Catador {
	l := &Catador{input: input}
	l.readChar()
	return l
}

// readChar reads the next character from the input and advances the lexer positions.
// If the end of the input is reached, it sets the current character to rune(0).
// It updates the current character, the current position, and the read position.
func (l *Catador) readChar() {
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

// NextPontilha retrieves the next pontilha from the input and advances the lexer.
// It skips any whitespace and then checks the current character to determine
// the type of pontilha to generate. It handles multi-character pontilhas for equality
// and inequality, single-character pontilhas for various operators and delimiters,
// and identifiers and integers. If the character is not recognized, it returns
// an MALFEITO pontilha. The function returns the generated pontilha.
func (l *Catador) NextPontilha() pontilha.Pontilha {
	var tok pontilha.Pontilha

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.readTwoCharPontilha('=', pontilha.EQ, pontilha.ASSIGN)
	case '!':
		tok = l.readTwoCharPontilha('=', pontilha.NOT_EQ, pontilha.BANG)
	case ';', '(', ')', ',', '<', '>', '/', '*', '+', '-', '{', '}':
		tok = newPontilha(l.lookupSingleCharPontilhaTipo(l.ch), l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = pontilha.ESTIO
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = pontilha.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = pontilha.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newPontilha(pontilha.MALFEITO, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readTwoCharPontilha reads a pontilha that can either be a single character pontilha or a two-character pontilha.
// It checks if the next character matches the expectedNextChar. If it does, it returns a pontilha of type twoCharType
// with the literal value being the combination of the current character and the next character.
// If the next character does not match the expectedNextChar, it returns a single character pontilha of type singleCharType.
//
// Parameters:
// - expectedNextChar: The rune that is expected to follow the current character to form a two-character pontilha.
// - twoCharType: The pontilha type to return if the next character matches the expectedNextChar.
// - singleCharType: The pontilha type to return if the next character does not match the expectedNextChar.
//
// Returns:
// - A pontilha of type twoCharType if the next character matches the expectedNextChar, otherwise a pontilha of type singleCharType.
func (l *Catador) readTwoCharPontilha(expectedNextChar rune, twoCharType, singleCharType pontilha.PontilhaTipo) pontilha.Pontilha {
	if l.peekChar() == expectedNextChar {
		ch := l.ch
		l.readChar()
		return pontilha.Pontilha{Type: twoCharType, Literal: string(ch) + string(l.ch)}
	}
	return newPontilha(singleCharType, l.ch)
}

func (l *Catador) lookupSingleCharPontilhaTipo(ch rune) pontilha.PontilhaTipo {
	switch ch {
	case ';':
		return pontilha.SEMICOLON
	case '(':
		return pontilha.LPAREN
	case ')':
		return pontilha.RPAREN
	case ',':
		return pontilha.COMMA
	case '<':
		return pontilha.LT
	case '>':
		return pontilha.GT
	case '/':
		return pontilha.SLASH
	case '*':
		return pontilha.ASTERISK
	case '+':
		return pontilha.PLUS
	case '-':
		return pontilha.MINUS
	case '{':
		return pontilha.LBRACE
	case '}':
		return pontilha.RBRACE
	default:
		return pontilha.MALFEITO
	}
}

func (l *Catador) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Catador) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		ch, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		return ch
	}
}

func (l *Catador) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Catador) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func newPontilha(pontilhaType pontilha.PontilhaTipo, ch rune) pontilha.Pontilha {
	return pontilha.Pontilha{Type: pontilhaType, Literal: string(ch)}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
