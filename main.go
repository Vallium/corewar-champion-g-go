package main

import (
	"fmt"

	parser "github.com/Vallium/corewar-champion-g-go/parser"
)

func main() {
	fmt.Println("Corewar champion G");
	parser.Parse("./winners-2014/_-clear.s")
}