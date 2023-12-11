package tasks

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func PrintLetters() {
	for char := 'a'; char <= 'j'; char++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c ", char)
	}
}

func Task1() {
	go PrintNumbers()
	go PrintLetters()

	time.Sleep(2 * time.Second)
	fmt.Println("Done atuh")
}
