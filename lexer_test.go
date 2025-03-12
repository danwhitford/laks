package laks

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLex(t *testing.T) {
	tests := []struct {
		src  string
		want []Token
	}{
		{
			src: "",
			want: []Token{
				{T_EOF, ""},
			},
		},
		{
			src: "1",
			want: []Token{
				{T_INT, "1"},
				{T_EOF, ""},
			},
		},
		{
			src: "14",
			want: []Token{
				{T_INT, "14"},
				{T_EOF, ""},
			},
		},
		{
			src: "14    41",
			want: []Token{
				{T_INT, "14"},
				{T_INT, "41"},
				{T_EOF, ""},
			},
		},
		{
			src: "42;",
			want: []Token{
				{T_INT, "42"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
		},
		{
			src: "1 + 7;",
			want: []Token{
				{T_INT, "1"},
				{T_PLUS, "+"},
				{T_INT, "7"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
		},
		{
			src: "1 + 7;\n18 *24;",
			want: []Token{
				{T_INT, "1"},
				{T_PLUS, "+"},
				{T_INT, "7"},
				{T_SEMI, ";"},
				{T_INT, "18"},
				{T_MULT, "*"},
				{T_INT, "24"},
				{T_SEMI, ";"},
				{T_EOF, ""},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			got, err := Lex(test.src)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
