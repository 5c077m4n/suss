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
	Hash                = "#"          // Hash `#`
	Percent             = "%"          // Percent `%`
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
	BackSlash           = `\`          // BackSlash `\`
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
	Pixel               = "px"         // Pixel `px` unit
	Em                  = "em"         // Em `em` unit
	Rem                 = "rem"        // Rem `rem` unit
	Vw                  = "vw"         // Vw `vw` unit
	Vh                  = "vh"         // Vh `vh` unit
)

// Token the token struct
type Token struct {
	Type          string
	Literal       string
	startPosition int
}

// IsWhitespace checks if the token is a whitespace one
func (t Token) IsWhitespace() bool {
	return t.Type == Space || t.Type == Tab || t.Type == CR || t.Type == LF
}

// IsEOF checks if the token is the last one
func (t Token) IsEOF() bool {
	return t.Type == EndOfFile
}

// Position get the start and end position of the token
func (t Token) Position() (int, int) {
	return t.startPosition, t.startPosition + len(t.Literal)
}

// New create a new token
func New[T string | byte](tokenType string, literal T, startPosition int) Token {
	return Token{
		Type:          tokenType,
		Literal:       string(literal),
		startPosition: startPosition,
	}
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
	Pixel:    Pixel,
	Em:       Em,
	Rem:      Rem,
	Vw:       Vw,
	Vh:       Vh,
}

// LookupIdentifier lookup keywords and fallback to identifier
func LookupIdentifier(ident string) string {
	if token, ok := keywords[ident]; ok {
		return token
	}

	return Identifier
}
