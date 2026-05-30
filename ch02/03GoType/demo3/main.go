package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}
func main() {
	var ss = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&ss)) // 将string类型变量地址显式转型为reflect.StringHeader
	fmt.Printf("0x%x\n", hdr.Data)                      // 0x10a30e0
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))           // 获取Data字段所指向的数组的指针
	dumpBytesArray((*p)[:])                             // [h e l l o ] // 输出底层数组的内容
	var s1 = "中国人"
	fmt.Printf("0x%x\n", s1[0]) // 0xe4：字符“中” utf-8编码的第一个字节
	//经过常规 for 迭代后，我们获取到的是字符串里字符的 UTF-8 编码中的一个字节
	for i := 0; i < len(s1); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, s1[i])
	}
	//通过 for range 迭代，我们每轮迭代得到的是字符串中 Unicode 字符的码点值
	for i, v := range s1 {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}
	//字符串连接
	s := "Rob Pike, "
	s = s + "Robert Griesemer, "
	s += " Ken Thompson"
	fmt.Println(s) // Rob Pike, Robert Griesemer, Ken Thompson
}
