/*
接口类型嵌入,I1等价于I2
接口类型嵌入的语义就是新接口类型（如接口类型 I2）将嵌入的接口类型（如接口类型 E）的方法集合，并入到自己的方法集合中。
通过在接口类型中嵌入其他接口类型可以实现接口的组合，这也是 Go 语言中基于已有接口类型构建新接口类型的惯用法
*/
package main

import (
	"fmt"
	"io"
	"strings"
)

type E interface {
	M1()
	M2()
}
type I1 interface {
	M1()
	M2()
	M3()
}
type I2 interface {
	E
	M3()
}

type myInt int

func (n *myInt) Add(m int) {
	*n = *n + myInt(m)
}

type t struct {
	a int
	b int
}

type s struct {
	*myInt
	t
	io.Reader
	// s string
	// n int
}

func main() {
	m := myInt(17)
	r := strings.NewReader("hello go")
	s :=
		s{
			myInt: &m,
			t: t{
				a: 1,
				b: 2,
			},
			Reader: r,
			// s:      "demo",
		}
	var s1 = make([]byte, len("hello world"))
	fmt.Printf("s1: %T\n", s1)
	s.Reader.Read(s1)
	fmt.Printf("string(s1): %v\n", string(s1))
	s.myInt.Add(5)
	fmt.Println(*(s.myInt))
}
