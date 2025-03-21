package parser

import (
	"fmt"
	"testing"

	"github.com/kaputi/sindar/ast"
	"github.com/kaputi/sindar/lexer"
	"github.com/stretchr/testify/assert"
)

func TestLetStatements(t *testing.T) {

	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		program := createParseProgram(tt.input, t)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		testLetStatement(t, stmt, tt.expectedIdentifier)

		val := stmt.(*ast.LetStatement).Value
		testLiteralExpression(t, val, tt.expectedValue)
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
  `

	program := createParseProgram(input, t)

	assert.Equal(t, len(program.Statements), 3, "program.Statements does not contain 3 statements. got=%d", len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	program := createParseProgram(input, t)

	assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

	ident, ok := stmt.Expression.(*ast.Identifier)

	assert.True(t, ok, "exp not *ast.Identifier. got=%T", stmt.Expression)

	assert.Equal(t, ident.Value, "foobar", "ident.Value not %s. got=%s", "foobar", ident.Value)

	assert.Equal(t, ident.TokenLiteral(), "foobar", "ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5"

	program := createParseProgram(input, t)

	assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	assert.True(t, ok, "exp not *ast.IntegerLiteral. got=%T", stmt.Expression)

	assert.Equal(t, literal.Value, int64(5), "literal.Value not %d. got=%d", 5, literal.Value)

	assert.Equal(t, literal.TokenLiteral(), "5", "literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTest := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		// {"!true;", "!", 5},
		// {"!false;", "!", 5},
	}

	for _, tt := range prefixTest {
		program := createParseProgram(tt.input, t)

		assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.True(t, ok, "stmt is not ast.PrefixExpression. got=%T", stmt.Expression)

		assert.Equal(t, exp.Operator, tt.operator, "exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)

		testIntegerLiteral(t, exp.Right, tt.integerValue)
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range infixTests {
		program := createParseProgram(tt.input, t)

		assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

		testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	test := []struct {
		input    string
		expected string
	}{
		{"-a * b;", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b - c;", "((a + b) - c)"},
		{"a + b / c;", "(a + (b / c))"},
		{"a + b + c;", "((a + b) + c)"},
		{"a * b * c;", "((a * b) * c)"},
		{"a * b / c;", "((a * b) / c)"},
		{"a + b * c + d / e - f;", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},

		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == true", "((3 < 5) == true)"},
		{"!(true == true)", "(!(true == true))"},

		{"a + add(b * c) + d", "((a + add((b * c))) + d)"},
		{"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7* 8))", "add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))"},
		{"add(a + b + c * d / f + g)", "add((((a + b) + ((c * d) / f)) + g))"},

		// {"a * [1, 2, 3, 4][b * c] * d", "((a * ([1, 2, 3, 4][(b * c)])) * d)"},
		// {"add(a * b[2], b[1], 2 * [1, 2][1])", "add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))"},
	}

	for _, tt := range test {
		program := createParseProgram(tt.input, t)

		actual := program.String()

		assert.Equal(t, actual, tt.expected, "actual=%q, expected=%q", actual, tt.expected)
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`

	program := createParseProgram(input, t)

	assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

	exp, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(t, ok, "stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)

	testInfixExpression(t, exp.Condition, "x", "<", "y")

	assert.Equal(t, len(exp.Consequence.Statements), 1, "consequence is not 1 statements. got=%d", len(exp.Consequence.Statements))

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])

	testIdentifier(t, consequence.Expression, "x")

	assert.Nil(t, exp.Alternative, "exp.Alternative was not nil. got=%+v", exp.Alternative)
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`

	program := createParseProgram(input, t)

	assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

	exp, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(t, ok, "stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)

	testInfixExpression(t, exp.Condition, "x", "<", "y")

	assert.Equal(t, len(exp.Consequence.Statements), 1, "consequence is not 1 statements. got=%d", len(exp.Consequence.Statements))

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])

	testIdentifier(t, consequence.Expression, "x")

	assert.Equal(t, len(exp.Alternative.Statements), 1, "exp.Alternative.Statements does not contain 1 statements. got=%d", len(exp.Alternative.Statements))

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "Statements[0] is not ast.ExpressionStatement. got=%T", exp.Alternative.Statements[0])

	testIdentifier(t, alternative.Expression, "y")
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	program := createParseProgram(input, t)

	assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	assert.True(t, ok, "stmt.Expression is not ast.FunctionLiteral. got=%T", stmt.Expression)

	assert.Equal(t, len(function.Parameters), 2, "function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	assert.Equal(t, len(function.Body.Statements), 1, "function.Body.Statements has not 1 statements. got=%d\n", len(function.Body.Statements))

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "function body stmt is not %T. got=%T", &ast.ExpressionStatement{}, function.Body.Statements[0])

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		program := createParseProgram(tt.input, t)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)

		assert.Equal(t, len(function.Parameters), len(tt.expectedParams), "function literal parameters wrong. want %d, got=%d\n", len(tt.expectedParams), len(function.Parameters))

		for i, identifier := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], identifier)
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"

	program := createParseProgram(input, t)

	assert.Equal(t, len(program.Statements), 1, "program.Statements does not contain 1 statement. got=%d", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])

	exp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(t, ok, "stmt.Expression is not ast.CallExpression. got=%T", stmt.Expression)

	testIdentifier(t, exp.Function, "add")

	assert.Equal(t, len(exp.Arguments), 3, "wrong length of arguments. got=%d", len(exp.Arguments))

	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

// HELPERS /................................................................

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func createParseProgram(input string, t *testing.T) *ast.Program {
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	return program
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, s.TokenLiteral(), "let", "s.TokenLiteral not 'let', got=%q", s.TokenLiteral())

	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok, "s not *ast.LetStatement. got=%T", s)

	assert.Equal(t, letStmt.Name.Value, name, "letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)

	assert.Equal(t, letStmt.Name.TokenLiteral(), name, "s.Name not '%s'. got=%s", name, letStmt.Name)
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) {
	integ, ok := il.(*ast.IntegerLiteral)
	assert.True(t, ok, "il not *ast.IntegerLiteral. got=%T", il)

	assert.Equal(t, integ.Value, value, "integ.Value not %d. got=%d", value, integ.Value)

	assert.Equal(t, integ.TokenLiteral(), fmt.Sprintf("%d", value), "integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) {
	bo, ok := exp.(*ast.Boolean)
	assert.True(t, ok, "exp not *ast.Boolean. got=%T", exp)

	assert.Equal(t, bo.Value, value, "bo.Value not %t. got=%t", value, bo.Value)

	assert.Equal(t, bo.TokenLiteral(), fmt.Sprintf("%t", value), "bo.TokenLiteral not %t. got=%s", value, bo.TokenLiteral())
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) {
	identifier, ok := exp.(*ast.Identifier)
	assert.True(t, ok, "exp not *ast.Identifier. got=%T", exp)

	assert.Equal(t, identifier.Value, value, "identifier.Value not %s. got=%s", value, identifier.Value)

	assert.Equal(t, identifier.TokenLiteral(), value, "identifier.TokenLiteral not %s. got=%s", value, identifier.TokenLiteral())
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, exp, int64(v))
		return
	case int64:
		testIntegerLiteral(t, exp, v)
		return
	case string:
		testIdentifier(t, exp, v)
		return
	case bool:
		testBooleanLiteral(t, exp, v)
		return
	}

	t.Errorf("type of exp not handled. got=%T", exp)
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) {
	opExp, ok := exp.(*ast.InfixExpression)
	assert.True(t, ok, "exp is not ast.InfixExpression. got=%T(%s)", exp, exp)

	testLiteralExpression(t, opExp.Left, left)

	assert.Equal(t, opExp.Operator, operator, "exp.Operator is not '%s'. got=%q", operator, opExp.Operator)

	testLiteralExpression(t, opExp.Right, right)
}
