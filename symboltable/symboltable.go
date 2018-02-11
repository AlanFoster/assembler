package symboltable

type SymbolTable struct {
	values map[string]int
}

func New() *SymbolTable {
	return &SymbolTable{
		// Populate with pre-fined symbols
		values: map[string]int{
			"SP":   0,
			"LCL":  1,
			"ARG":  2,
			"THIS": 3,
			"THAT": 4,

			// 'Registers'
			"R0":  0,
			"R1":  1,
			"R2":  2,
			"R3":  3,
			"R4":  4,
			"R5":  5,
			"R6":  6,
			"R7":  7,
			"R8":  8,
			"R9":  9,
			"R10": 10,
			"R11": 11,
			"R12": 12,
			"R13": 13,
			"R14": 14,
			"R15": 15,

			// Screen and keyboard, for Direct Memory Access
			"SCREEN": 0x4000,
			"KBD":    0x6000,
		},
	}
}

func (st *SymbolTable) Add(entry string, address int) {
	st.values[entry] = address
}

func (st *SymbolTable) Get(entry string) int {
	return st.values[entry]
}

func (st *SymbolTable) Contains(entry string) bool {
	_, ok := st.values[entry]
	return ok
}
