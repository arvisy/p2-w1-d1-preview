package tasks

import (
	"fmt"
	"sync"
)

func SendNumbers(evenCh, oddCh chan<- int, wg *sync.WaitGroup) {
	defer close(evenCh)
	defer close(oddCh)
	defer wg.Done()

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
}

func Task5() {
	var wg sync.WaitGroup

	evenCh := make(chan int)
	oddCh := make(chan int)

	wg.Add(1)

	go SendNumbers(evenCh, oddCh, &wg)

	for {
		select {
		case evenNum, ok := <-evenCh:
			if !ok {
				fmt.Println("Even channel closed!")
				evenCh = nil
			} else {
				fmt.Print("Received an even number: ", evenNum)
				fmt.Println()
			}

		case oddNum, ok := <-oddCh:
			if !ok {
				fmt.Println("Odd channel closed!")
				oddCh = nil
			} else {
				fmt.Print("Received an odd number: ", oddNum)
				fmt.Println()
			}
		}

		if evenCh == nil && oddCh == nil {
			break
		}
	}

	wg.Wait()

	fmt.Println("Done atuh")
}
