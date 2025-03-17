package parser

import (
	"fmt"

	"github.com/kaputi/sindar/token"
)

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead %s", t, p.peekToken.Type, p.curToken.LocationString())
	p.errors = append(p.errors, msg)
}
