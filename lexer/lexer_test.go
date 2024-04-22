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
		if !assert.Equalf(t, expectedToken, l.NextToken(), "wrong token @ test[%d]", index) {
			assert.FailNow(t, "failed token equality check")
		}
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
		token.New(token.Equal, "=", 0),
		token.New(token.Plus, "+", 1),
		token.New(token.OpenParens, "(", 2),
		token.New(token.CloseParens, ")", 3),
		token.New(token.OpenCurlyBrackets, "{", 4),
		token.New(token.CloseCurlyBrackets, "}", 5),
		token.New(token.Comma, ",", 6),
		token.New(token.Semicolon, ";", 7),
		token.New(token.EndOfFile, "\x00", 8),
	}

	runTest(t, input, tests)
}

func TestChannelNextChars(t *testing.T) {
	input := `=+(){},;`
	tests := []token.Token{
		token.New(token.Equal, "=", 0),
		token.New(token.Plus, "+", 1),
		token.New(token.OpenParens, "(", 2),
		token.New(token.CloseParens, ")", 3),
		token.New(token.OpenCurlyBrackets, "{", 4),
		token.New(token.CloseCurlyBrackets, "}", 5),
		token.New(token.Comma, ",", 6),
		token.New(token.Semicolon, ";", 7),
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
		token.New(token.Let, "let", 0),
		token.New(token.Space, " ", 3),
		token.New(token.Identifier, "five", 4),
		token.New(token.Space, " ", 8),
		token.New(token.Equal, "=", 9),
		token.New(token.Space, " ", 10),
		token.New(token.Integer, "5", 11),
		token.New(token.Semicolon, ";", 12),
		token.New(token.LF, "\n", 13),
		token.New(token.Let, "let", 14),
		token.New(token.Space, " ", 17),
		token.New(token.Identifier, "ten", 18),
		token.New(token.Space, " ", 21),
		token.New(token.Equal, "=", 22),
		token.New(token.Space, " ", 23),
		token.New(token.Integer, "10", 24),
		token.New(token.Semicolon, ";", 26),
		token.New(token.LF, "\n", 27),
		token.New(token.Function, "function", 28),
		token.New(token.Space, " ", 36),
		token.New(token.Identifier, "add", 37),
		token.New(token.OpenParens, "(", 40),
		token.New(token.Identifier, "x", 41),
		token.New(token.Comma, ",", 42),
		token.New(token.Space, " ", 43),
		token.New(token.Identifier, "y", 44),
		token.New(token.CloseParens, ")", 45),
		token.New(token.Space, " ", 46),
		token.New(token.OpenCurlyBrackets, "{", 47),
		token.New(token.LF, "\n", 48),
		token.New(token.Tab, "\t", 49),
		token.New(token.Identifier, "x", 50),
		token.New(token.Space, " ", 51),
		token.New(token.Plus, "+", 52),
		token.New(token.Space, " ", 53),
		token.New(token.Identifier, "y", 54),
		token.New(token.Semicolon, ";", 55),
		token.New(token.LF, "\n", 56),
		token.New(token.CloseCurlyBrackets, "}", 57),
		token.New(token.LF, "\n", 58),
		token.New(token.Const, "const", 59),
		token.New(token.Space, " ", 64),
		token.New(token.Identifier, "result", 65),
		token.New(token.Space, " ", 71),
		token.New(token.Equal, "=", 72),
		token.New(token.Space, " ", 73),
		token.New(token.Identifier, "add", 74),
		token.New(token.OpenParens, "(", 77),
		token.New(token.Identifier, "five", 78),
		token.New(token.Comma, ",", 82),
		token.New(token.Space, " ", 83),
		token.New(token.Identifier, "ten", 84),
		token.New(token.CloseParens, ")", 87),
		token.New(token.Semicolon, ";", 88),
		token.New(token.EndOfFile, "\x00", 89),
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
		token.New(token.Let, "let", 0),
		token.New(token.Identifier, "five", 4),
		token.New(token.Equal, "=", 9),
		token.New(token.Integer, "5", 11),
		token.New(token.Semicolon, ";", 12),
		token.New(token.Let, "let", 14),
		token.New(token.Identifier, "ten", 18),
		token.New(token.Equal, "=", 22),
		token.New(token.Integer, "10", 24),
		token.New(token.Semicolon, ";", 26),
		token.New(token.Function, "function", 28),
		token.New(token.Identifier, "add", 37),
		token.New(token.OpenParens, "(", 40),
		token.New(token.Identifier, "x", 41),
		token.New(token.Comma, ",", 42),
		token.New(token.Identifier, "y", 44),
		token.New(token.CloseParens, ")", 45),
		token.New(token.OpenCurlyBrackets, "{", 47),
		token.New(token.Identifier, "x", 50),
		token.New(token.Plus, "+", 52),
		token.New(token.Identifier, "y", 54),
		token.New(token.Semicolon, ";", 55),
		token.New(token.CloseCurlyBrackets, "}", 57),
		token.New(token.Const, "const", 59),
		token.New(token.Identifier, "result", 65),
		token.New(token.Equal, "=", 72),
		token.New(token.Identifier, "add", 74),
		token.New(token.OpenParens, "(", 77),
		token.New(token.Identifier, "five", 78),
		token.New(token.Comma, ",", 82),
		token.New(token.Identifier, "ten", 84),
		token.New(token.CloseParens, ")", 87),
		token.New(token.Semicolon, ";", 88),
	}

	runChannelTest(t, input, tests)
}

