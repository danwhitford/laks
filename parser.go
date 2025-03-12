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

func (expr Expr) Sexpr() (string, error) {
	switch(expr.Type) {
	case E_LITERAL:
		return fmt.Sprint(expr.Value), nil
	case E_BOP:
		l, err := expr.Left.Sexpr()
		if err != nil {
			return "", err
		}
		r, err := expr.Right.Sexpr()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("(%v %v %v)", expr.Operator, l, r), nil
	default:
		return "", fmt.Errorf("dont know how to sexpr '%v'", expr)
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

func (prsr *parser) read_token_checked(expected TokenType) (Token, error) {
	t := prsr.tokens[prsr.curr]
	if t.T != expected {
		return t, fmt.Errorf("token '%v' is not of type '%v'", t, expected)
	}
	prsr.curr++
	return t, nil
}

func (prsr *parser) read_literal() (Expr, error) {
	t := prsr.read_token()
	d, err := strconv.Atoi(t.Lexeme)
	if err != nil {
		return Expr{}, fmt.Errorf("error reading literal. %v", err)
	}
	return Expr{Type: E_LITERAL, Value: d}, nil
}

func (prsr *parser) read_expression() (Expr, error) {
	expr, err := prsr.read_binary_op_low()
	if err != nil {
		return expr, err
	}
	_, err = prsr.read_token_checked(T_SEMI)
	if err != nil {
		return expr, err
	}
	return expr, nil
}

func (prsr *parser) read_binary_op_high() (Expr, error) {
	expr , err := prsr.read_literal()
	if err != nil {
		return expr, err
	}

	for prsr.match(T_MULT, T_DIV) {
		t := prsr.read_token()
		operator := t.Lexeme
		right, err := prsr.read_literal()
		if err != nil {
			return expr, err
		}
		tmp := Expr{expr.Type, expr.Value, expr.Left, expr.Right, expr.Operator}		
		expr = Expr{Type: E_BOP, Left: &tmp, Right: &right, Operator: operator}
	}

	return expr, nil
}

func (prsr *parser) read_binary_op_low() (Expr, error) {
	expr, err := prsr.read_binary_op_high()
	if err != nil {
		return expr, err
	}

	for prsr.match(T_PLUS, T_MINUS) {
		t := prsr.read_token()
		operator := t.Lexeme
		right, err  := prsr.read_binary_op_high()
		if err != nil {
			return expr, err
		}
		tmp := Expr{expr.Type, expr.Value, expr.Left, expr.Right, expr.Operator}		
		expr = Expr{Type: E_BOP, Left: &tmp, Right: &right, Operator: operator}
	}

	return expr, nil
}

func (prsr *parser) match(tt... TokenType) bool {
	return slices.Contains(tt, prsr.peek().T)
}

func (prsr *parser) peek() Token {
	return prsr.tokens[prsr.curr]
}

func Parse(input []Token) ([]Expr, error) {
	prsr := parser{tokens: input, l: len(input)}

	for prsr.peek().T != T_EOF {
		expr, err := prsr.read_expression()
		if err != nil {
			return prsr.expressions, err
		}
		prsr.expressions = append(prsr.expressions, expr)
	}

	return prsr.expressions, nil
}
