package parser

import "fmt"

type SyntaxErrorTokenNotFound struct {
	position int
}

func (e SyntaxErrorTokenNotFound) Error() string {
	return fmt.Sprintf("Syntaxt Error: Token not found at position [%d]", e.position)
}

type SyntaxErrorWrongToken struct {
	position int
}

func (e SyntaxErrorWrongToken) Error() string {
	return fmt.Sprintf("Syntaxt Error: wrong token prvided at position [%d]", e.position)
}
