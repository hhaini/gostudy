package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

/*
零值可用例子:我们不需要对 bytes.Buffer 类型的变量 b 进行任何显式初始化，就可以直接通过处于零值状态的变量 b，调用它的方法进行写入和读取操作
*/
func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello Go"))
	fmt.Println(b.String())
	fmt.Println(unsafe.Sizeof(b)) // 结构体类型变量占用的内存大小
}
