// Package tokens holds all tokens to parse later
package tokens

const (
	Illegal             = "ILLEGAL"    // Illegal character
	Identifier          = "IDENTIFIER" // Identifier some ident
	OpenParens          = "("          // OpenParens `(`
	CloseParens         = ")"          // CloseParens `)`
	OpenCurlyBrackets   = "{"          // OpenCurlyBrackets `{`
	CloseCurlyBrackets  = "}"          // CloseCurlyBrackets `}`
	OpenSquareBrackets  = "["          // OpenSquareBrackets `[`
	CloseSquareBrackets = "]"          // CloseSquareBrackets `]`
	Colon               = ":"          // Colon `:`
	Semicolon           = ";"          // Semicolon `;`
	Dot                 = "."          // Dot `.`
	Comma               = ","          // Comma `,`
	Quote               = `'`          // Quote `'`
	DoubleQuote         = `"`          // DoubleQuote `"`
	Ampersand           = "&"          // Ampersand `&`
	Pipe                = "|"          // Pipe `|`
	At                  = "@"          // At `@`
	Dollar              = "$"          // Dollar `$`
	NumberSign          = "#"          // NumberSign `#`
	Plus                = "+"          // Plus `+`
	Minus               = "-"          // Minus `-`
	Bang                = "!"          // Bang `!`
	Asterisk            = "*"          // Asterisk `*`
	DoubleEqual         = "=="         // DoubleEqual `==`
	Equal               = "="          // Equal `=`
	GreaterThan         = ">"          // GreaterThan `>`
	LessThan            = "<"          // LessThan `<`
	DoubleSlash         = "//"         // DoubleSlash `//`
	Slash               = "/"          // Slash `/`
	Space               = " "          // Space ` `
	Tab                 = "\t"         // Tab `\t`
	CR                  = "\r"         // CR carriage return
	LF                  = "\n"         // LF new line
	If                  = "if"         // If the `if` keyword
	For                 = "for"        // For the `for` keyword
	Function            = "function"   // Function the `function` keyword
	Throw               = "throw"      // Throw the `throw` keyword
)

// Token the token struct
type Token struct {
	Type    string
	Literal string
}

var keywords = map[string]string{
	If:       If,
	For:      For,
	Function: Function,
	Throw:    Throw,
}

// LookupKeyword lookup keywords and fallback to identifier
func LookupKeyword(ident string) string {
	if token, ok := keywords[ident]; ok {
		return token
	}

	return Identifier
}

// New create a new token
func New[T string | byte](tokenType string, literal T) Token {
	return Token{Type: tokenType, Literal: string(literal)}
}
