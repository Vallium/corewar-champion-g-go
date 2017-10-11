package main

import (
	"fmt"
	"math/rand"

	popu "github.com/Vallium/corewar-champion-g-go/population"
)

func main() {
	fmt.Println("Corewar champion G")

	rand.Seed(42)
	population := popu.Create(100)

	population.ToFile("./champions-population")
	population.CompileCor()
	population.Evaluate()
}
