package lexer

import "regexp"

type TokenType int

type TokenDictionary map[TokenType]*regexp.Regexp

type Token struct {
	Type     TokenType
	Value    string
	Position int
}

type Lexer struct {
	dictionary map[TokenType]*regexp.Regexp
	text       string
	position   int
	tokens     []*Token
}

type ILexer interface {
	Tokenize(text string) ([]*Token, error)
}
