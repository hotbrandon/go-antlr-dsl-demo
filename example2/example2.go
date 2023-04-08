package main

import (
	"antlr-dsl-demo/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type calcListener struct {
	*parser.BaseCalcListener
	parser *parser.CalcParser 
}

// add a SetParser method to initialize the parser field
func (s *calcListener) SetParser(p *parser.CalcParser) { 
	s.parser = p
}

func (s *calcListener) EnterMulDiv(ctx *parser.MulDivContext) {
	// ruleName := ctx.GetRuleContext().GetRuleIndex()
	
	// fmt.Printf("Entering rule %s\n", s.parser.GetRuleNames()[ruleName])

}
func (s *calcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	ruleName := ctx.GetRuleContext().GetRuleIndex()	
	fmt.Printf("Entering rule %s\n", s.parser.GetRuleNames()[ruleName])

	if _, ok := ctx.(*parser.MulDivContext); ok {
        fmt.Printf("Entering context EnterMulDiv\n")
	}
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	parser := parser.NewCalcParser(stream)
	listener := &calcListener{}
	listener.SetParser(parser)
	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Start())
}