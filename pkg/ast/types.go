package ast

type NodeType int

const (
	StatementNode NodeType = iota
	BinaryOperationNode
	VariableNode
	ScalarNode
	GroupNode
	ArrayAccessNode  // not sure
	ObjectAccessNode // not sure
	FunctionNode     // not sure
)

func (n NodeType) String() string {
	switch n {
	case StatementNode:
		return "StatementNode"
	case BinaryOperationNode:
		return "BinaryOperationNode"
	case VariableNode:
		return "VariableNode"
	case ScalarNode:
		return "ScalarNode"
	case GroupNode:
		return "GroupNode"
	case ArrayAccessNode:
		return "ArrayAccessNode"
	case ObjectAccessNode:
		return "ObjectAccessNode"
	case FunctionNode:
		return "FunctionNode"
	default:
		return "Unknown"
	}
}
