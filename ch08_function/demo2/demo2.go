package main

import (
	"errors"
	"fmt"
)

/*

 */
// 函数返回值列表从形式上看主要有三种：
// 无返回值
func foo() {}

// 仅有一个返回值,如果一个函数仅有一个返回值，那么通常我们在函数声明中，就不需要将返回值用括号括起来
func foo1() error {
	return errors.New("this is one error")
}

// 有2或2个以上返回值,如果是 2 个或 2 个以上的返回值，那我们还是需要用括号括起来的
func foo2() (int, string, error) {
	return int(1), string("sss"), errors.New("this is two errors")
}
func main() {
	foo()
	fmt.Println(foo1())
	fmt.Println(foo2())
}
