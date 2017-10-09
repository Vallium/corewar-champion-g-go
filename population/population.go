package population

import (
	"os"
	"fmt"
	"io/ioutil"

	champ "github.com/Vallium/corewar-champion-g-go/champion"
	// parser "github.com/Vallium/corewar-champion-g-go/parser"
)

type Population struct {
	size int
	champions *[]champ.Champion
}

func Create(size int) *Population {
	return &Population{
		size: size,
	}
}

func (*Population) InjectPersonsFromFolder(path string) {
	files, err := ioutil.ReadDir(path)
	
	if err != nil {
		fmt.Println("readDir error: ", err)
		os.Exit(1);
    }

    for _, f := range files {
    	fmt.Println(f.Name())
    }
	// champ.Champion, err := 

}