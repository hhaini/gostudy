package main

import "fmt"

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
	b := t.a
	fmt.Printf("b: %v\n", b)
}
func (t *T) M2() {
	t.a = 11
}
func main() {
	var t T
	println(t.a) // 0
	t.M1()
	println(t.a) // 0
	t.M2()
	fmt.Printf("t.a: %v\n", t.a) //11

	fmt.Printf("t: %T\n", t)
	var t2 = &T{}
	fmt.Printf("t2: %T\n", t2)
	println(t2.a) // 0
	t2.M1()
	println(t2.a) // 0
	t2.M2()
	println(t2.a) // 11
}
