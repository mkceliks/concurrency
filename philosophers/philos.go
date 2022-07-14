package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

func (p Philo) eat(wg *sync.WaitGroup, hosting chan int) {
	for i := 0; i < 3; i++ {
		hosting <- i
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("<%d> starts eating (this philos %d. eat number) \n", p.number, i+1)
		fmt.Printf("<%d> finishes the (eating number %d) \n", p.number, i+1)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<-hosting
	}
	wg.Done()
}

// func hosting(host chan int, perm chan int, resp chan int) {
// 	for {
// 		hosting, ok := <-host

// 		if ok == false {
// 			fmt.Println(hosting, ok, "Finished.")
// 			break
// 		} else {
// 			fmt.Printf("[%d] wanted perm.\n", hosting)
// 			perm <- hosting
// 		}
// 	}
// }

var wg sync.WaitGroup

func main() {
	// perm := make(chan int)
	// resp := make(chan int)
	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {

		CSticks[i] = new(ChopS)

	}

	philos := make([]*Philo, 5)
	num := 1
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%5],
			number:  num + i,
		}
	}
	host := make(chan int, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, host)
	}
	wg.Wait()
}
