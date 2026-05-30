package main

import (
	"fmt"
	"time"
)

func main() {
	sl := []int{1, 2, 3, 4, 5}
	// for k, v := range sl {
	// 	fmt.Println(k, v)
	// }
	for k, v := range sl {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(k, v)
		}()
	}
	time.Sleep(time.Second * 10)
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func(i, v int) {

			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 5)
}
