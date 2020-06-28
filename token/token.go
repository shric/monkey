package token

type Type int

const (
	ILLEGAL = iota
	EOF

	// Identifiers + literals
	IDENT  // add, foobar, x, y, ...
	NUMBER // 1343456
	STRING // "foobar"

	// Operators
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /

	LT // <
	GT // >

	REGEX  // ~
	EQ     // ==
	NOT_EQ // !=

	AND // &&
	OR  // ||

	// Delimiters
	COMMA     // ,
	SEMICOLON // ;
	COLON     // :

	LPAREN   // (
	RPAREN   // )
	LBRACE   // {
	RBRACE   // }
	LBRACKET // [
	RBRACKET // ]

	// Keywords
	FUNCTION // FUNCTION
	LET      // LET
	TRUE     // TRUE
	FALSE    // FALSE
	IF       // IF
	ELSE     // ELSE
	RETURN   // RETURN

	NUMERIC_SUFFIX // NUMERIC_SUFFIX
)

var tokStrings = []string{
	"ILLEGAL",
	"EOF",

	// Identifiers + literals
	"IDENT",
	"NUMBER",
	"STRING",

	// Operators
	"=",
	"+",
	"-",
	"!",
	"*",
	"/",

	"<",
	">",

	"~",
	"==",
	"!=",

	"&&",
	"||",

	// Delimiters
	",",
	";",
	":",
	".",

	"(",
	")",
	"{",
	"}",
	"[",
	"]",

	// Keywords
	"FUNCTION",
	"LET",
	"TRUE",
	"FALSE",
	"IF",
	"ELSE",
	"RETURN",

	"NUMERIC SUFFIX",
}

func (t Type) String() string {
	return tokStrings[t]
}
func (t Token) String() string {
	return t.Type.String()
}

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

var numericSuffixes = map[string]Type{
	"KB":  NUMERIC_SUFFIX,
	"KiB": NUMERIC_SUFFIX,
	"MB":  NUMERIC_SUFFIX,
	"MiB": NUMERIC_SUFFIX,
	"GB":  NUMERIC_SUFFIX,
	"GiB": NUMERIC_SUFFIX,
	"TB":  NUMERIC_SUFFIX,
	"TiB": NUMERIC_SUFFIX,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	if tok, ok := numericSuffixes[ident]; ok {
		return tok
	}
	return IDENT
}
