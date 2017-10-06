package parameter

import (
	"strconv"
)

type Type uint

const (
	Reg = iota + 1
	Dir
	Ind
)

type Register uint8 // 1..=16
type Direct int16
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
		v, _ := strconv.ParseInt(s[1:len(s)], 10, 16)
		return &Parameter{
			_type: Dir,
			value: int16(v),
		}
	}
	v, _ := strconv.ParseInt(s, 10, 32)
	return &Parameter{
		_type: Dir,
		value: int32(v),
	}
}
