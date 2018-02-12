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

	JUMP

	EOF
)

type Token struct {
	Type   Type
	Lexeme string
}

// The set of possible JUMP lexemes. Stores as a map for simple lookups.
var jumpValues = map[string]bool {
	"JGT": true,
	"JEQ": true,
	"JGE": true,
	"JLT": true,
	"JNE": true,
	"JLE": true,
	"JMP": true,
}

func MapValue(value string) Type {
	if _, ok := jumpValues[value]; ok {
		return JUMP
	}

	return VALUE
}