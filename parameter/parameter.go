package parameter

import (
	"strconv"
	"bytes"
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

const RegPrefix = "r"
const DirPrefix = "%"
const IndPrefix = ""

type Parameter struct {
	_type Type
	value interface{}
}

func Create(_type Type, value interface{}) *Parameter {
	return &Parameter{
		_type: _type,
		value: value,
	}
}

func CreateByString(s string) *Parameter {
	if s[0] == 'r' {
		v, _ := strconv.ParseUint(s[1:len(s)], 10, 8)
		return &Parameter{
			_type: Reg,
			value: uint8(v),
		}
	} else if s[0] == '%' {
		v, _ := strconv.ParseInt(s[1:len(s)], 10, 32)
		return &Parameter{
			_type: Dir,
			value: int32(v),
		}
	}
	v, _ := strconv.ParseInt(s, 10, 32)
	return &Parameter{
		_type: Ind,
		value: int32(v),
	}
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