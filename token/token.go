// Package token holds all tokens to parse later
package token

const (
	Illegal             = "ILLEGAL"    // Illegal character
	Identifier          = "IDENTIFIER" // Identifier some ident
	EndOfFile           = "EOF"        // EndOfFile final token in a file
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
	NotEqual            = "!="         // NotEqual `!=`
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
	Else                = "else"       // Else the `else` keyword
	For                 = "for"        // For the `for` keyword
	Let                 = "let"        // Let the `let` keyword
	Const               = "const"      // Const the `const` keyword
	Function            = "function"   // Function the `function` keyword
	Return              = "return"     // Return the `return` keyword
	Throw               = "throw"      // Throw the `throw` keyword
	True                = "true"       // True the `true` keyword
	False               = "false"      // False the `false` keyword
	Integer             = "int"        // Integer int type
	String              = "string"     // String string type
)

// Token the token struct
type Token struct {
	Type    string
	Literal string
}

var keywords = map[string]string{
	If:       If,
	Else:     Else,
	For:      For,
	Let:      Let,
	Const:    Const,
	Function: Function,
	Throw:    Throw,
	Return:   Return,
	True:     True,
	False:    False,
}

// LookupIdentifier lookup keywords and fallback to identifier
func LookupIdentifier(ident string) string {
	if token, ok := keywords[ident]; ok {
		return token
	}

	return Identifier
}

// New create a new token
func New[T string | byte](tokenType string, literal T) Token {
	return Token{Type: tokenType, Literal: string(literal)}
}
