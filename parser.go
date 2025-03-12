package laks

import (
	"fmt"
	"strconv"
	"slices"
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

func (expr Expr) Sexpr() string {
	switch(expr.Type) {
	case E_LITERAL:
		return fmt.Sprint(expr.Value)
	case E_BOP:
		return fmt.Sprintf("(%v %v %v)", expr.Operator, expr.Left.Sexpr(), expr.Right.Sexpr())
	default:
		panic(fmt.Errorf("dont know how to sexpr '%v'", expr))
	}
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

func (prsr *parser) read_literal() Expr {
	t := prsr.read_token()
	d, err := strconv.Atoi(t.Lexeme)
	if err != nil {
		panic(err)
	}
	return Expr{Type: E_LITERAL, Value: d}	
}

func (prsr *parser) read_expression() Expr {
	expr := prsr.read_binary_op_low()
	prsr.read_token_checked(T_SEMI)
	return expr
}

func (prsr *parser) read_binary_op_high() Expr {
	expr := prsr.read_literal()

	for prsr.match(T_MULT, T_DIV) {
		t := prsr.read_token()
		operator := t.Lexeme
		right := prsr.read_literal()
		tmp := Expr{expr.Type, expr.Value, expr.Left, expr.Right, expr.Operator}		
		expr = Expr{Type: E_BOP, Left: &tmp, Right: &right, Operator: operator}
	}

	return expr
}

func (prsr *parser) read_binary_op_low() Expr {
	expr := prsr.read_binary_op_high()

	for prsr.match(T_PLUS, T_MINUS) {
		t := prsr.read_token()
		operator := t.Lexeme
		right := prsr.read_binary_op_high()
		tmp := Expr{expr.Type, expr.Value, expr.Left, expr.Right, expr.Operator}		
		expr = Expr{Type: E_BOP, Left: &tmp, Right: &right, Operator: operator}
	}

	return expr
}

func (prsr *parser) match(tt... TokenType) bool {
	return slices.Contains(tt, prsr.peek().T)
}

func (prsr *parser) peek() Token {
	return prsr.tokens[prsr.curr]
}

func Parse(input []Token) []Expr {
	prsr := parser{tokens: input, l: len(input)}

	for prsr.peek().T != T_EOF {
		expr := prsr.read_expression()
		prsr.expressions = append(prsr.expressions, expr)
	}

	return prsr.expressions
}
