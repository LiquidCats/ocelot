package parser

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
