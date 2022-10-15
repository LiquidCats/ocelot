package lexer

import (
	"strings"
)

func (l *Lexer) Tokens(text string) ([]*Token, error) {
	l.setText(text)

	for {
		shouldGoOn, err := l.next()

		if nil != err {
			return nil, err
		}

		if nil == err && false == shouldGoOn {
			break
		}
	}

	return l.tokens, nil
}

func (l *Lexer) next() (shouldGoOn bool, err error) {
	if l.position >= len(l.text) {
		return false, nil
	}

	for tokenType, r := range l.dictionary {
		stringIndex := r.FindStringIndex(l.text)

		if 0 != len(stringIndex) {
			if TokenIntend == tokenType { // skip all white spaces
				l.move(stringIndex)
				continue
			}

			token := &Token{
				Type:     tokenType,
				Position: l.position,
				Value:    l.text[stringIndex[0]:stringIndex[1]],
			}

			l.tokens = append(l.tokens, token)
			l.move(stringIndex)

			return true, nil
		}
	}

	return false, LexicalError{l.position}
}

func (l *Lexer) setText(text string) {
	l.text = strings.ToLower(text)
	l.position = 0
}

func (l *Lexer) move(pos []int) {
	l.position += pos[1]
	l.text = l.text[pos[1]:]
}

func NewLexer(dictionary TokenDictionary) *Lexer {
	return &Lexer{dictionary: dictionary}
}
