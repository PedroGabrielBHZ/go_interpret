package catador

import (
	"testing"

	"umbuLang/toco"
)

func TestNextToco(t *testing.T) {
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
		expectedType    toco.TocoTipo
		expectedLiteral string
	}{
		{toco.BOTE, "bote"},
		{toco.IDENT, "cinco"},
		{toco.ASSIGN, "="},
		{toco.INT, "5"},
		{toco.SEMICOLON, ";"},
		{toco.BOTE, "bote"},
		{toco.IDENT, "dez"},
		{toco.ASSIGN, "="},
		{toco.INT, "10"},
		{toco.SEMICOLON, ";"},
		{toco.BOTE, "bote"},
		{toco.IDENT, "soma"},
		{toco.ASSIGN, "="},
		{toco.MANDINGA, "mandinga"},
		{toco.LPAREN, "("},
		{toco.IDENT, "x"},
		{toco.COMMA, ","},
		{toco.IDENT, "y"},
		{toco.RPAREN, ")"},
		{toco.LBRACE, "{"},
		{toco.IDENT, "x"},
		{toco.PLUS, "+"},
		{toco.IDENT, "y"},
		{toco.SEMICOLON, ";"},
		{toco.RBRACE, "}"},
		{toco.SEMICOLON, ";"},
		{toco.BOTE, "bote"},
		{toco.IDENT, "resultado"},
		{toco.ASSIGN, "="},
		{toco.IDENT, "soma"},
		{toco.LPAREN, "("},
		{toco.IDENT, "cinco"},
		{toco.COMMA, ","},
		{toco.IDENT, "dez"},
		{toco.RPAREN, ")"},
		{toco.SEMICOLON, ";"},
		{toco.BANG, "!"},
		{toco.MINUS, "-"},
		{toco.SLASH, "/"},
		{toco.ASTERISK, "*"},
		{toco.INT, "5"},
		{toco.SEMICOLON, ";"},
		{toco.INT, "5"},
		{toco.LT, "<"},
		{toco.INT, "10"},
		{toco.GT, ">"},
		{toco.INT, "5"},
		{toco.SEMICOLON, ";"},
		{toco.AVERIGUE, "averigue"},
		{toco.LPAREN, "("},
		{toco.INT, "5"},
		{toco.LT, "<"},
		{toco.INT, "10"},
		{toco.RPAREN, ")"},
		{toco.LBRACE, "{"},
		{toco.RECEBA, "receba"},
		{toco.SINCERIDADE, "sinceridade"},
		{toco.SEMICOLON, ";"},
		{toco.RBRACE, "}"},
		{toco.DOUTRAFORMA, "doutraforma"},
		{toco.LBRACE, "{"},
		{toco.RECEBA, "receba"},
		{toco.FUXICO, "fuxico"},
		{toco.SEMICOLON, ";"},
		{toco.RBRACE, "}"},
		{toco.INT, "10"},
		{toco.EQ, "=="},
		{toco.INT, "10"},
		{toco.SEMICOLON, ";"},
		{toco.INT, "10"},
		{toco.NOT_EQ, "!="},
		{toco.INT, "9"},
		{toco.SEMICOLON, ";"},
		{toco.ESTIO, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToco()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tocotype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
