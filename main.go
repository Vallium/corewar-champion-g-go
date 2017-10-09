package main

import (
	"fmt"

	popu "github.com/Vallium/corewar-champion-g-go/population"
)

func main() {
	fmt.Println("Corewar champion G")

	population := popu.Create(100)

	population.InjectIndividualsFromFolder("./winners-2014")
}
