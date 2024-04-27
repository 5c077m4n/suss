use super::token::{Token, TokenData};

pub struct Lexer {
	input: &'static [u8],
	position: usize,
	read_position: usize,
	c: Option<u8>,
}
impl Lexer {
	fn read_char(&mut self) {
		self.c = self.input.get(self.read_position).copied();
		self.position = self.read_position;
		self.read_position += 1;
	}

	fn new(input: &'static [u8]) -> Self {
		let mut s = Self {
			input,
			position: 0,
			read_position: 0,
			c: None,
		};
		s.read_char(); // To fill up the first char into the `c` field

		s
	}

	fn is_c_letter(&self) -> bool {
		self.c.is_some_and(|c| {
			c.is_ascii_lowercase() || c.is_ascii_uppercase() || c == b'-' || c == b'_'
		})
	}
	fn is_c_digit(&self) -> bool {
		self.c.is_some_and(|c| c.is_ascii_digit())
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

	fn next_token(&mut self) -> TokenData {
		let token = match self.c.unwrap_or(b'\0') {
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
			b' ' => Token::Space,
			b'\t' => Token::Tab,
			b'\r' => Token::CarriageReturn,
			b'\n' => Token::NewLine,
			b'\0' => Token::EndOfFile,
			other => Token::Illegal(other.to_string()),
		};
		self.read_char();

		TokenData::new(token, self.position)
	}
}

impl Iterator for Lexer {
	type Item = TokenData;

	fn next(&mut self) -> Option<Self::Item> {
		let mut token_data = self.next_token();
		while !token_data.is_whitespace() {
			token_data = self.next_token();
		}

		if !token_data.is_eof() {
			Some(token_data)
		} else {
			None
		}
	}
}

#[cfg(test)]
mod test;
