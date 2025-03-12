package laks

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		input []Token
		want  []Expr
	}{
		{
			name: "literal",
			input: []Token{
				{T_INT, "4"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
			want: []Expr{
				{Type: E_LITERAL, Value: 4},
			},
		},
		{
			name: "simplebin",
			input: []Token{
				{T_INT, "6"},
				{T_PLUS, "+"},
				{T_INT, "14"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
			want: []Expr{
				{
					Type: E_BOP,
					Operator: "+",
					Left: &Expr{Type: E_LITERAL, Value: 6},
					Right: &Expr{Type: E_LITERAL, Value: 14},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Parse(test.input)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestParsePrecendece(t *testing.T) {
	tests := []struct {
		name  string
		input []Token
		want  string
	}{
		{
			name: "simplebin",
			input: []Token{
				{T_INT, "6"},
				{T_PLUS, "+"},
				{T_INT, "14"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
			want: "(+ 6 14)",
		},
		{
			name: "precedence1",
			input: []Token{
				{T_INT, "1"},
				{T_PLUS, "+"},
				{T_INT, "2"},
				{T_MULT, "*"},
				{T_INT, "3"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
			want: "(+ 1 (* 2 3))",
		},
		{
			name: "precedence2",
			input: []Token{
				{T_INT, "1"},
				{T_MULT, "*"},
				{T_INT, "2"},
				{T_PLUS, "+"},
				{T_INT, "3"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
			want: "(+ (* 1 2) 3)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Parse(test.input)
			var sexprs []string
			for _, g := range got {
				sexprs = append(sexprs, g.Sexpr())
			}
			sexpr := strings.Join(sexprs, "\n")
			if diff := cmp.Diff(test.want, sexpr); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

