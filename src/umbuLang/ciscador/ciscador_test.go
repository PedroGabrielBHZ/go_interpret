package ciscador

import (
	"testing"

	"umbuLang/pontilha"
)

func TestNextPontilha(t *testing.T) {
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
		expectedType    pontilha.PontilhaTipo
		expectedLiteral string
	}{
		{pontilha.BOTE, "bote"},
		{pontilha.IDENT, "cinco"},
		{pontilha.ASSIGN, "="},
		{pontilha.INT, "5"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.BOTE, "bote"},
		{pontilha.IDENT, "dez"},
		{pontilha.ASSIGN, "="},
		{pontilha.INT, "10"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.BOTE, "bote"},
		{pontilha.IDENT, "soma"},
		{pontilha.ASSIGN, "="},
		{pontilha.MANDINGA, "mandinga"},
		{pontilha.LPAREN, "("},
		{pontilha.IDENT, "x"},
		{pontilha.COMMA, ","},
		{pontilha.IDENT, "y"},
		{pontilha.RPAREN, ")"},
		{pontilha.LBRACE, "{"},
		{pontilha.IDENT, "x"},
		{pontilha.PLUS, "+"},
		{pontilha.IDENT, "y"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.RBRACE, "}"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.BOTE, "bote"},
		{pontilha.IDENT, "resultado"},
		{pontilha.ASSIGN, "="},
		{pontilha.IDENT, "soma"},
		{pontilha.LPAREN, "("},
		{pontilha.IDENT, "cinco"},
		{pontilha.COMMA, ","},
		{pontilha.IDENT, "dez"},
		{pontilha.RPAREN, ")"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.BANG, "!"},
		{pontilha.MINUS, "-"},
		{pontilha.SLASH, "/"},
		{pontilha.ASTERISK, "*"},
		{pontilha.INT, "5"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.INT, "5"},
		{pontilha.LT, "<"},
		{pontilha.INT, "10"},
		{pontilha.GT, ">"},
		{pontilha.INT, "5"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.AVERIGUE, "averigue"},
		{pontilha.LPAREN, "("},
		{pontilha.INT, "5"},
		{pontilha.LT, "<"},
		{pontilha.INT, "10"},
		{pontilha.RPAREN, ")"},
		{pontilha.LBRACE, "{"},
		{pontilha.RECEBA, "receba"},
		{pontilha.SINCERIDADE, "sinceridade"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.RBRACE, "}"},
		{pontilha.DOUTRAFORMA, "doutraforma"},
		{pontilha.LBRACE, "{"},
		{pontilha.RECEBA, "receba"},
		{pontilha.FUXICO, "fuxico"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.RBRACE, "}"},
		{pontilha.INT, "10"},
		{pontilha.EQ, "=="},
		{pontilha.INT, "10"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.INT, "10"},
		{pontilha.NOT_EQ, "!="},
		{pontilha.INT, "9"},
		{pontilha.SEMICOLON, ";"},
		{pontilha.ESTIO, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextPontilha()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - pontilhatype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
