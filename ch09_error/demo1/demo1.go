package main

import (
	"errors"
	"fmt"
)

/*
Go 语言中最常见的错误处理策略
err := doSomething()
if err != nil {
// 不关心err变量底层错误值所携带的具体上下文信息
// 执行简单错误处理逻辑并返回
... ...
return err
}

func doSomething(...) error {
... ...
return errors.New("some error occurred")
}
*/
func main() {
	//go语言中两种定义错误的方法：errors.New()、fmt.Errorf("")
	err := errors.New("this is error return")
	fmt.Println(err)
	i := 2
	err1 := fmt.Errorf("index %d is out of bounds", i)
	fmt.Println(err1)
	// 策略一：透明处理方式
	err2 := doSomething()
	if err2 != nil {

		fmt.Printf("err2: %v\n", err2)
	}
}
func doSomething() error {
	return errors.New("some error occurred")
}
