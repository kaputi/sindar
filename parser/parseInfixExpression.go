package parser

import "github.com/kaputi/sindar/ast"

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precendence := p.curPrecedence()
	p.nextToken()

	expression.Right = p.parseExpression(precendence)

	return expression
}
