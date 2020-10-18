package main

import (
	"fmt"
	"github.com/bbdLe/iLua/vm"
	"io/ioutil"

	"github.com/bbdLe/iLua/internal/binchunk"
	"github.com/bbdLe/iLua/internal/log"
	"go.uber.org/zap"
)

func main() {
	log.Logger.Info("lua start")
	data, err := ioutil.ReadFile("hello_world.out")
	if err != nil {
		log.Logger.Fatal("read file fail", zap.Error(err))
	}
	p := binchunk.Undump(data)
	list(p)
}

func list(p *binchunk.Prototype) {
	printHeader(p)
	printCode(p)
	printDetail(p)

	for _, proto := range p.Protos {
		list(proto)
	}
}

func printHeader(p *binchunk.Prototype) {
	funcType := "main"
	if p.LineDefine > 0 {
		funcType = "function"
	}

	varargFlag := ""
	if p.IsVararg > 0 {
		varargFlag = "+"
	}

	fmt.Printf("\n%s <%s:%d,%d> (%d instructions)\n", funcType, p.Source, p.LineDefine, p.LastLineDefine, len(p.Code))
	fmt.Printf("%d%s params, %d slots, %d upvalues, ", p.NumParams, varargFlag, p.MaxStackSize, len(p.Upvalues))
	fmt.Printf("%d locals, %d constants, %d functions\n", len(p.LocVars), len(p.Constants), len(p.Protos))
}

func printCode(p *binchunk.Prototype) {
	for pc, c := range p.Code {
		line := "-"
		if len(p.LineInfo) > 0 {
			line = fmt.Sprintf("%d", p.LineInfo[pc])
		}

		i := vm.Instruction(c)
		fmt.Printf("\t%d\t[%s]\t%s \t", pc+1, line, i.OpName())
		printOperands(i)
		fmt.Printf("\n")
	}
}

func printOperands(i vm.Instruction) {
	switch i.OpMode() {
	case vm.IABC:
		a, b, c := i.ABC()

		fmt.Printf("%d", a)
		if i.BMode() != vm.OpArgN {
			if b > 0xFF {
				fmt.Printf(" %d", -1-b&0xFF)
			} else {
				fmt.Printf(" %d", b)
			}
		}
		if i.CMode() != vm.OpArgN {
			if c > 0xFF {
				fmt.Printf(" %d", -1-c&0xFF)
			} else {
				fmt.Printf(" %d", c)
			}
		}
	case vm.IABx:
		a, bx := i.ABx()

		fmt.Printf("%d", a)
		if i.BMode() == vm.OpArgK {
			fmt.Printf(" %d", -1-bx)
		} else if i.BMode() == vm.OpArgU {
			fmt.Printf(" %d", bx)
		}
	case vm.IAsBx:
		a, sBx := i.AsBx()
		fmt.Printf("%d %d", a, sBx)
	case vm.IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}
}

func printDetail(p *binchunk.Prototype) {
	fmt.Printf("constants (%d):\n", len(p.Constants))
	for i, k := range p.Constants {
		fmt.Printf("\t%d\t%s\n", i+1, constantToString(k))
	}

	fmt.Printf("locals (%d):\n", len(p.LocVars))
	for i, localVar := range p.LocVars {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i+1, localVar.VarName, localVar.StartPC, localVar.EndPC)
	}

	fmt.Printf("upvalues (%d):\n", len(p.Upvalues))
	for i, upvalue := range p.Upvalues {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i+1, upvalName(p, i), upvalue.Idx, upvalue.Instack)
	}
}

func constantToString(k interface{}) string {
	switch k.(type) {
	case nil:
		return "nil"
	case bool:
		return fmt.Sprintf("%t", k)
	case float64:
		return fmt.Sprintf("%g", k)
	case int64:
		return fmt.Sprintf("%d", k)
	case string:
		return fmt.Sprintf("%q", k)
	default:
		return "?"
	}
}

func upvalName(p *binchunk.Prototype, idx int) string {
	if len(p.UpvalueNames) > 0 {
		return p.UpvalueNames[idx]
	} else {
		return "-"
	}
}
