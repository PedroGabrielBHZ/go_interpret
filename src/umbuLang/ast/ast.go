package ast

type Caroço interface {
	TocoLiteral() string
}

type Sentença interface {
	Caroço
	caroçoSentença()
}

type Calculação interface {
	Caroço
	caroçoCalculação()
}
