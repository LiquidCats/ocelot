package ast

import "github.com/LiquidCats/ocelot/pkg/lexer"

type Node struct {
	Type  NodeType
	Token *lexer.Token
	Left  *Node
	Right *Node
}
