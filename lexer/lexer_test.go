package lexer

import (
	"suss/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runTest(t *testing.T, input string, tests []token.Token) {
	t.Helper()

	l := New(input)

	for index, expectedToken := range tests {
		assert.Equalf(t, expectedToken, l.NextToken(), "wrong token @ test[%d]", index)
	}
}

func runChannelTest(t *testing.T, input string, expectedTokens []token.Token) {
	t.Helper()

	l := New(input)
	tokenChan := l.ToChannel()

	tokens := []token.Token{}
	for newToken := range tokenChan {
		tokens = append(tokens, newToken)
	}

	assert.Equal(t, expectedTokens, tokens)
}

func TestNextChars(t *testing.T) {
	input := `=+(){},;`
	tests := []token.Token{
		{Type: token.Equal, Literal: "="},
		{Type: token.Plus, Literal: "+"},
		{Type: token.OpenParens, Literal: "("},
		{Type: token.CloseParens, Literal: ")"},
		{Type: token.OpenCurlyBrackets, Literal: "{"},
		{Type: token.CloseCurlyBrackets, Literal: "}"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}

	runTest(t, input, tests)
}

func TestChannelNextChars(t *testing.T) {
	input := `=+(){},;`
	tests := []token.Token{
		{Type: token.Equal, Literal: "="},
		{Type: token.Plus, Literal: "+"},
		{Type: token.OpenParens, Literal: "("},
		{Type: token.CloseParens, Literal: ")"},
		{Type: token.OpenCurlyBrackets, Literal: "{"},
		{Type: token.CloseCurlyBrackets, Literal: "}"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Semicolon, Literal: ";"},
	}

	runChannelTest(t, input, tests)
}

func TestNextBasicToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
function add(x, y) {
	x + y;
}
const result = add(five, ten);`
	tests := []token.Token{
		{Type: token.Let, Literal: "let"},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "five"},
		{Type: token.Space, Literal: " "},
		{Type: token.Equal, Literal: "="},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Let, Literal: "let"},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "ten"},
		{Type: token.Space, Literal: " "},
		{Type: token.Equal, Literal: "="},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "10"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Function, Literal: "function"},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "add"},
		{Type: token.OpenParens, Literal: "("},
		{Type: token.Identifier, Literal: "x"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "y"},
		{Type: token.CloseParens, Literal: ")"},
		{Type: token.Space, Literal: " "},
		{Type: token.OpenCurlyBrackets, Literal: "{"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Tab, Literal: "\t"},
		{Type: token.Identifier, Literal: "x"},
		{Type: token.Space, Literal: " "},
		{Type: token.Plus, Literal: "+"},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "y"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.CloseCurlyBrackets, Literal: "}"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Const, Literal: "const"},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "result"},
		{Type: token.Space, Literal: " "},
		{Type: token.Equal, Literal: "="},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "add"},
		{Type: token.OpenParens, Literal: "("},
		{Type: token.Identifier, Literal: "five"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Space, Literal: " "},
		{Type: token.Identifier, Literal: "ten"},
		{Type: token.CloseParens, Literal: ")"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}

	runTest(t, input, tests)
}

func TestChannelNextBasicToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
function add(x, y) {
	x + y;
}
const result = add(five, ten);`
	tests := []token.Token{
		{Type: token.Let, Literal: "let"},
		{Type: token.Identifier, Literal: "five"},
		{Type: token.Equal, Literal: "="},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.Let, Literal: "let"},
		{Type: token.Identifier, Literal: "ten"},
		{Type: token.Equal, Literal: "="},
		{Type: token.Integer, Literal: "10"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.Function, Literal: "function"},
		{Type: token.Identifier, Literal: "add"},
		{Type: token.OpenParens, Literal: "("},
		{Type: token.Identifier, Literal: "x"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Identifier, Literal: "y"},
		{Type: token.CloseParens, Literal: ")"},
		{Type: token.OpenCurlyBrackets, Literal: "{"},
		{Type: token.Identifier, Literal: "x"},
		{Type: token.Plus, Literal: "+"},
		{Type: token.Identifier, Literal: "y"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.CloseCurlyBrackets, Literal: "}"},
		{Type: token.Const, Literal: "const"},
		{Type: token.Identifier, Literal: "result"},
		{Type: token.Equal, Literal: "="},
		{Type: token.Identifier, Literal: "add"},
		{Type: token.OpenParens, Literal: "("},
		{Type: token.Identifier, Literal: "five"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Identifier, Literal: "ten"},
		{Type: token.CloseParens, Literal: ")"},
		{Type: token.Semicolon, Literal: ";"},
	}

	runChannelTest(t, input, tests)
}

