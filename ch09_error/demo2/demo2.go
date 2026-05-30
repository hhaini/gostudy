package main

import (
	"errors"
	"fmt"
)

/*
策略二：哨兵错误处理策略
Go 标准库采用了定义导出的（Exported）“哨兵”错误值的方式，来辅助错误处理方检视（inspect）错误值并做出错误处理分支的决策
*/
var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

func Peek(n int) (int, error) {
	return n, fmt.Errorf("%w", ErrNegativeCount)
}

func main() {
	data, err := Peek(1)
	if errors.Is(err, ErrNegativeCount) {
		fmt.Println("bufio: negative count")
	}
	if err != nil {
		switch err {
		case ErrNegativeCount:
			fmt.Println(data, ErrNegativeCount)
			return
		case ErrBufferFull:
			fmt.Println(data, ErrBufferFull)
			return
		case ErrInvalidUnreadByte:
			fmt.Println(data, ErrInvalidUnreadByte)
			return
		default:
			fmt.Println(data, "other errors")
			return
		}
	}
}

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// var ErrSentinel = errors.New("the underlying sentinel error")

// func main() {
// 	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
// 	err2 := fmt.Errorf("wrap err1: %w", err1)
// 	println(err2 == ErrSentinel) //false
// 	if errors.Is(err2, ErrSentinel) {
// 		println("err2 is ErrSentinel")
// 		return
// 	}
// 	println("err2 is not ErrSentinel")
// }
