package done

import (
	"fmt"
	"sync"
)

type Chops struct{ sync.Mutex }
type Philosopher struct {
	number          int
	leftCs, rightCs *Chops
}

type Host struct{ philoSophers []*Philosopher }

func (h Host) startDinner(wg *sync.WaitGroup) {
	ch := make(chan Philosopher, 2)
	for i := 0; i < 5; i++ {
		go h.philoSophers[i].eat(ch)
	}

	for index := 0; index < 15; index++ {
		<-ch
	}
	wg.Done()
}

func (p Philosopher) eat(ch chan<- Philosopher) {
	n := 0

	for n < 3 {
		ch <- p
		p.leftCs.Lock()
		p.rightCs.Lock()
		fmt.Println("Starting to eat ", p.number)
		fmt.Println("Finishing eating ", p.number)
		p.rightCs.Unlock()
		p.leftCs.Unlock()
		n++
	}
}

func main() {
	CSticks := make([]*Chops, 5)
	Philos := make([]*Philosopher, 5)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chops)
	}

	for i := 0; i < 5; i++ {
		Philos[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	host := Host{Philos}
	wg.Add(1)
	go host.startDinner(&wg)
	wg.Wait()
}
