package vm

const MAXARG_Bx = 1<<18 - 1
const MAXARG_sBx = MAXARG_Bx >> 1

type Instruction uint32

func (self Instruction) Opcode() int {
	return int(self & 0x3F)
}

func (self Instruction) ABC() (int, int, int) {
	a := int(self >> 6 & 0xFF)
	c := int(self >> 14 & 0x1FF)
	b := int(self >> 23 & 0x1FF)
	return a, b, c
}

func (self Instruction) ABx() (int, int) {
	a := int(self >> 6 & 0xFF)
	bx := int(self >> 14)
	return a, bx
}

func (self Instruction) AsBx() (int, int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx
}

func (self Instruction) Ax() int {
	return int(self >> 6)
}

func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}
