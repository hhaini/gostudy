package main

import "fmt"

//map 变量被传递到函数或方法内部后，我们在函数内部对map类型参数的修改在函数外部也是可见的
func foo(m map[string]int) {
	m["key1"] = 11
	m["key2"] = 12
}

func main() {
	m := make(map[string]int, 8)
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println(m)
	foo(m)
	fmt.Println(m)
}
