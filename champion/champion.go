package champion

import (
	"os"
	"fmt"

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

func (c *Champion) PushInstruction(instruction string) {
	i := inst.CreateByString(instruction)
	c.instructions = append(c.instructions, i)
}

func (c *Champion) ToFile() {
	f, err := os.Create("./" + c.name + ".s")
    if err != nil {
		fmt.Println("os.Create error: ", err)
		os.Exit(1)
	}
	defer f.Close()
	
	f.WriteString(".name \"" + c.name + "\"\n")
	f.WriteString(".comment \"" + c.comment + "\"\n\n")
	for _, ins := range c.instructions {
		f.WriteString(ins.ToString() + "\n")
	}
}

func Create(name string, comment string) *Champion {
	return &Champion {
		name: name,
		comment: comment,
	}
}
