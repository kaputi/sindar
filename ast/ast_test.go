package ast

import (
	"testing"

	"github.com/kaputi/sindar/token"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testString := "let myVar = anotherVar;"

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	programString := program.String()

	assert.Equal(t, programString, testString)
}
