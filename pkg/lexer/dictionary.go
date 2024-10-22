package lexer

import (
	"fmt"
	"regexp"
)

const (
	TokenIntend TokenType = iota
	TokenDot
	TokenComma

	// Variables
	TokenVariable

	// Paranthesis
	TokenParenthesisSquareLeft
	TokenParenthesisSquareRight
	TokenParenthesisRoundLeft
	TokenParenthesisRoundRight

	// Scalars
	TokenScalarString
	TokenScalarInt
	TokenScalarFloat
	TokenScalarBool
	TokenScalarNull

	// Operators: Logical
	TokenOperatorIs
	TokenOperatorNot
	TokenOperatorIn
	TokenOperatorNin
	TokenOperatorAnd
	TokenOperatorOr
	TokenOperatorXor
	TokenOperatorGt
	TokenOperatorGte
	TokenOperatorLt
	TokenOperatorLte

	// Operators: Math
	TokenOperatorAdd
	TokenOperatorSub
	TokenOperatorDivide
	TokenOperatorMultiply
)

func (tt TokenType) String() string {
	switch tt {
	case TokenIntend:
		return "Intend"
	case TokenDot:
		return "Dot"
	case TokenComma:
		return "Comma"
	case TokenVariable:
		return "Variable"
	case TokenParenthesisSquareLeft:
		return "ParenthesisSquareLeft"
	case TokenParenthesisSquareRight:
		return "ParenthesisSquareRight"
	case TokenParenthesisRoundLeft:
		return "ParenthesisRoundLeft"
	case TokenParenthesisRoundRight:
		return "ParenthesisRoundRight"
	case TokenScalarString:
		return "ScalarString"
	case TokenScalarInt:
		return "ScalarInt"
	case TokenScalarFloat:
		return "ScalarFloat"
	case TokenScalarBool:
		return "ScalarBool"
	case TokenScalarNull:
		return "ScalarNull"
	case TokenOperatorIs:
		return "OperatorIs"
	case TokenOperatorNot:
		return "OperatorNot"
	case TokenOperatorIn:
		return "OperatorIn"
	case TokenOperatorNin:
		return "OperatorNin"
	case TokenOperatorAnd:
		return "OperatorAnd"
	case TokenOperatorOr:
		return "OperatorOr"
	case TokenOperatorXor:
		return "OperatorXor"
	case TokenOperatorGt:
		return "OperatorGt"
	case TokenOperatorGte:
		return "OperatorGte"
	case TokenOperatorLt:
		return "OperatorLt"
	case TokenOperatorLte:
		return "OperatorLte"
	case TokenOperatorAdd:
		return "OperatorAdd"
	case TokenOperatorSub:
		return "OperatorSub"
	case TokenOperatorDivide:
		return "OperatorDivide"
	case TokenOperatorMultiply:
		return "OperatorMultiply"
	default:
		return "Unknown"
	}
}

type TokenDictionary map[TokenType]*regexp.Regexp

func (d TokenDictionary) Add(key TokenType, re *regexp.Regexp) {
	if d[key] == nil {
		d[key] = re
	}
}

func initializeAsserter(regexpString string) *regexp.Regexp {
	rx := fmt.Sprintf("^%s", regexpString)
	return regexp.MustCompile(rx)
}

func CreateDictionary() TokenDictionary {
	return map[TokenType]*regexp.Regexp{
		TokenIntend:                 initializeAsserter(`[\s\n\t\r]`),
		TokenVariable:               initializeAsserter("[a-z][a-z0-9]{1,80}"),
		TokenParenthesisSquareLeft:  initializeAsserter(`\[`),
		TokenParenthesisSquareRight: initializeAsserter(`\]`),
		TokenScalarString:           initializeAsserter(`\"([a-z0-9]{0,255})\"`),
		TokenScalarInt:              initializeAsserter(`[0-9]{1,10}`),
		TokenScalarFloat:            initializeAsserter(`([0-9]{0,10}[.])[0-9]{1,8}`),
		TokenScalarBool:             initializeAsserter(`true|false`),
		TokenScalarNull:             initializeAsserter(`null`),
		TokenOperatorIs:             initializeAsserter(`is|=`),
		TokenOperatorNot:            initializeAsserter(`!=|not`),
		TokenOperatorIn:             initializeAsserter(`in`),
		TokenOperatorNin:            initializeAsserter(`nin`),
		TokenOperatorAnd:            initializeAsserter(`and`),
		TokenOperatorOr:             initializeAsserter(`or`),
		TokenOperatorXor:            initializeAsserter(`xor`),
		TokenOperatorGt:             initializeAsserter(`>`),
		TokenOperatorGte:            initializeAsserter(`>=`),
		TokenOperatorLt:             initializeAsserter(`<`),
		TokenOperatorLte:            initializeAsserter(`<=`),
		TokenDot:                    initializeAsserter(`[\.]`),
		TokenComma:                  initializeAsserter(`[\,]`),
		TokenParenthesisRoundLeft:   initializeAsserter(`\(`),
		TokenParenthesisRoundRight:  initializeAsserter(`\)`),
		TokenOperatorAdd:            initializeAsserter(`[\+]`),
		TokenOperatorSub:            initializeAsserter(`[\-]`),
		TokenOperatorDivide:         initializeAsserter(`[\/]`),
		TokenOperatorMultiply:       initializeAsserter(`[\*]`),
	}
}
