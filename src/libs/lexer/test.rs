use anyhow::Result;

use super::super::token::{Token, TokenData};

use super::Lexer;

#[test]
fn sanity() -> Result<()> {
	let input = "=;:{} ()\t\r\n";
	let mut lexer = Lexer::new(input.as_bytes());

	let expected_results = &[
		TokenData::new(Token::Equal, 0),
		TokenData::new(Token::Semicolon, 1),
		TokenData::new(Token::Colon, 2),
		TokenData::new(Token::OpenCurlyBrackets, 3),
		TokenData::new(Token::CloseCurlyBrackets, 4),
		TokenData::new(Token::Space, 5),
		TokenData::new(Token::OpenParens, 6),
		TokenData::new(Token::CloseParens, 7),
		TokenData::new(Token::Tab, 8),
		TokenData::new(Token::CarriageReturn, 9),
		TokenData::new(Token::NewLine, 10),
	];

	for expected in expected_results {
		assert_eq!(&lexer.next_token(), expected);
	}
	assert_eq!(lexer.next_token(), TokenData::new(Token::EndOfFile, 11));

	Ok(())
}

#[test]
fn sanity_iter() -> Result<()> {
	let input = "=;:{} ()\t\r\n";
	let lexer = Lexer::new(input.as_bytes());

	let results = lexer.collect::<Vec<TokenData>>();
	let expected_results = &[
		TokenData::new(Token::Equal, 0),
		TokenData::new(Token::Semicolon, 1),
		TokenData::new(Token::Colon, 2),
		TokenData::new(Token::OpenCurlyBrackets, 3),
		TokenData::new(Token::CloseCurlyBrackets, 4),
		TokenData::new(Token::OpenParens, 6),
		TokenData::new(Token::CloseParens, 7),
	];

	assert_eq!(results, expected_results);

	Ok(())
}
