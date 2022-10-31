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

	// Operators
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
)

func initializeAsserter(regexpString string) *regexp.Regexp {
	rx := fmt.Sprintf("^%s", regexpString)
	r, _ := regexp.Compile(rx)

	return r
}

func CreateDictionary() TokenDictionary {
	return map[TokenType]*regexp.Regexp{
		TokenIntend:                 initializeAsserter("[\\s\\n\\t\\r]"),
		TokenVariable:               initializeAsserter("[a-z0-9]{1,80}"),
		TokenParenthesisSquareLeft:  initializeAsserter("\\["),
		TokenParenthesisSquareRight: initializeAsserter("\\]"),
		TokenScalarString:           initializeAsserter("\"([a-z0-9]{1,255}?)\""),
		TokenScalarInt:              initializeAsserter("[-]?[0-9]{1,10}"),
		TokenScalarFloat:            initializeAsserter("[-]?([0-9]{0,10}[.])?[0-9]{1,8}"),
		TokenScalarBool:             initializeAsserter("true|false"),
		TokenScalarNull:             initializeAsserter("null"),
		TokenOperatorIs:             initializeAsserter("is|="),
		TokenOperatorNot:            initializeAsserter("!=|not"),
		TokenOperatorIn:             initializeAsserter("in"),
		TokenOperatorNin:            initializeAsserter("nin"),
		TokenOperatorAnd:            initializeAsserter("and"),
		TokenOperatorOr:             initializeAsserter("or"),
		TokenOperatorXor:            initializeAsserter("xor"),
		TokenOperatorGt:             initializeAsserter(">"),
		TokenOperatorGte:            initializeAsserter(">="),
		TokenOperatorLt:             initializeAsserter("<"),
		TokenOperatorLte:            initializeAsserter("<="),
		TokenDot:                    initializeAsserter("\\."),
		TokenComma:                  initializeAsserter(","),
		TokenParenthesisRoundLeft:   initializeAsserter("\\("),
		TokenParenthesisRoundRight:  initializeAsserter("\\)"),
	}
}
