package tasks

import (
	"fmt"
	"sync"
)

func SendNumbersWithErr(evenCh, oddCh chan<- int, errCh chan<- error, wg *sync.WaitGroup) {
	defer close(evenCh)
	defer close(oddCh)
	defer close(errCh)
	defer wg.Done()

	for i := 1; i <= 25; i++ {
		if i > 20 {
			errCh <- fmt.Errorf("Error: number %d is greater than 20", i)
		} else {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddCh <- i
			}
		}
	}
}

func Task6() {
	var wg sync.WaitGroup

	evenCh := make(chan int)
	oddCh := make(chan int)
	errCh := make(chan error)

	wg.Add(1)

	go SendNumbersWithErr(evenCh, oddCh, errCh, &wg)

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
		case err, ok := <-errCh:
			if !ok {
				fmt.Println("Error channel closed!")
				errCh = nil
			} else {
				fmt.Println(err)
			}
		}

		if evenCh == nil && oddCh == nil && errCh == nil {
			break
		}
	}

	wg.Wait()

	fmt.Println("Done atuh")
}
