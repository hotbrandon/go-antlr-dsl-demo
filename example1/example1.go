package main

import (
	"antlr-dsl-demo/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

const Red = "\033[31m"

func main() {
	fmt.Println("--- lexer demo ---")
	// seup the input
	is := antlr.NewInputStream("1 + 2 * 3")
	// create the lexer
	lexer := parser.NewCalcLexer(is)

	// read all tokens
	for {
		t := lexer.NextToken()
		// fmt.Println("t=",t)
		// fmt.Println("token type=",t.GetTokenType())
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s%s (%q)%s\n", Red, 
			lexer.SymbolicNames[t.GetTokenType()], t.GetText(), "\033[0m")
	}
}