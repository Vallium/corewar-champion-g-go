package population

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	champ "github.com/Vallium/corewar-champion-g-go/champion"
)

type Population struct {
	size      int
	champions []*champ.Champion
}

func Create(size int) *Population {
	var ret Population

	ret.size = size
	ret.injectIndividualsFromFolder("./winners-2014")
	for i := len(ret.champions); i <= size; i++ {
		ret.champions = append(ret.champions, champ.Random())
	}
	return &ret
}

func (p *Population) injectIndividualsFromFolder(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("ReadDir error: ", err)
		os.Exit(1)
	}

	for _, f := range files {
		c, err := champ.CreateFromFile(string(path + "/" + f.Name()))
		if err != nil {
			fmt.Println("Cahmpion::CreateFromFile error: ", err)
			os.Exit(1)
		}
		p.champions = append(p.champions, c)
	}
}

func (p *Population) ToFile(path string) {
	for _, c := range p.champions {
		c.ToFile(path)
	}
}

func (p *Population) Evaluate() {
	for index, c := range p.champions {
		playMatch(c, p.champions[len(p.champions)-index-1])
	}
}

func (p *Population) CompileCor() {
	for _, c := range p.champions {
		// go func() {
		cmd := exec.Command("./bin/asm", "./champions-population/"+c.GetAssFileName())
		cmd.Run()
		// }()
	}
}

func playMatch(c1 *champ.Champion, c2 *champ.Champion) {
	path := "./champions-population/"

	// go func() {
	cmd := exec.Command("./bin/corewar", path+c1.GetCorFileName(), path+c2.GetCorFileName())
	output, _ := cmd.Output()
	fmt.Println(string(output))
	// cmd.Run()
	// }()
}
