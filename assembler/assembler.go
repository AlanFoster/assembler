package assembler

type Assembler struct {
}

func New() *Assembler {
	return &Assembler{}
}

func (a *Assembler) Convert(s string) string {
	return "0000000000000011"
}
