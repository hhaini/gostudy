package main

import "fmt"

func main() {
	//字符串常量定义
	const (
		GO_SLOGAN = "less is more"  // GO_SLOGAN是string类型常量
		s1        = "hello, gopher" // s1是string类型常量
	)
	//字符串变量定义
	var s2 = "I love go" // s2是string类型变量
	fmt.Println(s1)
	fmt.Println(s2)
}
