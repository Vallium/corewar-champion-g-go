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
const CrossOverNb int = 4

type Champion struct {
	name         string
	comment      string
	score        float64
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

func (c *Champion) CreateByCopy() *Champion {
	var n Champion

	n.name = c.name
	n.comment = c.comment
	n.score = c.score

	for _, i := range c.instructions {
		n.instructions = append(n.instructions, i.CreateByCopy())
	}
	return &n
}

func Random() *Champion {
	champion := Create(ng.GetRandomName(0)+"_g1", haddock.HaddockSays())
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

func CrossOver(father Champion, mother Champion) (*Champion, *Champion) {
	var crossLen int

	child1 := father.CreateByCopy()
	child2 := mother.CreateByCopy()

	child1Insts := child1.instructions
	child2Insts := child2.instructions
	child1Len := len(child1Insts)
	child2Len := len(child2Insts)

	if child1Len > child2Len {
		crossLen = child2Len
	} else {
		crossLen = child1Len
	}
	if crossLen%CrossOverNb != 0 {
		crossLen -= crossLen % CrossOverNb
	}
	fmt.Println("crossLen: ", child1Len)
	fmt.Println("crossLen: ", child2Len)
	fmt.Println("crossLen: ", crossLen/CrossOverNb)
	for j := 0; j < CrossOverNb; j++ {
		if ((crossLen/CrossOverNb)*j)%2 != 0 {
			for i := (crossLen / CrossOverNb) * j; i < (crossLen/CrossOverNb)*(j+1); i++ {
				child1Insts[i], child2Insts[i] = child2Insts[i], child1Insts[i]
			}
		}
	}

	for child1.GetMemSize() > ChampMaxSize {
		child1.instructions = child1.instructions[:len(child1.instructions)-1]
	}
	for child2.GetMemSize() > ChampMaxSize {
		child2.instructions = child2.instructions[:len(child2.instructions)-1]
	}
	fmt.Println("size c1", child1.GetMemSize())
	fmt.Println("size c2", child2.GetMemSize())
	return child1, child2
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

func (c *Champion) IncScore(score float64) {
	c.score += score
}

func (c *Champion) ResetScore() {
	c.score = 0
}

func (c *Champion) GetScore() float64 {
	return c.score
}

func (c *Champion) GetName() string {
	return c.name
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
