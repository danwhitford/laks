package laks

import (
	"fmt"
	"strconv"
)

type ExpressionType byte

const (
	E_LITERAL ExpressionType = iota
	E_BOP
)

type Expr struct {
	Type     ExpressionType
	Value    int
	Left     *Expr
	Right    *Expr
	Operator string
}

type parser struct {
	tokens      []Token
	curr        int
	l           int
	expressions []Expr
}

func (prsr *parser) read_token() Token {
	t := prsr.tokens[prsr.curr]
	prsr.curr++
	return t
}

func (prsr *parser) read_token_checked(expected TokenType) Token {
	t := prsr.tokens[prsr.curr]
	if t.T != expected {
		panic(fmt.Errorf("token '%v' is not of type '%v'", t, expected))
	}
	prsr.curr++
	return t
}

func (prsr *parser) read_literal() {
	t := prsr.read_token()
	d, err := strconv.Atoi(t.Lexeme)
	if err != nil {
		panic(err)
	}
	prsr.read_token_checked(T_SEMI)
	e := Expr{Type: E_LITERAL, Value: d}
	prsr.expressions = append(prsr.expressions, e
}

func Parse(input []Token) []Expr {
	prsr := parser{tokens: input, l: len(input)}

	for prsr.tokens[prsr.curr].T != T_EOF {
		switch prsr.tokens[prsr.curr].T {
		case T_INT:
			prsr.read_literal()
		}
	}

	return prsr.expressions
}
