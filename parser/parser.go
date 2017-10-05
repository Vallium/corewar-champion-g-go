package parser

import (
	"fmt"
	"strings"
	"bufio"
	"os"

	champ "github.com/Vallium/corewar-champion-g-go/champion"
)

func Parse(path string) (*champ.Champion, error) {
	var champion *champ.Champion
	file, err := os.Open(path)

	champion = champ.Create("", "")

	if err != nil {
		return champion, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	index := 0
	for scanner.Scan() {
		if (index == 0) {
			champion.SetName(strings.Split(scanner.Text(), "\"")[1])
		} else if (index == 1) {
			champion.SetComment(strings.Split(scanner.Text(), "\"")[1])
		} else {
			
		}
		lines = append(lines, scanner.Text())
		index++
	}
	fmt.Println(champion.GetName())
	fmt.Println(champion.GetComment())
	return champion, scanner.Err()
}