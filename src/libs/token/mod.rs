#[derive(Debug)]
pub enum Token {
	/// `.`
	Dot,
	/// `,`
	Comma,
	/// `#`
	Hash,
	/// `>`
	GreaterThan,
	/// `<`
	LessThan,
	/// `=`
	Equal,
	/// `(`
	OpenParens,
	/// `)`
	CloseParens,
	/// `{`
	OpenCurlyBrackets,
	/// `}`
	CloseCurlyBrackets,
	/// `+`
	Plus,
	/// `-`
	Minus,
	/// `*`
	Asterisk,
	/// `/`
	Slash,
	/// `//`
	DoubleSlash,
	/// `:`
	Colon,
	/// `$`
	Dollar,
	/// `&`
	Ampersand,
	/// `@`
	At,
	/// `%`
	Percent,
	/// `_`
	Underscore,
	/// `'`
	Quote,
	/// `"`
	DoubleQuote,
	/// `;`
	Semicolon,
	/// `px`
	Pixel,
	/// `em`
	Em,
	/// `rem`
	Rem,
	/// ` `
	Space,
	/// `\t`
	Tab,
	/// `\r`
	CarriageReturn,
	/// `\n`
	NewLine,
	/// `EOF`
	EndOfFile,
	Identifier(String),
	Illegal(String),
}
