package lexer

import "github.com/kaputi/sindar/token"

type Lexer struct {
	input        string
	currPosition int // current position in input (current character)
	readPosition int // current reading position in input (after current character)
	Line         int
	Col          int
	ch           byte // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.Line = 1
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	tok := token.Token{
		Location: [2]int{l.Line, l.Col},
	}

	tok.Literal = string(l.ch)

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Literal = string(ch) + string(l.ch)
			tok.Type = token.EQ
		} else {
			tok.Type = token.ASSIGN
		}
	case '+':
		tok.Type = token.PLUS
	case '-':
		tok.Type = token.MINUS
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Literal = string(ch) + string(l.ch)
			tok.Type = token.NOT_EQ
		} else {
			tok.Type = token.BANG
		}
	case '/':
		tok.Type = token.SLASH
	case '*':
		tok.Type = token.ASTERISK
	case '<':
		tok.Type = token.LT
	case '>':
		tok.Type = token.GT
	case ';':
		tok.Type = token.SEMICOLON
	case ',':
		tok.Type = token.COMA
	case '{':
		tok.Type = token.LBRACE
	case '}':
		tok.Type = token.RBRACE
	case '(':
		tok.Type = token.LPAREN
	case ')':
		tok.Type = token.RPAREN
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok.Type = token.ILLEGAL
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.currPosition = l.readPosition
	l.readPosition++

	l.Col++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.Line++
			l.Col = 0
		}
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.currPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currPosition]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
