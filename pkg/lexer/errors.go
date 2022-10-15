package lexer

import "fmt"

type LexicalError struct {
	pos int
}

func (e LexicalError) Error() string {
	return fmt.Sprintf("Lexical issues was detected on position [%d]", e.pos)
}
