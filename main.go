package main

import (
	"fmt"
	"os"

	parser "github.com/Vallium/corewar-champion-g-go/parser"
)

func main() {
	fmt.Println("Corewar champion G")
	_, err := parser.Parse("./winners-2014/_-clear.s")
	if err != nil {
		fmt.Println("parser error:", err)
		os.Exit(1)
	}
	// fmt.Printf("%v\n", champion.GetInstruction())
}
