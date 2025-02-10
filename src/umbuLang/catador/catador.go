package catador

import (
	"umbuLang/toco"
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
//	input - the string input to be tocoized by the lexer.
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

// NextToco retrieves the next toco from the input and advances the lexer.
// It skips any whitespace and then checks the current character to determine
// the type of toco to generate. It handles multi-character tocos for equality
// and inequality, single-character tocos for various operators and delimiters,
// and identifiers and integers. If the character is not recognized, it returns
// an MALFEITO toco. The function returns the generated toco.
func (l *Lexer) NextToco() toco.Toco {
	var tok toco.Toco

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.readTwoCharToco('=', toco.EQ, toco.ASSIGN)
	case '!':
		tok = l.readTwoCharToco('=', toco.NOT_EQ, toco.BANG)
	case ';', '(', ')', ',', '<', '>', '/', '*', '+', '-', '{', '}':
		tok = newToco(l.lookupSingleCharTocoTipo(l.ch), l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = toco.ESTIO
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = toco.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = toco.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToco(toco.MALFEITO, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readTwoCharToco reads a toco that can either be a single character toco or a two-character toco.
// It checks if the next character matches the expectedNextChar. If it does, it returns a toco of type twoCharType
// with the literal value being the combination of the current character and the next character.
// If the next character does not match the expectedNextChar, it returns a single character toco of type singleCharType.
//
// Parameters:
// - expectedNextChar: The rune that is expected to follow the current character to form a two-character toco.
// - twoCharType: The toco type to return if the next character matches the expectedNextChar.
// - singleCharType: The toco type to return if the next character does not match the expectedNextChar.
//
// Returns:
// - A toco of type twoCharType if the next character matches the expectedNextChar, otherwise a toco of type singleCharType.
func (l *Lexer) readTwoCharToco(expectedNextChar rune, twoCharType, singleCharType toco.TocoTipo) toco.Toco {
	if l.peekChar() == expectedNextChar {
		ch := l.ch
		l.readChar()
		return toco.Toco{Type: twoCharType, Literal: string(ch) + string(l.ch)}
	}
	return newToco(singleCharType, l.ch)
}

func (l *Lexer) lookupSingleCharTocoTipo(ch rune) toco.TocoTipo {
	switch ch {
	case ';':
		return toco.SEMICOLON
	case '(':
		return toco.LPAREN
	case ')':
		return toco.RPAREN
	case ',':
		return toco.COMMA
	case '<':
		return toco.LT
	case '>':
		return toco.GT
	case '/':
		return toco.SLASH
	case '*':
		return toco.ASTERISK
	case '+':
		return toco.PLUS
	case '-':
		return toco.MINUS
	case '{':
		return toco.LBRACE
	case '}':
		return toco.RBRACE
	default:
		return toco.MALFEITO
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

func newToco(tocoType toco.TocoTipo, ch rune) toco.Toco {
	return toco.Toco{Type: tocoType, Literal: string(ch)}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
