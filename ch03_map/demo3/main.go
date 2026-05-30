package main

import (
	"fmt"
	"time"
)

// map 实例不是并发写安全的，也不支持并发读写
func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d, %d] ", k, v)
	}
}
func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}
func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()
	time.Sleep(5 * time.Second)
}
