package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

/*
Go 语言的函数作为“一等公民”，表现出的各种行为特征

*/
// 特征一：Go 函数可以存储在变量中
var (
	myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, a...)
	}
)

// 特征二：支持在函数内创建并通过返回值返回
func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

// 特征三：作为参数传入函数
func three() {
	// 通过 AfterFunc 函数设置了一个 2 秒的定时器，并传入了时间到了后要执行的函数。这里传入的就是一个匿名函数
	time.AfterFunc(time.Second*2, func() { println("timer fired") })
}

// 特征四：拥有自己的类型
// HandlerFunc、visitFunc 就是 Go 标准库中，基于函数类型进行自定义的类型
// type HandlerFunc func(ResponseWriter, *Request)
// type visitFunc func(ast.Node) ast.Visitor

func main() {
	fmt.Printf("%T\n", myFprintf)             // func(io.Writer, string, ...interface {}) (int, error)
	myFprintf(os.Stdout, "%s\n", "Hello, Go") // 输出Hello，Go

	teardown := setup("demo")
	defer teardown()
	println("do some bussiness stuff")

	three()
	time.Sleep(time.Second * 3)
}
