package main

import "fmt"

func main() {
	//通用变量声明
	var var1 int = 10
	fmt.Println(var1) //print:10
	//未对变量赋初值
	var var2 int
	fmt.Println(var2) //print:0
	// 声明块声明变量
	var (
		a int    = 128
		b int8   = 6
		s string = "hello"
		c rune   = 'A'
		t bool   = true
	)
	fmt.Println("a=", a, "、b=", b, "、s=", s, "、c=", c, "、t=", t) //print:a= 128 、b= 6 、s= hello 、c= 65 、t= true
	// 多变量声明
	var a1, b1, c1 int = 5, 6, 7
	fmt.Println(a1, b1, c1) //print:5 6 7
	//省略类型变量声明
	var b22 = 13
	fmt.Println(b22) //print:13
	var a2, b2, c2 = 12, 'A', "hello"
	fmt.Println(a2, b2, c2) //print:12 65 hello
	//短变量声明
	a3 := 12
	b3 := 'A'
	c3 := "hello"
	fmt.Printf("%c\n", b3)
	fmt.Println(a3, b3, c3) //print:12 65 hello
	a4, b4, c4 := 12, 'A', "hello"
	fmt.Println(a4, b4, c4) //print:12 65 hello
}
