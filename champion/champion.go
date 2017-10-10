package champion

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	inst "github.com/Vallium/corewar-champion-g-go/instruction"
	haddock "github.com/Vallium/corewar-champion-g-go/pkg/captainhaddock"
	ng "github.com/Vallium/corewar-champion-g-go/pkg/namegenerator"
)

const MemSize int = 4 * 1024
const ChampMaxSize int = MemSize / 6
const MinIns int = 30
const MaxIns int = ChampMaxSize / inst.Smallest

type Champion struct {
	name         string
	comment      string
	score        float32
	instructions []*inst.Instruction
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
			champion.pushInstruction(scanner.Text())
		}
		index++
	}
	return champion, scanner.Err()
}

func Random() *Champion {
	champion := Create(ng.GetRandomName(0), haddock.HaddockSays())
	// nbInstucions := rand.Intn(MaxIns-MinIns+1) + MinIns
	var total_size int

	for true {
		i := inst.CreateRandom()
		total_size += i.GetMemSize()
		if total_size > ChampMaxSize {
			break
		}
		champion.instructions = append(champion.instructions, i)
	}
	return champion
}

func (c *Champion) pushInstruction(instruction string) {
	i := inst.CreateByString(instruction)
	c.instructions = append(c.instructions, i)
}

func (c *Champion) SetName(name string) {
	c.name = name
}

func (c *Champion) SetComment(comment string) {
	c.comment = comment
}

func (c *Champion) GetMemSize() int {
	var memSize int

	for _, i := range c.instructions {
		memSize += i.GetMemSize()
	}
	return memSize
}

func (c *Champion) GetAssFileName() string {
	return c.name + ".s"
}

func (c *Champion) GetCorFileName() string {
	return c.name + ".cor"
}

func (c *Champion) ToFile(path string) {
	f, err := os.Create(path + "/" + c.name + ".s")
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
