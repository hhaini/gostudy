package main

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// for i := 0; i < 2; i++ {
	go fmt.Println("I am a goroutine!")
	// wg.Done()
	// }
	// wg.Wait()
	var c = make(chan int)
	go func(a, b int) {
		c <- a + b
	}(3, 4)
	fmt.Println(<-c)
	time.Sleep(time.Millisecond * 1)
}
