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
const Smallest int = 3

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
			specialDir := false
			switch ret.opCode {
			case ZJump, LoadIndex, StoreIndex, Fork, LongLoadIndex, LongFork:
				specialDir = true
			}
			p := parameter.CreateByString(elem, specialDir)
			ret.params = append(ret.params, p)
		}
	}
	return &ret
}

func CreateRandom() *Instruction {
	var ret Instruction

	switch rand.Intn(16) + 1 {
	case 1:
		ret.opCode = Live
		ret.params = append(ret.params, parameter.RandDirect(false))
	case 2:
		ret.opCode = Load
		ret.params = append(ret.params, parameter.RandDirInd(false))
		ret.params = append(ret.params, parameter.RandRegister())
	case 3:
		ret.opCode = Store
		ret.params = append(ret.params, parameter.RandRegister())
		ret.params = append(ret.params, parameter.RandIndReg())
	case 4:
		ret.opCode = Addition
		ret.params = append(ret.params, parameter.RandRegister())
		ret.params = append(ret.params, parameter.RandRegister())
		ret.params = append(ret.params, parameter.RandRegister())
	case 5:
		ret.opCode = Substraction
		ret.params = append(ret.params, parameter.RandRegister())
		ret.params = append(ret.params, parameter.RandRegister())
		ret.params = append(ret.params, parameter.RandRegister())
	case 6:
		ret.opCode = And
		ret.params = append(ret.params, parameter.RandDirIndReg(false))
		ret.params = append(ret.params, parameter.RandDirIndReg(false))
		ret.params = append(ret.params, parameter.RandRegister())
	case 7:
		ret.opCode = Or
		ret.params = append(ret.params, parameter.RandDirIndReg(false))
		ret.params = append(ret.params, parameter.RandDirIndReg(false))
		ret.params = append(ret.params, parameter.RandRegister())
	case 8:
		ret.opCode = Xor
		ret.params = append(ret.params, parameter.RandDirIndReg(false))
		ret.params = append(ret.params, parameter.RandDirIndReg(false))
		ret.params = append(ret.params, parameter.RandRegister())
	case 9:
		ret.opCode = ZJump
		ret.params = append(ret.params, parameter.RandDirect(true))
	case 10:
		ret.opCode = LoadIndex
		ret.params = append(ret.params, parameter.RandDirIndReg(true))
		ret.params = append(ret.params, parameter.RandDirReg(true))
		ret.params = append(ret.params, parameter.RandRegister())
	case 11:
		ret.opCode = StoreIndex
		ret.params = append(ret.params, parameter.RandRegister())
		ret.params = append(ret.params, parameter.RandDirIndReg(true))
		ret.params = append(ret.params, parameter.RandDirReg(true))
	case 12:
		ret.opCode = Fork
		ret.params = append(ret.params, parameter.RandDirect(true))
	case 13:
		ret.opCode = LongLoad
		ret.params = append(ret.params, parameter.RandDirInd(false))
		ret.params = append(ret.params, parameter.RandRegister())
	case 14:
		ret.opCode = LongLoadIndex
		ret.params = append(ret.params, parameter.RandDirIndReg(true))
		ret.params = append(ret.params, parameter.RandDirReg(true))
		ret.params = append(ret.params, parameter.RandRegister())
	case 15:
		ret.opCode = LongFork
		ret.params = append(ret.params, parameter.RandDirect(true))
	case 16:
		ret.opCode = Display
		ret.params = append(ret.params, parameter.RandRegister())
	}
	return &ret
}

func (i *Instruction) GetMemSize() int {
	var size int

	switch i.opCode {
	case Load, Store, Addition, Substraction, And, Or, Xor, LoadIndex, StoreIndex, LongLoad, LongLoadIndex, Display:
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
