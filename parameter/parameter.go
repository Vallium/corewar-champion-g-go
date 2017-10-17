package parameter

import (
	"bytes"
	"math/rand"
	"strconv"
)

type Type uint

const (
	Reg = iota + 1
	Dir
	Ind
)

// TODO: Fix problem with Direct int16 / int32

type Register int // 1..=16
type Direct int32
type Indirect int32

const RegSize int = 1
const IndSize int = 2
const DirSize int = 4

const RegPrefix = "r"
const DirPrefix = "%"
const IndPrefix = ""

type Parameter struct {
	_type      Type
	value      interface{}
	specialDir bool
}

// func Create(_type Type, value interface{}, specialDir bool) *Parameter {
// 	return &Parameter{
// 		_type:      _type,
// 		value:      value,
// 		specialDir: specialDir,
// 	}
// }

func (p *Parameter) CreateByCopy() *Parameter {
	n := *p
	return &n
}

func CreateByString(s string, specialDir bool) *Parameter {
	if s[0] == 'r' {
		v, _ := strconv.ParseInt(s[1:len(s)], 10, 16)
		return &Parameter{
			_type:      Reg,
			value:      int(v),
			specialDir: specialDir,
		}
	} else if s[0] == '%' {
		v, _ := strconv.ParseInt(s[1:len(s)], 10, 32)
		return &Parameter{
			_type:      Dir,
			value:      int32(v),
			specialDir: specialDir,
		}
	}
	v, _ := strconv.ParseInt(s, 10, 32)
	return &Parameter{
		_type:      Ind,
		value:      int(v),
		specialDir: specialDir,
	}
}

func RandRegister() *Parameter {
	var ret Parameter

	ret._type = Reg
	ret.value = rand.Intn(16) + 1
	ret.specialDir = false
	return &ret
}
func RandDirect(specialDir bool) *Parameter {
	var ret Parameter

	ret._type = Dir
	ret.value = rand.Int31()
	ret.specialDir = specialDir
	return &ret
}
func RandIndirect() *Parameter {
	var ret Parameter

	ret._type = Ind
	ret.value = rand.Int()
	ret.specialDir = false
	return &ret
}

func RandIndReg() *Parameter {
	if rand.Intn(2) == 0 {
		return RandIndirect()
	}
	return RandRegister()
}

func RandDirReg(specialDir bool) *Parameter {
	if rand.Intn(2) == 0 {
		return RandDirect(specialDir)
	}
	return RandRegister()
}

func RandDirInd(specialDir bool) *Parameter {
	if rand.Intn(2) == 0 {
		return RandDirect(specialDir)
	}
	return RandIndirect()
}

func RandDirIndReg(specialDir bool) *Parameter {
	r := rand.Intn(3)
	if r == 0 {
		return RandDirect(specialDir)
	} else if r == 1 {
		return RandIndirect()
	}
	return RandRegister()
}

func (p *Parameter) GetMemSize() int {
	var memSize int

	switch p._type {
	case Reg:
		memSize = RegSize
	case Dir:
		if p.specialDir == true {
			memSize = IndSize
		} else {
			memSize = DirSize
		}
	case Ind:
		memSize = IndSize
	}
	return memSize
}

func (p *Parameter) ToString() string {
	var s string
	buff := bytes.NewBufferString(s)

	if p._type == Reg {
		buff.WriteByte('r')
		buff.WriteString(strconv.Itoa(int(p.value.(int))))
	} else if p._type == Dir {
		buff.WriteByte('%')
		buff.WriteString(strconv.Itoa(int(p.value.(int32))))
	} else {
		buff.WriteString(strconv.Itoa(int(p.value.(int))))
	}
	return buff.String()
}