func TestNextNewCharsToken(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;
`
	tests := []token.Token{
		token.New(token.Bang, "!", 0),
		token.New(token.Minus, "-", 1),
		token.New(token.Slash, "/", 2),
		token.New(token.Asterisk, "*", 3),
		token.New(token.Integer, "5", 4),
		token.New(token.Semicolon, ";", 5),
		token.New(token.LF, "\n", 6),
		token.New(token.Tab, "\t", 7),
		token.New(token.Integer, "5", 8),
		token.New(token.Space, " ", 9),
		token.New(token.LessThan, "<", 10),
		token.New(token.Space, " ", 11),
		token.New(token.Integer, "10", 12),
		token.New(token.Space, " ", 14),
		token.New(token.GreaterThan, ">", 15),
		token.New(token.Space, " ", 16),
		token.New(token.Integer, "5", 17),
		token.New(token.Semicolon, ";", 18),
		token.New(token.LF, "\n", 19),
		token.New(token.EndOfFile, "\x00", 20),
	}
	runTest(t, input, tests)
}

func TestChannelNextNewCharsToken(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;
`
	tests := []token.Token{
		token.New(token.Bang, "!", 0),
		token.New(token.Minus, "-", 1),
		token.New(token.Slash, "/", 2),
		token.New(token.Asterisk, "*", 3),
		token.New(token.Integer, "5", 4),
		token.New(token.Semicolon, ";", 5),
		token.New(token.Integer, "5", 8),
		token.New(token.LessThan, "<", 10),
		token.New(token.Integer, "10", 12),
		token.New(token.GreaterThan, ">", 15),
		token.New(token.Integer, "5", 17),
		token.New(token.Semicolon, ";", 18),
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
		token.New(token.If, "if", 0),
		token.New(token.Space, " ", 2),
		token.New(token.OpenParens, "(", 3),
		token.New(token.Integer, "5", 4),
		token.New(token.Space, " ", 5),
		token.New(token.LessThan, "<", 6),
		token.New(token.Space, " ", 7),
		token.New(token.Integer, "10", 8),
		token.New(token.CloseParens, ")", 10),
		token.New(token.Space, " ", 11),
		token.New(token.OpenCurlyBrackets, "{", 12),
		token.New(token.LF, "\n", 13),
		token.New(token.Tab, "\t", 14),
		token.New(token.Return, "return", 15),
		token.New(token.Space, " ", 21),
		token.New(token.True, "true", 22),
		token.New(token.Semicolon, ";", 26),
		token.New(token.LF, "\n", 27),
		token.New(token.CloseCurlyBrackets, "}", 28),
		token.New(token.Space, " ", 29),
		token.New(token.Else, "else", 30),
		token.New(token.Space, " ", 34),
		token.New(token.OpenCurlyBrackets, "{", 35),
		token.New(token.LF, "\n", 36),
		token.New(token.Tab, "\t", 37),
		token.New(token.Return, "return", 38),
		token.New(token.Space, " ", 44),
		token.New(token.False, "false", 45),
		token.New(token.Semicolon, ";", 50),
		token.New(token.LF, "\n", 51),
		token.New(token.CloseCurlyBrackets, "}", 52),
		token.New(token.LF, "\n", 53),
		token.New(token.EndOfFile, "\x00", 54),
	}
	runTest(t, input, tests)
}

func TestNextEQToken(t *testing.T) {
	input := `10 == 10;
	10 != 9;`
	tests := []token.Token{
		token.New(token.Integer, "10", 0),
		token.New(token.Space, " ", 2),
		token.New(token.DoubleEqual, "==", 3),
		token.New(token.Space, " ", 5),
		token.New(token.Integer, "10", 6),
		token.New(token.Semicolon, ";", 8),
		token.New(token.LF, "\n", 9),
		token.New(token.Tab, "\t", 10),
		token.New(token.Integer, "10", 11),
		token.New(token.Space, " ", 13),
		token.New(token.NotEqual, "!=", 14),
		token.New(token.Space, " ", 16),
		token.New(token.Integer, "9", 17),
		token.New(token.Semicolon, ";", 18),
		token.New(token.EndOfFile, "\x00", 19),
	}
	runTest(t, input, tests)
}

