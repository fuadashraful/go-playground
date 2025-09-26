package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Person struct {
	name      string
	rightFork int
	leftFork  int
}

var persons = []Person{
	{"Plato", 0, 1},
	{"Socrates", 1, 2},
	{"Aristotle", 2, 3},
	{"Descartes", 3, 4},
	{"Kant", 4, 0},
}

var hunger = 3
var eatTimes = 1 * time.Second
var thinkTimes = 3 * time.Second
var sleepTimes = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func diningPerson(wg, seatedWg *sync.WaitGroup, forks *map[int]*sync.Mutex, p *Person) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", p.name)
	seatedWg.Done()
	seatedWg.Wait()

	for i := 0; i < hunger; i++ {
		fmt.Printf("%s is hungry.\n", p.name)

		if p.leftFork < p.rightFork {
			(*forks)[p.leftFork].Lock()
			fmt.Printf("%s picked up left fork %d.\n", p.name, p.leftFork)
			(*forks)[p.rightFork].Lock()
			fmt.Printf("%s picked up right fork %d.\n", p.name, p.rightFork)
		} else {
			(*forks)[p.rightFork].Lock()
			fmt.Printf("%s picked up right fork %d.\n", p.name, p.rightFork)
			(*forks)[p.leftFork].Lock()
			fmt.Printf("%s picked up left fork %d.\n", p.name, p.leftFork)
		}
		fmt.Printf("%s has both forks and is eating.\n", p.name)
		time.Sleep(eatTimes)

		fmt.Printf("%s is thinking.\n", p.name)
		time.Sleep(thinkTimes)

		(*forks)[p.leftFork].Unlock()
		(*forks)[p.rightFork].Unlock()
		fmt.Printf("%s has finished eating and put down both forks.\n", p.name)
	}

	fmt.Printf("%s leaves the table.\n", p.name)

	orderMutex.Lock()
	orderFinished = append(orderFinished, p.name)
	orderMutex.Unlock()
}
func dine() {
	wg := sync.WaitGroup{}
	wg.Add(len(persons))

	seatedWg := sync.WaitGroup{}
	seatedWg.Add(len(persons))

	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(persons); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(persons); i++ {
		go diningPerson(&wg, &seatedWg, &forks, &persons[i])
	}

	wg.Wait()
}

func main() {
	fmt.Println("The table is empty !!")

	time.Sleep(sleepTimes)

	dine()

	time.Sleep(sleepTimes)

	fmt.Println("The table is empty again !!")
	fmt.Printf("Order of finishing: %s\n", strings.Join(orderFinished, ", "))

}
