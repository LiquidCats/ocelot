package main

import (
	"ocelot-ruler/pkg/lexer"
	"ocelot-ruler/pkg/parser"
)

func main() {
	var lexerInstance lexer.ILexer
	var parserInstance parser.IParser
	var dictionary lexer.TokenDictionary

	dictionary = lexer.CreateDictionary()
	
	lexerInstance = lexer.NewLexer(dictionary)
	parserInstance = parser.NewParser(lexerInstance)

	//fmt.Println(tokens)
	//fmt.Println(err)
	//fmt.Println(lex)
}
