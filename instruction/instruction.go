package instruction

import (
	"strings"
	"bytes"

	parameter "github.com/Vallium/corewar-champion-g-go/parameter"
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

type Register uint8 // 1..=16
type Direct int16
type Indirect int32

type IndReg uint8
type DirReg uint8
type DirInd int32
type DirIndReg int32

type Instruction struct {
	opCode OpCode
	params []*parameter.Parameter
}

func CreateByString(s string) *Instruction {
	var ret Instruction
	ins := strings.Split(s, " ")

	for i, elem := range ins {
		elem = strings.Replace(elem, ",", "", -1)
		if i == 0 {
			ret.setOpCode(elem)
		} else {
			p := parameter.CreateByString(elem)
			ret.params = append(ret.params, p)
		}
	}
	return &ret
}

func (i *Instruction) ToString() string {
	var s string
	buff := bytes.NewBufferString(s)

	switch i.opCode {
	default:
	case Live:
		buff.WriteString("live")
	case Load:
		buff.WriteString("ld")
	case Store:
		buff.WriteString("st")
	case Addition:
		buff.WriteString("add")
	case Substraction:
		buff.WriteString("sub")
	case And:
		buff.WriteString("and")
	case Or:
		buff.WriteString("or")
	case Xor:
		buff.WriteString("xor")
	case ZJump:
		buff.WriteString("zjmp")
	case LoadIndex:
		buff.WriteString("ldi")
	case StoreIndex:
		buff.WriteString("sti")
	case Fork:
		buff.WriteString("fork")
	case LongLoad:
		buff.WriteString("lld")
	case LongLoadIndex:
		buff.WriteString("lldi")
	case LongFork:
		buff.WriteString("lfork")
	case Display:
		buff.WriteString("aff")
	}

	for index, param := range i.params {
		buff.WriteByte(' ')
		buff.WriteString(param.ToString())
		if (index < len(i.params) - 1) {
			buff.WriteByte(',')
		} 
	}
	return buff.String()
}

func (i *Instruction) setOpCode(s string) {
	switch s {
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