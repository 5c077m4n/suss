#[derive(Debug, PartialEq, Eq)]
pub enum Token {
	/// `.`
	Dot,
	/// `,`
	Comma,
	/// `#`
	Hash,
	/// `>`
	GreaterThan,
	/// `>=`
	GreaterThanOrEqual,
	/// `<`
	LessThan,
	/// `<=`
	LessThanOrEqual,
	/// `=`
	Equal,
	/// `==`
	DoubleEqual,
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
	/// `!`
	Bang,
	/// `!=`
	NotEqual,
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

#[derive(Debug, PartialEq, Eq)]
pub struct TokenData {
	pub token: Token,
	pub position: usize,
}
impl TokenData {
	pub fn new(token: Token, position: usize) -> Self {
		Self { token, position }
	}

	pub fn is_whitespace(&self) -> bool {
		matches!(
			self.token,
			Token::Space | Token::Tab | Token::CarriageReturn | Token::NewLine
		)
	}
	pub fn is_eof(&self) -> bool {
		self.token == Token::EndOfFile
	}
}
