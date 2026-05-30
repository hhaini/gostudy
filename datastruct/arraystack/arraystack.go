package main

import "fmt"

/*
基于切片实现顺序栈
*/
var items []string //切片
var count, n int   //栈中元素个数，栈大小
// 初始化栈
func arraystack(n1 int) {
	items = make([]string, n1)
	count = 0
	n = n1
}

// 入栈操作
func push(s1 string) bool {
	if count == n {
		return false
	}
	items[count] = s1
	count++
	return true
}

// 出栈操作
func pop() string {
	if count == 0 {
		return ""
	}
	tmp := items[count-1]
	count--
	return tmp
}
func main() {
	arraystack(2)
	push("1")
	push("2")
	res1 := push("3")
	fmt.Println(res1)
	pop()
	pop()
	tmp := pop()
	if len(tmp) != 0 {
		fmt.Println(tmp)
	} else {
		fmt.Println("栈为空")
	}

}
