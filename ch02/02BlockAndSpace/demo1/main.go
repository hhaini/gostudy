// var shadowing变量遮蔽例子
package main

import "fmt"

var a = 11

func foo(n int) {
	a := 15
	a += n
}

func main() {
	fmt.Println("a=", a) //print:a=11
	foo(5)
	fmt.Println("after foo a = ", a) //print:after foo a =  11
}
