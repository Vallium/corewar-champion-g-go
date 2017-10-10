package parameter

import (
	"bytes"
	"strconv"
)

type Type uint

const (
	Reg = iota + 1
	Dir
	Ind
)

// TODO: Fix problem with Direct int16 / int32

type Register uint8 // 1..=16
type Direct int32
type Indirect int32

const RegSize int = 1
const IndSize int = 2
const DirSize int = 4

const RegPrefix = "r"
const DirPrefix = "%"
const IndPrefix = ""

type Parameter struct {
	_type Type
	value interface{}
	flag  bool
}

func Create(_type Type, value interface{}, flag bool) *Parameter {
	return &Parameter{
		_type: _type,
		value: value,
		flag:  flag,
	}
}

func CreateByString(s string, flag bool) *Parameter {
	if s[0] == 'r' {
		v, _ := strconv.ParseUint(s[1:len(s)], 10, 8)
		return &Parameter{
			_type: Reg,
			value: uint8(v),
			flag:  flag,
		}
	} else if s[0] == '%' {
		v, _ := strconv.ParseInt(s[1:len(s)], 10, 32)
		return &Parameter{
			_type: Dir,
			value: int32(v),
			flag:  flag,
		}
	}
	v, _ := strconv.ParseInt(s, 10, 32)
	return &Parameter{
		_type: Ind,
		value: int32(v),
		flag:  flag,
	}
}

func (p *Parameter) GetMemSize() int {
	var memSize int

	switch p._type {
	case Reg:
		memSize = RegSize
	case Dir:
		if p.flag == true {
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
		buff.WriteString(strconv.Itoa(int(p.value.(uint8))))
	} else if p._type == Dir {
		buff.WriteByte('%')
		buff.WriteString(strconv.Itoa(int(p.value.(int32))))
	} else {
		buff.WriteString(strconv.Itoa(int(p.value.(int32))))
	}
	return buff.String()
}
