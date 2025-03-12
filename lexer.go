package laks

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType byte

const (
	T_EOF TokenType = iota
	T_INT
	T_SEMI
	T_PLUS
	T_MULT
	T_MINUS
	T_DIV
)

type Token struct {
	T      TokenType
	Lexeme string
}

type lexer struct {
	curr   int
	l      int
	src    string
	tokens []Token
}

func (lxr *lexer) lex() {
	for lxr.canRead() {
		r := lxr.read()

		if unicode.IsSpace(r) {
			continue
		}

		if unicode.IsDigit(r) {
			lxr.read_int(r)
		} else if r == ';' {
			lxr.tokens = append(lxr.tokens, Token{T_SEMI, string(r)})
		} else if slices.Contains([]rune{'+', '-', '*', '/'}, r) {
			lxr.read_operator(r)
		} else {
			panic("dont understand '" + string(r) + "'")
		}
	}

	if lxr.curr == lxr.l {
		lxr.tokens = append(lxr.tokens, Token{T_EOF, ""})
	}
}

func (lxr *lexer) read() rune {
	r, s := utf8.DecodeRuneInString(lxr.src[lxr.curr:])
	if s == 0 {
		panic("could not read from source")
	}
	lxr.curr += s
	return r
}

func (lxr *lexer) canRead() bool {
	return lxr.curr < lxr.l
}

func (lxr *lexer) peek() rune {
	r, s := utf8.DecodeRuneInString(lxr.src[lxr.curr:])
	if s == 0 {
		panic("could not peek from source")
	}
	return r
}

func (lxr *lexer) read_int(start rune) {
	var sb strings.Builder
	sb.WriteRune(start)
	for lxr.canRead() {
		r := lxr.peek()
		if unicode.IsDigit(r) {
			lxr.read()
			sb.WriteRune(r)
		} else {
			break
		}
	}
	lxr.tokens = append(lxr.tokens, Token{T_INT, sb.String()})
}

func (lxr *lexer) read_operator(op rune) {
	switch op {
	case '+':
		lxr.tokens = append(lxr.tokens, Token{T_PLUS, string(op)})
	case '*':
		lxr.tokens = append(lxr.tokens, Token{T_MULT, string(op)})
	case '-':
		lxr.tokens = append(lxr.tokens, Token{T_MINUS, string(op)})
	case '/':
		lxr.tokens = append(lxr.tokens, Token{T_DIV, string(op)})
	default:
		panic(string(op) + " is not an operator")
	}
}

func Lex(src string) []Token {
	var lxr lexer

	lxr.src = src
	lxr.curr = 0
	lxr.l = utf8.RuneCountInString(src)

	lxr.lex()

	return lxr.tokens
}
