use anyhow::Result;

use super::super::token::Token;

use super::Lexer;

#[test]
fn sanity() -> Result<()> {
	let input = "=;:{}()";
	let mut lexer = Lexer::new(input.as_bytes());

	let expected_results = &[
		Token::Equal,
		Token::Semicolon,
		Token::Colon,
		Token::OpenCurlyBrackets,
		Token::CloseCurlyBrackets,
		Token::OpenParens,
		Token::CloseParens,
	];

	for expected in expected_results {
		assert_eq!(&lexer.next_token(), expected);
	}
	assert_eq!(lexer.next_token(), Token::EndOfFile);

	Ok(())
}
