package tasks

import (
	"fmt"
	"sync"
)

func ProduceSize5(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func ConsumeSize5(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("%d ", num)
	}
}

func Task4() {
	var wg sync.WaitGroup

	numCha := make(chan int, 5)

	wg.Add(2)

	go ProduceSize5(numCha, &wg)
	go ConsumeSize5(numCha, &wg)

	wg.Wait()

	fmt.Println("Done atuh")
}
