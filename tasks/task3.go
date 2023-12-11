package tasks

import (
	"fmt"
	"sync"
)

func Produce(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func Consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("%d ", num)
	}
}

func Task3() {
	var wg sync.WaitGroup

	numCha := make(chan int)

	wg.Add(2)

	go Produce(numCha, &wg)
	go Consume(numCha, &wg)

	wg.Wait()

	fmt.Println("Done atuh")
}
