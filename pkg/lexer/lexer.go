package lexer

import (
	"regexp"
)

type Lexer struct {
	dictionary map[TokenType]*regexp.Regexp
}

type Tokenizer interface {
	Tokenize(text string) ([]Token, error)
}

func (l *Lexer) Tokenize(source string) (tokens []Token, err error) {
	var position int

	for position < len(source) {
		matched := false

		for tokenType, handler := range l.dictionary {
			stringIndex := handler.FindStringIndex(source[position:])
			if stringIndex != nil && stringIndex[0] == 0 {
				if tokenType == TokenIntend { // skip all white spaces
					position += stringIndex[1]
					continue
				}

				token := Token{
					Type:     tokenType,
					Position: position,
					Value:    source[position+stringIndex[0] : position+stringIndex[1]],
				}

				tokens = append(tokens, token)
				position += stringIndex[1]
				matched = true

				break
			}
		}

		if !matched {
			err = LexicalError{position}
			return
		}
	}

	return
}

func NewLexer(dictionary TokenDictionary) *Lexer {
	return &Lexer{dictionary: dictionary}
}
