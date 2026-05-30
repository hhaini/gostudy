package main

import (
	"fmt"
	"unsafe"
)

// 定义一个空结构体,空结构体类型变量的内存占用为 0
// 以空结构体为元素类建立的 channel，是目前能实现的、内存占用最小的 Goroutine 间通信方式。
type Empty struct{}

func main() {
	var s Empty
	fmt.Println(unsafe.Sizeof(s)) //0
	//声明一个元素类型为Empty的channel
	var c = make(chan Empty)
	// 向channel写入一个“事件”
	c <- Empty{}
}
