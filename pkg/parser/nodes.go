package parser

import "ocelot-ruler/pkg/lexer"

type Node struct {
	nodeType NodeType
	token    *lexer.Token
	value    []*Node
}

func (n *Node) AddNode(node *Node) {
	n.value = append(n.value, node)
}
