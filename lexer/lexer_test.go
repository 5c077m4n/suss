package lexer

import (
	"suss/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tokenTypeLiteral struct {
	expectedType    string
	expectedLiteral string
}

func runTest(t *testing.T, input string, tests []tokenTypeLiteral) {
	t.Helper()

	l := New(input)
	for index, testCase := range tests {
		token := l.NextToken()

		assert.Equalf(t, testCase.expectedType, token.Type, "tests[%d] wrong token type", index)
		assert.Equalf(
			t,
			testCase.expectedLiteral,
			token.Literal,
			"tests[%d] wrong token literal",
			index,
		)
	}
}

func TestNextChars(t *testing.T) {
	input := `=+(){},;`
	tests := []tokenTypeLiteral{
		{token.Equal, "="},
		{token.Plus, "+"},
		{token.OpenParens, "("},
		{token.CloseParens, ")"},
		{token.OpenCurlyBrackets, "{"},
		{token.CloseCurlyBrackets, "}"},
		{token.Comma, ","},
		{token.Semicolon, ";"},
		{token.EndOfFile, ""},
	}

	runTest(t, input, tests)
}

func TestNextBasicToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
function add(x, y) {
	x + y;
}
const result = add(five, ten);`
	tests := []tokenTypeLiteral{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Equal, "="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Equal, "="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Function, "function"},
		{token.Identifier, "add"},
		{token.OpenParens, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.CloseParens, ")"},
		{token.OpenCurlyBrackets, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.CloseCurlyBrackets, "}"},
		{token.Const, "const"},
		{token.Identifier, "result"},
		{token.Equal, "="},
		{token.Identifier, "add"},
		{token.OpenParens, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.CloseParens, ")"},
		{token.Semicolon, ";"},
		{token.EndOfFile, ""},
	}
	runTest(t, input, tests)
}

func TestNextNewCharsToken(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;
`
	tests := []tokenTypeLiteral{
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.GreaterThan, ">"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.EndOfFile, ""},
	}
	runTest(t, input, tests)
}

func TestNextIfElseToken(t *testing.T) {
	input := `if (5 < 10) {
	return true;
} else {
	return false;
}
`
	tests := []tokenTypeLiteral{
		{token.If, "if"},
		{token.OpenParens, "("},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.CloseParens, ")"},
		{token.OpenCurlyBrackets, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.CloseCurlyBrackets, "}"},
		{token.Else, "else"},
		{token.OpenCurlyBrackets, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.CloseCurlyBrackets, "}"},
		{token.EndOfFile, ""},
	}
	runTest(t, input, tests)
}

func TestNextEQToken(t *testing.T) {
	input := `10 == 10;
	10 != 9;
	`
	tests := []tokenTypeLiteral{
		{token.Integer, "10"},
		{token.DoubleEqual, "=="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Integer, "10"},
		{token.NotEqual, "!="},
		{token.Integer, "9"},
		{token.Semicolon, ";"},
		{token.EndOfFile, ""},
	}
	runTest(t, input, tests)
}

func TestNextStringToken(t *testing.T) {
	input := `
	"foobar";
	"foo bar";
	`
	tests := []tokenTypeLiteral{
		{token.String, "foobar"},
		{token.Semicolon, ";"},
		{token.String, "foo bar"},
		{token.Semicolon, ";"},
		{token.EndOfFile, ""},
	}
	runTest(t, input, tests)
}

func TestNextArrayToken(t *testing.T) {
	input := `[1, 2];`
	tests := []tokenTypeLiteral{
		{token.OpenSquareBrackets, "["},
		{token.Integer, "1"},
		{token.Comma, ","},
		{token.Integer, "2"},
		{token.CloseSquareBrackets, "]"},
		{token.Semicolon, ";"},
		{token.EndOfFile, ""},
	}
	runTest(t, input, tests)
}

func TestNextHashToken(t *testing.T) {
	input := `{ "foo": "bar" };`
	tests := []tokenTypeLiteral{
		{token.OpenCurlyBrackets, "{"},
		{token.String, "foo"},
		{token.Colon, ":"},
		{token.String, "bar"},
		{token.CloseCurlyBrackets, "}"},
		{token.Semicolon, ";"},
		{token.EndOfFile, ""},
	}
	runTest(t, input, tests)
}

func TestNextCSSUnitToken(t *testing.T) {
	input := `5rem; 6em; 18px;`
	tests := []tokenTypeLiteral{
		{token.Integer, "5"},
		{token.Rem, "rem"},
		{token.Semicolon, ";"},
		{token.Integer, "6"},
		{token.Em, "em"},
		{token.Semicolon, ";"},
		{token.Integer, "18"},
		{token.Pixel, "px"},
		{token.Semicolon, ";"},
	}
	runTest(t, input, tests)
}
