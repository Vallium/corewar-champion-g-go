package champion

import (
	inst "github.com/Vallium/corewar-champion-g-go/instruction"
)

type Champion struct {
	name         string
	comment      string
	instructions []*inst.Instruction
}

func (c *Champion) SetName(name string) {
	c.name = name
}

func (c *Champion) SetComment(comment string) {
	c.comment = comment
}

func (c *Champion) GetName() string {
	return c.name
}

func (c *Champion) GetComment() string {
	return c.comment
}

func (c *Champion) GetInstruction() []*inst.Instruction {
	return c.instructions
}

func (c *Champion) PushInstruction(instruction string) {
	i := inst.Create(instruction)
	c.instructions = append(c.instructions, i)
}

func Create(name string, comment string) *Champion {
	var c Champion

	c.SetName(name)
	c.SetComment(comment)
	return &c
}
