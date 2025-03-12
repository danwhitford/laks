package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danwhitford/laks"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := laks.Lex(line)
		exprs := laks.Parse(tokens)
		for _, e := range exprs {
			fmt.Println("    |  " + e.Sexpr())
		}
	}

	fmt.Println("===")
}
