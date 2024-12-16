package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/LiquidCats/ocelot/pkg/lexer"
)

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   []lexer.Token
		err        error
	}{
		{
			name:       "simple equation",
			expression: "1+1.2=2.2",
			expected: []lexer.Token{
				{
					Type:     lexer.TokenScalarInt,
					Value:    "1",
					Position: 0,
				},
				{
					Type:     lexer.TokenOperatorAdd,
					Value:    "+",
					Position: 1,
				},
				{
					Type:     lexer.TokenScalarFloat,
					Value:    "1.2",
					Position: 2,
				},
				{
					Type:     lexer.TokenOperatorIs,
					Value:    "=",
					Position: 5,
				},
				{
					Type:     lexer.TokenScalarFloat,
					Value:    "2.2",
					Position: 6,
				},
			},
		},
		{
			name:       "complex lexical comparison",
			expression: `(animal is "cat" or animal is "dog") and color is "black"`,
			expected: []lexer.Token{
				{
					Type:     lexer.TokenParenthesisRoundLeft,
					Value:    "(",
					Position: 0,
				},
				{
					Type:     lexer.TokenVariable,
					Value:    "animal",
					Position: 5,
				},
				{
					Type:     lexer.TokenOperatorIs,
					Value:    `is`,
					Position: 6,
				},
				{
					Type:     lexer.TokenScalarString,
					Value:    `"cat"`,
					Position: 6,
				},
				{
					Type:     lexer.TokenOperatorOr,
					Value:    `or`,
					Position: 6,
				},
				{
					Type:     lexer.TokenVariable,
					Value:    `animal`,
					Position: 6,
				},
				{
					Type:     lexer.TokenOperatorIs,
					Value:    `is`,
					Position: 6,
				},
				{
					Type:     lexer.TokenScalarString,
					Value:    `"dog"`,
					Position: 6,
				},
				{
					Type:     lexer.TokenParenthesisRoundRight,
					Value:    ")",
					Position: 0,
				},
				{
					Type:     lexer.TokenOperatorAnd,
					Value:    "and",
					Position: 0,
				},
				{
					Type:     lexer.TokenVariable,
					Value:    "color",
					Position: 0,
				},
				{
					Type:     lexer.TokenOperatorIs,
					Value:    `is`,
					Position: 6,
				},
				{
					Type:     lexer.TokenScalarString,
					Value:    `"black"`,
					Position: 6,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.NewLexer(lexer.CreateDictionary())

			tokens, err := lex.Tokenize(tt.expression)

			if tt.err != nil {
				require.ErrorAs(t, err, &tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, tokens)
			}
		})
	}
}