func TestNextNewCharsToken(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;
`
	tests := []token.Token{
		{Type: token.Bang, Literal: "!"},
		{Type: token.Minus, Literal: "-"},
		{Type: token.Slash, Literal: "/"},
		{Type: token.Asterisk, Literal: "*"},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Tab, Literal: "\t"},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Space, Literal: " "},
		{Type: token.LessThan, Literal: "<"},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "10"},
		{Type: token.Space, Literal: " "},
		{Type: token.GreaterThan, Literal: ">"},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}
	runTest(t, input, tests)
}

func TestChannelNextNewCharsToken(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;
`
	tests := []token.Token{
		{Type: token.Bang, Literal: "!"},
		{Type: token.Minus, Literal: "-"},
		{Type: token.Slash, Literal: "/"},
		{Type: token.Asterisk, Literal: "*"},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.Integer, Literal: "5"},
		{Type: token.LessThan, Literal: "<"},
		{Type: token.Integer, Literal: "10"},
		{Type: token.GreaterThan, Literal: ">"},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Semicolon, Literal: ";"},
	}
	runChannelTest(t, input, tests)
}

func TestNextIfElseToken(t *testing.T) {
	input := `if (5 < 10) {
	return true;
} else {
	return false;
}
`
	tests := []token.Token{
		{Type: token.If, Literal: "if"},
		{Type: token.Space, Literal: " "},
		{Type: token.OpenParens, Literal: "("},
		{Type: token.Integer, Literal: "5"},
		{Type: token.Space, Literal: " "},
		{Type: token.LessThan, Literal: "<"},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "10"},
		{Type: token.CloseParens, Literal: ")"},
		{Type: token.Space, Literal: " "},
		{Type: token.OpenCurlyBrackets, Literal: "{"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Tab, Literal: "\t"},
		{Type: token.Return, Literal: "return"},
		{Type: token.Space, Literal: " "},
		{Type: token.True, Literal: "true"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.CloseCurlyBrackets, Literal: "}"},
		{Type: token.Space, Literal: " "},
		{Type: token.Else, Literal: "else"},
		{Type: token.Space, Literal: " "},
		{Type: token.OpenCurlyBrackets, Literal: "{"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Tab, Literal: "\t"},
		{Type: token.Return, Literal: "return"},
		{Type: token.Space, Literal: " "},
		{Type: token.False, Literal: "false"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.CloseCurlyBrackets, Literal: "}"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}
	runTest(t, input, tests)
}

func TestNextEQToken(t *testing.T) {
	input := `10 == 10;
	10 != 9;`
	tests := []token.Token{
		{Type: token.Integer, Literal: "10"},
		{Type: token.Space, Literal: " "},
		{Type: token.DoubleEqual, Literal: "=="},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "10"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Tab, Literal: "\t"},
		{Type: token.Integer, Literal: "10"},
		{Type: token.Space, Literal: " "},
		{Type: token.NotEqual, Literal: "!="},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "9"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}
	runTest(t, input, tests)
}

func TestNextStringToken(t *testing.T) {
	input := `"foobar";
	"foo bar";`
	tests := []token.Token{
		{Type: token.String, Literal: "foobar"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.LF, Literal: "\n"},
		{Type: token.Tab, Literal: "\t"},
		{Type: token.String, Literal: "foo bar"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}
	runTest(t, input, tests)
}

func TestNextArrayToken(t *testing.T) {
	input := `[1, 2];`
	tests := []token.Token{
		{Type: token.OpenSquareBrackets, Literal: "["},
		{Type: token.Integer, Literal: "1"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "2"},
		{Type: token.CloseSquareBrackets, Literal: "]"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}
	runTest(t, input, tests)
}

func TestNextHashToken(t *testing.T) {
	input := `{ "foo": "bar" };`
	tests := []token.Token{
		{Type: token.OpenCurlyBrackets, Literal: "{"},
		{Type: token.Space, Literal: " "},
		{Type: token.String, Literal: "foo"},
		{Type: token.Colon, Literal: ":"},
		{Type: token.Space, Literal: " "},
		{Type: token.String, Literal: "bar"},
		{Type: token.Space, Literal: " "},
		{Type: token.CloseCurlyBrackets, Literal: "}"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}
	runTest(t, input, tests)
}

func TestNextCSSUnitToken(t *testing.T) {
	input := `5rem; 6em; 18px;`
	tests := []token.Token{
		{Type: token.Integer, Literal: "5"},
		{Type: token.Rem, Literal: "rem"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "6"},
		{Type: token.Em, Literal: "em"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.Space, Literal: " "},
		{Type: token.Integer, Literal: "18"},
		{Type: token.Pixel, Literal: "px"},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EndOfFile, Literal: "\x00"},
	}
	runTest(t, input, tests)
}
