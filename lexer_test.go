package laks

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLex(t *testing.T) {
	tests  := []struct{
		src string
		want []Token
	} {
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
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			got := Lex(test.src)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

