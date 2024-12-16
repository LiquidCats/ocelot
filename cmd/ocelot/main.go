package main

import (
	"flag"
	"log"

	"github.com/LiquidCats/ocelot/pkg/lexer"
)

func main() {
	sourcePtr := flag.String("s", "", "code to compile")

	lex := lexer.NewLexer(lexer.CreateDictionary())

	_, err := lex.Tokenize(*sourcePtr)
	if err != nil {
		log.Fatal(err)
	}

}
