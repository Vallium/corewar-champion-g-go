package population

import (
	"fmt"
	"io/ioutil"
	"os"

	champ "github.com/Vallium/corewar-champion-g-go/champion"
)

type Population struct {
	size      int
	champions []*champ.Champion
}

func Create(size int) *Population {
	var ret Population

	ret.size = size
	for i := 0; i < size; i++ {
		ret.champions = append(ret.champions, champ.Random())
	}
	return &ret
}

func (p *Population) InjectIndividualsFromFolder(path string) {
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
