package parser

import (
	"github.com/kaputi/sindar/ast"
	"github.com/kaputi/sindar/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}
