package main

//取消了默认执行下一个 case 代码逻辑的语义,Go 提供的关键字 fallthrough 来实现执行下一个 case 的代码逻辑
func case1() int {
	println("eval case1 expr")
	return 2
}
func case2() int {
	println("eval case2 expr")
	return 1
}
func case3() int {
	println("eval case3 expr")
	return 3
}
func switchexpr() int {
	println("eval switch expr")
	return 1
}
func main() {
	switch switchexpr() {
	case case1():
		println("exec case1")
		fallthrough
	case case2():
		println("exec case2")
		fallthrough
	case case3():
		println("exec case3")
		fallthrough
	default:
		println("exec default")
	}
}
