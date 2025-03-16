package lexer

import (
	"testing"

	"umbuLang/token"
)

func TestNextToken(t *testing.T) {
	input := `
		bote cinco = 5;
		bote dez = 10;

		bote soma = mandinga(x, y) { x + y; };

		bote resultado = soma(cinco, dez);
		!-/*5;

		5 < 10 > 5;

		averigue (5 < 10) {
			receba sinceridade;
		} doutraforma {
			receba fuxico;
		}

		10 == 10;
		10 != 9;
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "bote"},
		{token.IDENT, "cinco"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "bote"},
		{token.IDENT, "dez"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "bote"},
		{token.IDENT, "soma"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "mandinga"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "bote"},
		{token.IDENT, "resultado"},
		{token.ASSIGN, "="},
		{token.IDENT, "soma"},
		{token.LPAREN, "("},
		{token.IDENT, "cinco"},
		{token.COMMA, ","},
		{token.IDENT, "dez"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "averigue"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "receba"},
		{token.TRUE, "sinceridade"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "doutraforma"},
		{token.LBRACE, "{"},
		{token.RETURN, "receba"},
		{token.FALSE, "fuxico"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.ESTIO, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - Tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
