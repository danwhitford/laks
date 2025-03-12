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
		tokens, err := laks.Lex(line)
		if err != nil {
			fmt.Printf("TOKERR '%s'\n", err)
			continue
		}
		exprs, err := laks.Parse(tokens)
		if err != nil {
			fmt.Printf("PRSERR '%s'\n", err)
			continue
		}
		for _, e := range exprs {
			s, err := e.Sexpr()
			if err != nil {
				fmt.Printf("SXPERR '%s'\n", err)				
			}
			fmt.Println("    |  " + s)
		}
	}

	fmt.Println("===")
}