func TestNextStringToken(t *testing.T) {
	input := `"foobar";
	"foo bar";`
	tests := []token.Token{
		token.New(token.String, "foobar", 0),
		token.New(token.Semicolon, ";", 8),
		token.New(token.LF, "\n", 9),
		token.New(token.Tab, "\t", 10),
		token.New(token.String, "foo bar", 11),
		token.New(token.Semicolon, ";", 20),
		token.New(token.EndOfFile, "\x00", 21),
	}
	runTest(t, input, tests)
}

func TestNextEscapedStringToken(t *testing.T) {
	input := `"foo\"bar"; "foobar\"";`
	tests := []token.Token{
		token.New(token.String, `foo\"bar`, 0),
		token.New(token.Semicolon, ";", 10),
		token.New(token.Space, " ", 11),
		token.New(token.String, `foobar\"`, 12),
		token.New(token.Semicolon, ";", 22),
	}
	runTest(t, input, tests)
}

func TestNextArrayToken(t *testing.T) {
	input := `[1, 2];`
	tests := []token.Token{
		token.New(token.OpenSquareBrackets, "[", 0),
		token.New(token.Integer, "1", 1),
		token.New(token.Comma, ",", 2),
		token.New(token.Space, " ", 3),
		token.New(token.Integer, "2", 4),
		token.New(token.CloseSquareBrackets, "]", 5),
		token.New(token.Semicolon, ";", 6),
		token.New(token.EndOfFile, "\x00", 7),
	}
	runTest(t, input, tests)
}

func TestNextHashToken(t *testing.T) {
	input := `{ "foo": "bar" };`
	tests := []token.Token{
		token.New(token.OpenCurlyBrackets, "{", 0),
		token.New(token.Space, " ", 1),
		token.New(token.String, "foo", 2),
		token.New(token.Colon, ":", 7),
		token.New(token.Space, " ", 8),
		token.New(token.String, "bar", 9),
		token.New(token.Space, " ", 14),
		token.New(token.CloseCurlyBrackets, "}", 15),
		token.New(token.Semicolon, ";", 16),
		token.New(token.EndOfFile, "\x00", 17),
	}
	runTest(t, input, tests)
}

func TestNextCSSUnitToken(t *testing.T) {
	input := `5rem; 6em; 18px;`
	tests := []token.Token{
		token.New(token.Integer, "5", 0),
		token.New(token.Rem, "rem", 1),
		token.New(token.Semicolon, ";", 4),
		token.New(token.Space, " ", 5),
		token.New(token.Integer, "6", 6),
		token.New(token.Em, "em", 7),
		token.New(token.Semicolon, ";", 9),
		token.New(token.Space, " ", 10),
		token.New(token.Integer, "18", 11),
		token.New(token.Pixel, "px", 13),
		token.New(token.Semicolon, ";", 15),
		token.New(token.EndOfFile, "\x00", 16),
	}
	runTest(t, input, tests)
}

func TestChannelNextCSSSelectorToken(t *testing.T) {
	input := `.some-classname > #some-id {
	background-color: transparent;
}`
	tests := []token.Token{
		token.New(token.Dot, ".", 0),
		token.New(token.Identifier, "some-classname", 1),
		token.New(token.GreaterThan, ">", 16),
		token.New(token.NumberSign, "#", 18),
		token.New(token.Identifier, "some-id", 19),
		token.New(token.OpenCurlyBrackets, "{", 27),
		token.New(token.Identifier, "background-color", 30),
		token.New(token.Colon, ":", 46),
		token.New(token.Identifier, "transparent", 48),
		token.New(token.Semicolon, ";", 59),
		token.New(token.CloseCurlyBrackets, "}", 61),
	}
	runChannelTest(t, input, tests)
}

func TestChannelNextCSSNestedSelectorToken(t *testing.T) {
	input := `.some-classname {
	& > #some-id {
		background-color: transparent;
	}
}`
	tests := []token.Token{
		token.New(token.Dot, ".", 0),
		token.New(token.Identifier, "some-classname", 1),
		token.New(token.OpenCurlyBrackets, "{", 16),
		token.New(token.Ampersand, "&", 19),
		token.New(token.GreaterThan, ">", 21),
		token.New(token.NumberSign, "#", 23),
		token.New(token.Identifier, "some-id", 24),
		token.New(token.OpenCurlyBrackets, "{", 32),
		token.New(token.Identifier, "background-color", 36),
		token.New(token.Colon, ":", 52),
		token.New(token.Identifier, "transparent", 54),
		token.New(token.Semicolon, ";", 65),
		token.New(token.CloseCurlyBrackets, "}", 68),
		token.New(token.CloseCurlyBrackets, "}", 70),
	}
	runChannelTest(t, input, tests)
}
