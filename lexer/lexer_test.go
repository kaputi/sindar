package lexer

import (
	"testing"

	"github.com/kaputi/sindar/token"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5
5 < 10 > 5;

if (5 < 10) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType     token.TokenType
		expectedLiteral  string
		expectedLocation [2]int
		testNumber       int
	}{

		{token.LET, "let", [2]int{1, 1}, 0},
		{token.IDENT, "five", [2]int{1, 5}, 1},
		{token.ASSIGN, "=", [2]int{1, 10}, 2},
		{token.INT, "5", [2]int{1, 12}, 3},
		{token.SEMICOLON, ";", [2]int{1, 13}, 4},
		{token.LET, "let", [2]int{2, 1}, 5},
		{token.IDENT, "ten", [2]int{2, 5}, 6},
		{token.ASSIGN, "=", [2]int{2, 9}, 7},
		{token.INT, "10", [2]int{2, 11}, 8},
		{token.SEMICOLON, ";", [2]int{2, 13}, 9},
		{token.LET, "let", [2]int{4, 1}, 10},
		{token.IDENT, "add", [2]int{4, 5}, 11},
		{token.ASSIGN, "=", [2]int{4, 9}, 12},
		{token.FUNCTION, "fn", [2]int{4, 11}, 13},
		{token.LPAREN, "(", [2]int{4, 13}, 14},
		{token.IDENT, "x", [2]int{4, 14}, 15},
		{token.COMMA, ",", [2]int{4, 15}, 16},
		{token.IDENT, "y", [2]int{4, 17}, 17},
		{token.RPAREN, ")", [2]int{4, 18}, 18},
		{token.LBRACE, "{", [2]int{4, 20}, 19},
		{token.IDENT, "x", [2]int{5, 3}, 20},
		{token.PLUS, "+", [2]int{5, 5}, 21},
		{token.IDENT, "y", [2]int{5, 7}, 22},
		{token.SEMICOLON, ";", [2]int{5, 8}, 23},
		{token.RBRACE, "}", [2]int{6, 1}, 24},
		{token.SEMICOLON, ";", [2]int{6, 2}, 25},
		{token.LET, "let", [2]int{8, 1}, 26},
		{token.IDENT, "result", [2]int{8, 5}, 27},
		{token.ASSIGN, "=", [2]int{8, 12}, 28},
		{token.IDENT, "add", [2]int{8, 14}, 29},
		{token.LPAREN, "(", [2]int{8, 17}, 30},
		{token.IDENT, "five", [2]int{8, 18}, 31},
		{token.COMMA, ",", [2]int{8, 22}, 32},
		{token.IDENT, "ten", [2]int{8, 24}, 33},
		{token.RPAREN, ")", [2]int{8, 27}, 34},
		{token.SEMICOLON, ";", [2]int{8, 28}, 35},
		{token.BANG, "!", [2]int{9, 1}, 36},
		{token.MINUS, "-", [2]int{9, 2}, 37},
		{token.SLASH, "/", [2]int{9, 3}, 38},
		{token.ASTERISK, "*", [2]int{9, 4}, 39},
		{token.INT, "5", [2]int{9, 5}, 40},
		{token.INT, "5", [2]int{10, 1}, 41},
		{token.LT, "<", [2]int{10, 3}, 42},
		{token.INT, "10", [2]int{10, 5}, 43},
		{token.GT, ">", [2]int{10, 8}, 44},
		{token.INT, "5", [2]int{10, 10}, 45},
		{token.SEMICOLON, ";", [2]int{10, 11}, 46},
		{token.IF, "if", [2]int{12, 1}, 47},
		{token.LPAREN, "(", [2]int{12, 4}, 48},
		{token.INT, "5", [2]int{12, 5}, 49},
		{token.LT, "<", [2]int{12, 7}, 50},
		{token.INT, "10", [2]int{12, 9}, 51},
		{token.RPAREN, ")", [2]int{12, 11}, 52},
		{token.LBRACE, "{", [2]int{12, 13}, 53},
		{token.RETURN, "return", [2]int{13, 3}, 54},
		{token.TRUE, "true", [2]int{13, 10}, 55},
		{token.SEMICOLON, ";", [2]int{13, 14}, 56},
		{token.RBRACE, "}", [2]int{14, 1}, 57},
		{token.ELSE, "else", [2]int{14, 3}, 58},
		{token.LBRACE, "{", [2]int{14, 8}, 59},
		{token.RETURN, "return", [2]int{15, 3}, 60},
		{token.FALSE, "false", [2]int{15, 10}, 61},
		{token.SEMICOLON, ";", [2]int{15, 15}, 62},
		{token.RBRACE, "}", [2]int{16, 1}, 63},
		{token.INT, "10", [2]int{18, 1}, 64},
		{token.EQ, "==", [2]int{18, 4}, 65},
		{token.INT, "10", [2]int{18, 7}, 66},
		{token.SEMICOLON, ";", [2]int{18, 9}, 67},
		{token.INT, "10", [2]int{19, 1}, 68},
		{token.NOT_EQ, "!=", [2]int{19, 4}, 69},
		{token.INT, "9", [2]int{19, 7}, 70},
		{token.SEMICOLON, ";", [2]int{19, 8}, 71},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()
		assert.Equal(t, tt.expectedType, tok.Type, "test[%d] tokentype wrong for: %v", tt.testNumber, tt.expectedLiteral)
		assert.Equal(t, tt.expectedLiteral, tok.Literal, "test[%d] literal wrong for: %v", tt.testNumber, tt.expectedLiteral)
		assert.Equal(t, tt.expectedLocation, tok.Location, "test[%d] location wrong for: %v", tt.testNumber, tt.expectedLiteral)
	}
}
