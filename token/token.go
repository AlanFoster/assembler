package token

type Type int

//go:generate stringer -type=Type
const (
	VALUE     Type = iota
	LEFT_BRACKET
	RIGHT_BRACKET
	AT
	EQUALS
	OPERATOR
	SEMICOLON
	INVALID
	EOF
)

type Token struct {
	Type   Type
	Lexeme string
}
