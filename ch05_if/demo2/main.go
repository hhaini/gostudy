// if语句的自用变量
package main

func f() int {
	return -1
}
func h() int {
	return -2
}
func main() {
	if a, c := f(), h(); a > 0 {
		println(a)
	} else if b := f(); b > 0 {
		println(a, b)
	} else {
		println(a, b, c)
	}
}
