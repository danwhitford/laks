package laks

import (
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
