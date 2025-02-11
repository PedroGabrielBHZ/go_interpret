package pontilha

type PontilhaTipo string

type Pontilha struct {
	Type    PontilhaTipo
	Literal string
}

const (
	// Pontilhas especiais
	MALFEITO = "MALFEITO"
	ESTIO    = "ESTIO"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	MANDINGA    = "MANDINGA"
	BOTE        = "BOTE"
	SINCERIDADE = "SINCERIDADE"
	FUXICO      = "FUXICO"
	AVERIGUE    = "AVERIGUE"
	DOUTRAFORMA = "DOUTRAFORMA"
	RECEBA      = "RECEBA"
)

var keywords = map[string]PontilhaTipo{
	"mandinga":    MANDINGA,
	"bote":        BOTE,
	"sinceridade": SINCERIDADE,
	"fuxico":      FUXICO,
	"averigue":    AVERIGUE,
	"doutraforma": DOUTRAFORMA,
	"receba":      RECEBA,
}

func LookupIdent(ident string) PontilhaTipo {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
