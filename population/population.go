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
	return &Population{
		size: size,
	}
}

func (p *Population) InjectPersonsFromFolder(path string) {
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
