// The lexer(tokenizer) gets the users input and transforms it into a token that will then be used in the parser

// @TODO attach filenames and line numbers to the tokens to track down lexing and parsing errors (initialize lexer with io.Reader and the filename)

package token

type TokenType string // Defining this as a string allows us to use many different values as TokenTypes @TODO Look into making this an int or a byte for increased performance

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
