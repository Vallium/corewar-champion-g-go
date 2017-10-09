package main

import (
	"fmt"
	"os"

	popu "github.com/Vallium/corewar-champion-g-go/population"
	champ "github.com/Vallium/corewar-champion-g-go/champion"
)

func main() {
	fmt.Println("Corewar champion G")

	population := popu.Create(100)

	population.InjectPersonsFromFolder("./winners-2014")
	champion, err := champ.CreateFromFile("./winners-2014/_-clear.s")
	if err != nil {
		fmt.Println("parser error:", err)
		os.Exit(1)
	}
	champion.ToFile()
}
