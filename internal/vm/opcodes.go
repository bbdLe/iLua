package vm

/* OpMode */
const (
	IABC  = iota //	B:9	C:9	A:8	OP:6
	IABx         //	Bx:18	A:8	OP:6
	IAsBx        //	sBx:18	A:8	OP:6
	IAx          //	Ax:26		OP:6
)

/* OpArgMask */
const (
	OpArgN = iota // argument is not used
	OpArgU        // argument is used
	OpArgR        // argument is a register or a jump offset
	OpArgK        // argument is a constant or register/constant
)

/* OpCode */
const (
	OP_MOVE = iota
	OP_LOADK
	OP_LOADKX
	OP_LOADBOOL
	OP_LOADNIL
	OP_GETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETUPVAL
	OP_SETTABLE
	OP_NEWTABLE
	OP_SELF
	OP_ADD
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TAILCALL
	OP_RETURN
	OP_FORLOOP
	OP_FORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAAGR
)

type opcode struct {
	testFlag byte
	setAFlag byte
	argBMode byte
	argCMode byte
	opMode   byte
	name     string
}

var opcodes = []opcode{
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IABC, name: "Move	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgN, opMode: IABx, name: "LOADK	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgN, argCMode: OpArgN, opMode: IABx, name: "LOADKX	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgU, opMode: IABC, name: "LOADBOOL"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgN, opMode: IABC, name: "LOADNIL	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgN, opMode: IABC, name: "GETUPVAL"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgK, opMode: IABC, name: "GETTABUP"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgK, opMode: IABC, name: "GETTABLE"},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "SETTABUP"},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgU, argCMode: OpArgN, opMode: IABC, name: "SETUPVAL"},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "SETTABLE"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgU, opMode: IABC, name: "NEWTABLE"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgK, opMode: IABC, name: "SELF	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "ADD		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "SUB		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "MUL		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "MOD		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "POW		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "DIV		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "IDVI	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "BAND	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "BOR		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "BXOR	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "SHL		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "SHR		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IABC, name: "UNM		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IABC, name: "BNOT	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IABC, name: "NOT		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IABC, name: "LEN		"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgR, opMode: IABC, name: "CONCAT	"},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgR, argCMode: OpArgN, opMode: IAsBx, name: "JMP	"},
	{testFlag: 1, setAFlag: 0, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "EQ		"},
	{testFlag: 1, setAFlag: 0, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "LT		"},
	{testFlag: 1, setAFlag: 0, argBMode: OpArgK, argCMode: OpArgK, opMode: IABC, name: "LE		"},
	{testFlag: 1, setAFlag: 0, argBMode: OpArgN, argCMode: OpArgU, opMode: IABC, name: "TEST	"},
	{testFlag: 1, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgU, opMode: IABC, name: "TESTSET	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgU, opMode: IABC, name: "CALL	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgU, opMode: IABC, name: "TAILCALL"},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgU, argCMode: OpArgN, opMode: IABC, name: "RETURN	"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IAsBx, name: "FORLOOP"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IAsBx, name: "FORPREP"},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgN, argCMode: OpArgU, opMode: IABC, name: "TFORCALL"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgR, argCMode: OpArgN, opMode: IAsBx, name: "TFORLOOP"},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgU, argCMode: OpArgU, opMode: IABC, name: "SETLIST"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgN, opMode: IABx, name: "CLOSURE"},
	{testFlag: 0, setAFlag: 1, argBMode: OpArgU, argCMode: OpArgN, opMode: IABC, name: "VARARG "},
	{testFlag: 0, setAFlag: 0, argBMode: OpArgU, argCMode: OpArgU, opMode: IAx, name: "EXTRAAGR"},
}
