package lexer

type TokenType int

type Token struct {
	Type     TokenType
	Value    string
	Position int
}
