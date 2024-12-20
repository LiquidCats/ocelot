package parser

import (
	"github.com/LiquidCats/ocelot/pkg/ast"
	"github.com/LiquidCats/ocelot/pkg/lexer"
)

type Builder interface {
	BuildAST(source string) (*ast.Node, error)
}

type Parser struct {
	pos    int
	tokens []lexer.Token
}

func NewBuilder(t []lexer.Token) *Parser {
	return &Parser{
		tokens: t,
		pos:    0,
	}
}

func (p *Parser) Parse() (*ast.Node, error) {
	root := &ast.Node{Type: ast.StatementNode}

	// ["sad"] in ...
	// ctx =
	// ctx.field >=
	// "" =
	// 1 not
	// ctx.func() is

	for p.pos < len(p.tokens) {
		leftNode, err := p.parseVariableOrScalar()
		if nil != err {
			return nil, err
		}

		token := p.next()

		binaryNode := &ast.Node{Type: ast.BinaryOperationNode, Token: token}
		binaryNode.Left = leftNode
	}

	return root, nil
}

func (p *Parser) parseGroup() (*ast.Node, error) {
	token := p.next(lexer.TokenParenthesisSquareLeft)
	if nil == token {
		return nil, nil
	}

	groupNode := &ast.Node{Type: ast.GroupNode}

	_, err := p.parseListed(groupNode, lexer.TokenParenthesisSquareRight)
	if err != nil {
		return nil, err
	}

	return groupNode, nil
}

func (p *Parser) parseFunction(token *lexer.Token) (*ast.Node, error) {
	functionNode := &ast.Node{Type: ast.FunctionNode, Token: token}

	_, err := p.parseListed(functionNode, lexer.TokenParenthesisRoundRight)
	if err != nil {
		return nil, err
	}

	return functionNode, nil
}

func (p *Parser) parseArrayAccess(token *lexer.Token) (*ast.Node, error) {
	arrayNode := &ast.Node{Type: ast.ArrayAccessNode, Token: token}

	if left := p.next(lexer.TokenParenthesisSquareLeft); nil == left {
		return nil, SyntaxErrorTokenNotFound{p.pos}
	}

	index := p.next(lexer.TokenScalarInt)

	arrayNode.Left = &ast.Node{Type: ast.VariableNode, Token: index}

	if right := p.next(lexer.TokenParenthesisSquareRight); nil == right {
		return nil, SyntaxErrorTokenNotFound{p.pos}
	}

	return arrayNode, nil
}

func (p *Parser) parseObjectAccess(token *lexer.Token) (*ast.Node, error) {
	objectNode := &ast.Node{Type: ast.ObjectAccessNode, Token: token}

	for {
		objectKeyToken := p.next(lexer.TokenVariable, lexer.TokenScalarInt)

		if nil == objectKeyToken {
			return nil, SyntaxErrorTokenNotFound{p.pos}
		}

		objectNode.Left = &ast.Node{Type: ast.VariableNode, Token: token}

		if next := p.next(lexer.TokenDot); nil == next {
			break
		}
	}

	return objectNode, nil
}

func (p *Parser) parseListed(node *ast.Node, closeWith lexer.TokenType) (*ast.Node, error) {
	for {
		variable, err := p.parseVariableOrScalar()
		if err != nil {
			return nil, err
		}

		if nil == variable {
			break
		}

		node.Left = variable

		if commaToken := p.next(lexer.TokenComma); nil == commaToken {
			break
		}
	}

	if groupClose := p.next(closeWith); nil != groupClose {
		return node, nil
	}

	return nil, SyntaxErrorTokenNotFound{p.pos}
}

func (p *Parser) parseVariableOrScalar() (*ast.Node, error) {
	var node *ast.Node
	var err error

	node = p.parseScalar()

	if nil == node {
		node, err = p.parseVariable()
		if nil != err {
			return nil, err
		}
	}

	return node, nil
}

func (p *Parser) parseVariable() (*ast.Node, error) {
	token := p.next(lexer.TokenVariable)
	if nil == token {
		return nil, nil
	}

	if t := p.next(lexer.TokenParenthesisRoundLeft); nil != t {
		return p.parseFunction(token)
	}

	if t := p.next(lexer.TokenParenthesisSquareLeft); nil != t {
		return p.parseArrayAccess(token)
	}

	if t := p.next(lexer.TokenDot); nil != t {
		return p.parseObjectAccess(token)
	}

	return &ast.Node{Type: ast.VariableNode, Token: token}, nil
}

func (p *Parser) parseScalar() *ast.Node {
	token := p.next(lexer.TokenScalarString,
		lexer.TokenScalarInt,
		lexer.TokenScalarFloat,
		lexer.TokenScalarBool,
		lexer.TokenScalarNull)

	if nil == token {
		return nil
	}

	return &ast.Node{
		Type:  ast.ScalarNode,
		Token: token,
	}
}

func (p *Parser) next(tokenTypes ...lexer.TokenType) *lexer.Token {
	if p.pos < len(p.tokens) {
		currentToken := p.tokens[p.pos]

		if p.contains(tokenTypes, currentToken.Type) {
			p.pos++
			return &currentToken
		}
	}

	return nil
}

func (p *Parser) contains(s []lexer.TokenType, e lexer.TokenType) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}
