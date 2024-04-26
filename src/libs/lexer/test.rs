use anyhow::Result;

use super::super::token::Token;

use super::Lexer;

#[test]
fn sanity() -> Result<()> {
	let input = "=;:{} ()\t\r\n";
	let mut lexer = Lexer::new(input.as_bytes());

	let expected_results = &[
		Token::Equal,
		Token::Semicolon,
		Token::Colon,
		Token::OpenCurlyBrackets,
		Token::CloseCurlyBrackets,
		Token::Space,
		Token::OpenParens,
		Token::CloseParens,
		Token::Tab,
		Token::CarriageReturn,
		Token::NewLine,
	];

	for expected in expected_results {
		assert_eq!(&lexer.next_token(), expected);
	}
	assert_eq!(lexer.next_token(), Token::EndOfFile);

	Ok(())
}

#[test]
fn sanity_iter() -> Result<()> {
	let input = "=;:{} ()\t\r\n";
	let lexer = Lexer::new(input.as_bytes());

	let results = lexer.collect::<Vec<Token>>();
	let expected_results = &[
		Token::Equal,
		Token::Semicolon,
		Token::Colon,
		Token::OpenCurlyBrackets,
		Token::CloseCurlyBrackets,
		Token::OpenParens,
		Token::CloseParens,
	];

	assert_eq!(results, expected_results);

	Ok(())
}
