package token

type Type int

const (
	VALUE     Type = iota
	AT
	EQUALS
	OPERATOR
	SEMICOLON
	EOF
)

type Token struct {
	Type   Type
	Lexeme string
}
