package champion

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	inst "github.com/Vallium/corewar-champion-g-go/instruction"
)

const MemSize = 4 * 1024
const ChampMaxSize = MemSize / 6

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
	return &Champion{
		name:    name,
		comment: comment,
	}
}

func CreateFromFile(path string) (*Champion, error) {
	var champion *Champion
	file, err := os.Open(path)

	champion = Create("", "")

	if err != nil {
		return champion, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index := 0
	for scanner.Scan() {
		if index == 0 {
			champion.SetName(strings.Split(scanner.Text(), "\"")[1])
		} else if index == 1 {
			champion.SetComment(strings.Split(scanner.Text(), "\"")[1])
		} else if index > 2 {
			champion.PushInstruction(scanner.Text())
		}
		index++
	}
	return champion, scanner.Err()
}
