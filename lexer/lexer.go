// Package lexer the tokenizer
package lexer

import (
	"errors"
	"log"

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
		case '\\':
			l.readChar()
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

	switch c, pos := l.char, l.position; c {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.New(token.DoubleEqual, "==", pos)
		} else {
			t = token.New(token.Equal, c, pos)
		}
	case '+':
		t = token.New(token.Plus, c, pos)
	case '-':
		t = token.New(token.Minus, c, pos)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.New(token.NotEqual, "!=", pos)
		} else {
			t = token.New(token.Bang, c, pos)
		}
	case '/':
		t = token.New(token.Slash, c, pos)
	case '%':
		t = token.New(token.Percent, c, pos)
	case '#':
		t = token.New(token.NumberSign, c, pos)
	case '&':
		t = token.New(token.Ampersand, c, pos)
	case '*':
		t = token.New(token.Asterisk, c, pos)
	case '<':
		t = token.New(token.LessThan, c, pos)
	case '>':
		t = token.New(token.GreaterThan, c, pos)
	case ';':
		t = token.New(token.Semicolon, c, pos)
	case ':':
		t = token.New(token.Colon, c, pos)
	case '(':
		t = token.New(token.OpenParens, c, pos)
	case ')':
		t = token.New(token.CloseParens, c, pos)
	case '.':
		t = token.New(token.Dot, c, pos)
	case ',':
		t = token.New(token.Comma, c, pos)
	case '[':
		t = token.New(token.OpenSquareBrackets, c, pos)
	case ']':
		t = token.New(token.CloseSquareBrackets, c, pos)
	case '{':
		t = token.New(token.OpenCurlyBrackets, c, pos)
	case '}':
		t = token.New(token.CloseCurlyBrackets, c, pos)
	case '"':
		stringLiteral, err := l.readString()
		if err != nil {
			log.Fatal(err)
		}

		t = token.New(token.String, stringLiteral, pos)
	case '\x00':
		t = token.New(token.EndOfFile, c, pos)
	case ' ':
		t = token.New(token.Space, c, pos)
	case '\t':
		t = token.New(token.Tab, c, pos)
	case '\r':
		t = token.New(token.CR, c, pos)
	case '\n':
		t = token.New(token.LF, c, pos)
	default:
		if isLetter(c) {
			ident := l.readIdentifier()
			return token.New(token.LookupIdentifier(ident), ident, pos)
		} else if isDigit(c) {
			n := l.readNumber()
			return token.New(token.Integer, n, pos)
		}
		t = token.New(token.Illegal, c, pos)
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
