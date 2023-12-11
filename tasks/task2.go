package tasks

import (
	"fmt"
	"sync"
)

func PrintNumbersWG(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func PrintLettersWG(wg *sync.WaitGroup) {
	defer wg.Done()
	for char := 'a'; char <= 'j'; char++ {
		fmt.Printf("%c ", char)
	}
}

func Task2() {
	var wg sync.WaitGroup

	wg.Add(2)

	go PrintNumbersWG(&wg)
	go PrintLettersWG(&wg)

	wg.Wait()

	fmt.Println("Done atuh")
}
