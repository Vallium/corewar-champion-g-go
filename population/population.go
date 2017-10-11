package population

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	champ "github.com/Vallium/corewar-champion-g-go/champion"
)

type Population struct {
	size      int
	champions []*champ.Champion
}

type MTCommand struct {
	cmd *exec.Cmd
	c1  *champ.Champion
	c2  *champ.Champion
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
	path := "./champions-population/"
	tasks := make(chan *MTCommand, 256)

	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			for task := range tasks {
				output, _ := task.cmd.Output()
				tab := strings.Split(string(output), "\n")
				var cycles float64
				var winner *champ.Champion

				for i, s := range tab {
					if i > 3 && len(s) != 0 {
						if s[0] == 'C' && s[1] == 'o' {
							// fmt.Println(strings.Split(s, " ")[1], " WON")
							w, _ := strconv.ParseInt(strings.Split(s, " ")[1][:1], 10, 32)
							if w == 1 {
								winner = task.c1
							} else if w == 2 {
								winner = task.c2
							}
						} else if s[0] == 'P' {
							// fmt.Println(strings.Split(s, " ")[1], " says alive")
						} else if s[0] == 'I' {
							// fmt.Println(strings.Split(s, " ")[4], " cycles")
							cycles, _ = strconv.ParseFloat(strings.Split(s, " ")[4], 64)
						}
					}
				}

				fitness := 1.0 + (1.0 / cycles)
				winner.IncScore(fitness)
				// fmt.Println(winner.GetScore())
			}
			wg.Done()
		}()
	}

	for _, c1 := range p.champions {
		for _, c2 := range p.champions {
			if c1 == c2 {
				continue
			}
			tasks <- &MTCommand{
				cmd: exec.Command("./bin/corewar", "-v", "2", path+c1.GetCorFileName(), path+c2.GetCorFileName()),
				c1:  c1,
				c2:  c2,
			}
		}
	}

	close(tasks)
	wg.Wait()

	for _, c := range p.champions {
		fmt.Println(c.GetName(), ": ", c.GetScore())
	}
}

func (p *Population) CompileCor() {
	tasks := make(chan *exec.Cmd, 64)

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for cmd := range tasks {
				cmd.Run()
			}
			wg.Done()
		}()
	}

	for _, c := range p.champions {
		tasks <- exec.Command("./bin/asm", "./champions-population/"+c.GetAssFileName())
	}
	close(tasks)
	wg.Wait()
}
