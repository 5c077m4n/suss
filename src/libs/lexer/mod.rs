use super::token::Token;

pub struct Lexer {
	input: &'static [u8],
	position: usize,
	read_position: usize,
	c: Option<u8>,
}
impl Lexer {
	fn new(input: &'static [u8]) -> Self {
		Self {
			input,
			position: 0,
			read_position: 0,
			c: None,
		}
	}

	fn is_c_letter(&self) -> bool {
		self.c.is_some_and(|c| {
			c.is_ascii_lowercase() || c.is_ascii_uppercase() || c == b'-' || c == b'_'
		})
	}
	fn is_c_digit(&self) -> bool {
		self.c.is_some_and(|c| c.is_ascii_digit())
	}
	fn read_char(&mut self) {
		self.c = self.input.get(self.read_position).copied();
		self.position = self.read_position;
		self.read_position += 1;
	}
	fn peek_char(&self) -> Option<u8> {
		self.input.get(self.read_position).copied()
	}
	fn read_indentifier(&mut self) -> Vec<u8> {
		let pos = self.position + 1;
		self.read_char();

		while self.is_c_letter() {
			self.read_char();
		}

		self.input[pos..self.position].to_owned()
	}

	fn next_token(&mut self) -> Token {
		let token = match self.c.unwrap_or(0) {
			b'.' => Token::Dot,
			b',' => Token::Comma,
			b'#' => Token::Hash,
			b'@' => Token::At,
			b'&' => Token::Ampersand,
			b'*' => Token::Asterisk,
			b':' => Token::Colon,
			b';' => Token::Semicolon,
			b'=' => {
				if self.peek_char().is_some_and(|c| c == b'=') {
					self.read_char();
					Token::DoubleEqual
				} else {
					Token::Equal
				}
			}
			b'!' => {
				if self.peek_char().is_some_and(|c| c == b'=') {
					self.read_char();
					Token::NotEqual
				} else {
					Token::Bang
				}
			}
			b'>' => {
				if self.peek_char().is_some_and(|c| c == b'=') {
					self.read_char();
					Token::GreaterThanOrEqual
				} else {
					Token::GreaterThan
				}
			}
			b'<' => {
				if self.peek_char().is_some_and(|c| c == b'=') {
					self.read_char();
					Token::LessThanOrEqual
				} else {
					Token::LessThan
				}
			}
			b'(' => Token::OpenParens,
			b')' => Token::CloseParens,
			b'{' => Token::OpenCurlyBrackets,
			b'}' => Token::CloseCurlyBrackets,
			b'/' => {
				if self.peek_char().is_some_and(|c| c == b'/') {
					self.read_char();
					Token::DoubleSlash
				} else {
					Token::Slash
				}
			}
			0 => Token::EndOfFile,
			other => Token::Illegal(other.to_string()),
		};
		self.read_char();

		token
	}
}

#[cfg(test)]
mod test;
