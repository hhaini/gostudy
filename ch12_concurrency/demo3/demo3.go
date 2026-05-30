package main

import (
	"fmt"
	"time"
)

/*
channel 是用于 Goroutine 间通信的

对无缓冲 channel 类型的发送与接收操作，一定要放在两个不同的 Goroutine 中进行，否则会导致 deadlock
*/
func main() {
	// 声明一个chanel变量，如果 channel 类型变量在声明时没有被赋予初值，那么它的默认值为 nil
	// var ch chan int
	// 无缓冲 channel
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()
	// ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	n := <-ch1
	println(n)

	// 带缓冲 channel
	ch2 := make(chan int, 5)
	// 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起,但是发送操作不收影响。
	// n2 := <-ch2 //fatal error: all goroutines are asleep - deadlock!
	ch2 <- 2
	go func() {
		n2 := <-ch2
		fmt.Println(n2)
	}()
	time.Sleep(time.Second * 2)
}
