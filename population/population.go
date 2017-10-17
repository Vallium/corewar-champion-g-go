package population

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	champ "github.com/Vallium/corewar-champion-g-go/champion"
	"github.com/bradfitz/slice"
)

type Population struct {
	size       int
	generation int
	champions  []*champ.Champion
}

type MTCommand struct {
	cmd *exec.Cmd
	c1  *champ.Champion
	c2  *champ.Champion
}

func Create(size int) *Population {
	var ret Population

	ret.size = size
	ret.generation = 1
	ret.injectIndividualsFromFolder("./winners-2014")
	for i := len(ret.champions); i <= size; i++ {
		ret.champions = append(ret.champions, champ.Random())
	}
	return &ret
}

func (p *Population) incGeneration() {
	p.generation++
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

func (p *Population) toFile(path string) {
	for _, c := range p.champions {
		c.ToFile(path)
	}
}

func (p *Population) evaluate() {
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
}

func (p *Population) compileCor() {
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

func (p *Population) GeneticLoopStart() {
	// for true {
	p.toFile("./champions-population")
	p.compileCor()
	p.evaluate()
	// p.newGeneration()
	p.tournamentSelection()
	// }
}

func (p *Population) tournamentSelection() {
	var chosens []champ.Champion

	for j := 0; j < 2; j++ {
		var pool []champ.Champion

		for i := 0.0; i < 0.10*float64(p.size); i += 1.0 {
			c := p.champions[rand.Intn(p.size)]
			// TODO: Check if duplicate random chosen
			pool = append(pool, *c)
		}
		slice.Sort(pool[:], func(i, j int) bool {
			return pool[i].GetScore() > pool[j].GetScore()
		})
		chosens = append(chosens, *pool[0].CreateByCopy())
	}

	child1, child2 := champ.CrossOver(chosens[0], chosens[1])
	fmt.Println(child1.GetScore())
	fmt.Println(child2.GetScore())
}

// func (p *Population) rouletteWheel(champ.Champion, champ.Champion) {
// 	var sum float64
// 	var probability float64

// 	for _, c := range p.champions {
// 		sum += c.GetScore()
// 	}

// 	var
// 	for all members of population
// 		probability = sum of probabilities + (fitness / sum)
// 		sum of probabilities += probability
// 	end for

// // loop until new population is full
// // do this twice
// // 	number = Random between 0 and 1
// //   for all members of population
// // 	  if number > probability but less than next probability
// // 		   then you have been selected
// //   end for
// // end
// // create offspring
// // end loop
// }

func (p *Population) newGeneration() {
	// var p1 Population

	slice.Sort(p.champions[:], func(i, j int) bool {
		return p.champions[i].GetScore() > p.champions[j].GetScore()
	})

	var newGen Population

	for _, c := range p.champions {
		cc := c.CreateByCopy()
		newGen.champions = append(newGen.champions, cc)
		cc.ResetScore()
	}

	for _, c := range p.champions {
		fmt.Println(c.GetName(), ": ", c.GetScore())
	}
	fmt.Println("\n")

	for _, c := range newGen.champions {
		fmt.Println(c.GetName(), ": ", c.GetScore())
	}
}
