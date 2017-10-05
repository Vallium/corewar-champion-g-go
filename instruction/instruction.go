package instruction

import (
	"strings"
)

type OpCode int

const (
	NoOp = 1 + iota
	Live
	Load
	Store
	Addition
	Substraction
	And
	Or
	Xor
	ZJump
	LoadIndex
	StoreIndex
	Fork
	LongLoad
	LongLoadIndex
	LongFork
	Display
)

type Register	uint8	// 1..=16
type Direct		int16
type Indirect	int32

type IndReg	uint8
type DirReg	uint8
type DirInd	int32
type DirIndReg	int32

type Instruction struct {
	opCode OpCode
	params []interface{}
}

func Create(s string) (Instruction) {
	var ret Instruction 
	ins := strings.Split(s, " ")
	
	for _, elem := range ins {
		elem = strings.Replace(elem, ",", "", -1)
	}
	return ret
}

func (i Instruction) setOpCode(s string) {
	switch (s) {
	default:
		return
	case "live":
		i.opCode = Live
	case "ld":
		i.opCode = Load
	case "st":
		i.opCode = Store
	case "add":
		i.opCode = Addition
	case "sub":
		i.opCode = Substraction
	case "and":
		i.opCode = And
	case "or":
		i.opCode = Or
	case "xor":
		i.opCode = Xor
	case "zjmp":
		i.opCode = ZJump
	case "ldi":
		i.opCode = LoadIndex
	case "sti":
		i.opCode = StoreIndex
	case "fork":
		i.opCode = Fork
	case "lld":
		i.opCode = LongLoad
	case "lldi":
		i.opCode = LongLoadIndex
	case "lfork":
		i.opCode = LongFork
	case "aff":
		i.opCode = Display
	}
}

func live(Direct) {

}

func load(DirInd, Register) {

}

func store(Register, IndReg) {

}

func addition(Register, Register, Register) {

}

func substraction(Register, Register, Register) {

}

func and(DirIndReg, DirIndReg, Register) {

}

func or(DirIndReg, DirIndReg, Register) {

}

func xor(DirIndReg, DirIndReg, Register) {

}

func zJump(Direct) {

}

func loadIndex(DirIndReg, DirReg, Register) {

}

func storeIndex(Register, DirIndReg, DirReg) {

}

func fork(Direct) {

}

func longLoad(DirInd, Register) {

}

func longLoadIndex(DirIndReg, DirReg, Register) {

}

func longFork(Direct) {

}

func display(Register) {

}

func instructionRep(code OpCode) {
	switch (code) {
	default:
		return
	case Live:
		live(0)
	case Load:
		load(0, 0)
	case Store:
		store(0, 0)
	case Addition:
		addition(0, 0, 0)
	case Substraction:
		substraction(0, 0, 0)
	case And:
		and(0, 0, 0)
	case Or:
		or(0, 0, 0)
	case Xor:
		xor(0, 0, 0)
	case ZJump:
		zJump(0)
	case LoadIndex:
		loadIndex(0, 0, 0)
	case StoreIndex:
		storeIndex(0, 0, 0)
	case Fork:
		fork(0)
	case LongLoad:
		longLoad(0, 0)
	case LongLoadIndex:
		longLoadIndex(0, 0, 0)
	case LongFork:
		longFork(0)
	case Display:
		display(0)
	}
}