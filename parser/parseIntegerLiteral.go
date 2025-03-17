package parser

import (
	"fmt"
	"strconv"

	"github.com/kaputi/sindar/ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer. %v", p.curToken.Literal, p.curToken.LocationString())
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}
