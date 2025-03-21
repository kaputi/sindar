package parser

import (
	"github.com/kaputi/sindar/ast"
	"github.com/kaputi/sindar/token"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}
