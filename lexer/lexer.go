// Package lexer the tokenizer
package lexer

import (
	"errors"
	"suss/token"
)

// Lexer struct
type Lexer struct {
	input         string
	position      int
	readPoisition int
	char          byte
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char == '-'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

// New lexer constructor
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() byte {
	if l.readPoisition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPoisition]
	}
	l.position = l.readPoisition
	l.readPoisition++

	return l.char
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() (string, error) {
	position := l.position + 1
	for {
		switch l.readChar() {
		case '"':
			return l.input[position:l.position], nil
		case 0:
			return "", errors.New("unfinished string")
		}
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPoisition <= len(l.input) {
		return l.input[l.readPoisition]
	}
	return 0
}

// NextToken get the next token
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch c := l.char; c {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.New(token.DoubleEqual, "==")
		} else {
			t = token.New(token.Equal, c)
		}
	case '+':
		t = token.New(token.Plus, c)
	case '-':
		t = token.New(token.Minus, c)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.New(token.NotEqual, "!=")
		} else {
			t = token.New(token.Bang, c)
		}
	case '/':
		t = token.New(token.Slash, c)
	case '%':
		t = token.New(token.Percent, c)
	case '#':
		t = token.New(token.NumberSign, c)
	case '*':
		t = token.New(token.Asterisk, c)
	case '<':
		t = token.New(token.LessThan, c)
	case '>':
		t = token.New(token.GreaterThan, c)
	case ';':
		t = token.New(token.Semicolon, c)
	case ':':
		t = token.New(token.Colon, c)
	case '(':
		t = token.New(token.OpenParens, c)
	case ')':
		t = token.New(token.CloseParens, c)
	case ',':
		t = token.New(token.Comma, c)
	case '[':
		t = token.New(token.OpenSquareBrackets, c)
	case ']':
		t = token.New(token.CloseSquareBrackets, c)
	case '{':
		t = token.New(token.OpenCurlyBrackets, c)
	case '}':
		t = token.New(token.CloseCurlyBrackets, c)
	case '"':
		stringLiteral, err := l.readString()
		if err != nil {
			panic(err)
		}

		t = token.New(token.String, stringLiteral)
	case '\x00':
		t = token.New(token.EndOfFile, c)
	case ' ':
		t = token.New(token.Space, c)
	case '\t':
		t = token.New(token.Tab, c)
	case '\r':
		t = token.New(token.CR, c)
	case '\n':
		t = token.New(token.LF, c)
	default:
		if isLetter(c) {
			ident := l.readIdentifier()
			return token.New(token.LookupIdentifier(ident), ident)
		} else if isDigit(c) {
			n := l.readNumber()
			return token.New(token.Integer, n)
		}
		t = token.New(token.Illegal, c)
	}

	l.readChar()
	return t
}

// ToChannel converts the lexer object to a (non-whitespace) token channel
func (l *Lexer) ToChannel() <-chan token.Token {
	ch := make(chan token.Token)

	go func() {
		defer close(ch)

		for t := l.NextToken(); !t.IsEOF(); t = l.NextToken() {
			if !t.IsWhitespace() {
				ch <- t
			}
		}
	}()

	return ch
}
