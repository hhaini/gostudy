package main

import "fmt"

/*
panicking过程
*/
func foo() {
	println("call foo")
	bar()

	println("exit foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic :", e)
		}
	}()
	println("call bar")
	panic("panic occurs in bar")
	// zoo()
	// println("exit in bar")
}

// func zoo() {
// 	println("call zoo")
// 	println("exit zoo")
// }
func main() {
	println("call main")
	foo()
	println("exit main")
}
