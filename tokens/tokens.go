// Package tokens holds all tokens to parse later
package tokens

import (
	"log"
	"slices"
)

const (
	Illegal            = "ILLEGAL" // Illegal character
	OpenParens         = "("       // OpenParens `(`
	CloseParens        = ")"       // CloseParens `)`
	OpenCurlyBrackets  = "{"       // OpenCurlyBrackets `{`
	CloseCurlyBrackets = "}"       // CloseCurlyBrackets `}`
	Colon              = ":"       // Colon `:`
	Semicolon          = ";"       // Semicolon `;`
	Dot                = "."       // Dot `.`
	Comma              = ","       // Comma `,`
	Ampersand          = "&"       // Ampersand `&`
	At                 = "@"       // At `@`
	Dollar             = "$"       // Dollar `$`
	NumberSign         = "#"       // NumberSign `#`
	Plus               = "+"       // Plus `+`
	Minus              = "-"       // Minus `-`
	Bang               = "!"       // Bang `!`
	Asterisk           = "*"       // Asterisk `*`
	EqEq               = "=="      // EqEq `==`
	Eq                 = "="       // Eq `=`
	Gt                 = ">"       // Gt `>`
	Lt                 = "<"       // Lt `<`
	DoubleSlash        = "//"      // DoubleSlash `//`
	Slash              = "/"       // Slash `/`
	Space              = " "       // Space ` `
	Tab                = "\t"      // Tab `\t`
	CR                 = "\r"      // CR carriage return
	LF                 = "\n"      // LF new line
	CRLF               = "\r\n"    // CRLF new line
)

// TokenTypes list of all tokens
var TokenTypes = []string{
	Illegal,
	OpenParens,
	CloseParens,
	OpenCurlyBrackets,
	CloseCurlyBrackets,
	Colon,
	Semicolon,
	Dot,
	Comma,
	Ampersand,
	At,
	Dollar,
	NumberSign,
	Plus,
	Minus,
	Bang,
	Asterisk,
	EqEq,
	Eq,
	Gt,
	Lt,
	DoubleSlash,
	Slash,
	Space,
	Tab,
	CR,
	LF,
	CRLF,
}

// Token the token struct
type Token struct {
	Type    string
	Literal string
}

// New create a new token
func New[T string | byte](tokenType string, literal T) Token {
	if !slices.Contains(TokenTypes, tokenType) {
		log.Fatalf("Invalid token type: %s\n", tokenType)
	}
	return Token{Type: tokenType, Literal: string(literal)}
}
