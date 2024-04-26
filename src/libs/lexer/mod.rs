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
			b'>' => Token::GreaterThan,
			b'<' => Token::LessThan,
			0 => Token::EndOfFile,
			other => Token::Illegal(other.to_string()),
		};
		self.read_char();

		token
	}
}

#[cfg(test)]
mod test;
