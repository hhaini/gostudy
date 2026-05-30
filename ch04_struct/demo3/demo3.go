package main

import "fmt"

/*
	使用其他结构体作为自定义结构体中的类型
*/

type Person struct {
	Name  string
	Phone string
	Addr  string
}

type Book struct {
	Title string
	// Author Person
	// 嵌入式字段，也称为匿名字段
	Person
}

func main() {
	var b Book
	fmt.Println(b.Person.Name) // 将类型名当作嵌入字段的名字
	fmt.Println(b.Name)        // 支持直接访问嵌入字段所属类型中字段
}
