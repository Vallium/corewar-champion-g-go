package instruction

import (
	"bytes"
	"math/rand"
	"strings"

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

const OpCodeSize int = 1
const ParamCodeSize int = 1

type Register uint8 // 1..=16
type Indirect int16
type Direct int32

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
			flag := false
			switch ret.opCode {
			case ZJump, LoadIndex, StoreIndex, Fork, LongLoadIndex, LongFork:
				flag = true
			}
			p := parameter.CreateByString(elem, flag)
			ret.params = append(ret.params, p)
		}
	}
	return &ret
}

func CreateRandom() *Instruction {
	var ret Instruction

	switch rand.Intn(17) {
	default:
	case 1:
		ret.opCode = Live
		// live()
		// case 2:
		// 	ret.opCode = Load
		// 	load()
		// case 3:
		// 	ret.opCode = Store
		// 	store()
		// case 4:
		// 	ret.opCode = Addition
		// 	addition()
		// case 5:
		// 	ret.opCode = Substraction
		// 	substraction()
		// case 6:
		// 	ret.opCode = And
		// 	and()
		// case 7:
		// 	ret.opCode = Or
		// 	or()
		// case 8:
		// 	ret.opCode = Xor
		// 	xor()
		// case 9:
		// 	ret.opCode = ZJump
		// 	zJump()
		// case 10:
		// 	ret.opCode = LoadIndex
		// 	loadIndex()
		// case 11:
		// 	ret.opCode = StoreIndex
		// 	storeIndex()
		// case 12:
		// 	ret.opCode = Fork
		// 	fork()
		// case 13:
		// 	ret.opCode = LongLoad
		// 	longLoad()
		// case 14:
		// 	ret.opCode = LongLoadIndex
		// 	longLoadIndex()
		// case 15:
		// 	ret.opCode = LongFork
		// 	longFork()
		// case 16:
		// 	ret.opCode = Display
		// 	display()
	}
	return &ret
}

func (i *Instruction) GetMemSize() int {
	var size int

	switch i.opCode {
	case Load, Store, And, Or, Xor, LoadIndex, StoreIndex, LongLoad, LongLoadIndex:
		size += ParamCodeSize
	}
	for _, p := range i.params {
		size += p.GetMemSize()
	}
	size += OpCodeSize
	return size
}

func (i *Instruction) setOpCode(s string) {
	switch s {
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
		if index < len(i.params)-1 {
			buff.WriteByte(',')
		}
	}
	return buff.String()
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
