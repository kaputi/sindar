package parser

import (
	"fmt"

	"github.com/kaputi/sindar/ast"
	"github.com/kaputi/sindar/token"
)

func (p *Parser) noPrefixParseFnError(t token.Token) {
	msg := fmt.Sprintf("no prefix parse function for %s found. %s", t.Type, t.LocationString())
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}
