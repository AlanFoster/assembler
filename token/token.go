package token

type Type int

//go:generate stringer -type=Type
const (
	VALUE     Type = iota
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
