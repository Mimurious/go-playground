package main

import (
	"fmt"
	"sync"
)

type chopStick struct {
	sync.Mutex
}

type Philosopher struct {
	Name                string
	RightChop, LeftChop *chopStick
}

func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	for {
		p.LeftChop.Lock()
		p.RightChop.Lock()

		fmt.Println(p.Name, "is Eating...")

		p.LeftChop.Unlock()
		p.RightChop.Unlock()
	}
}

func main() {
	Count := 5

	Chopsticks := make([]*chopStick, Count)
	for i := 0; i < Count; i++ {
		Chopsticks[i] = new(chopStick)
	}

	Philosophers := make([]*Philosopher, Count)
	for i := 0; i < Count; i++ {
		// Dijkstra's Solution
		left := i
		right := (i + 1) % Count

		if left > right {
			left, right = right, left
		}
		var name = fmt.Sprintf("#%d", i)
		Philosophers[i] = &Philosopher{name, Chopsticks[left], Chopsticks[right]}
	}

	var wg sync.WaitGroup
	wg.Add(Count)
	for i := 0; i < Count; i++ {
		go Philosophers[i].Eat(&wg)
	}
	wg.Wait()
}
