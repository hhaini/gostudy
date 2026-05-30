package main

//case支持表达式列表
func checkWorkday(a int) {
	switch a {
	case 1, 2, 3, 4, 5:
		println("it is a work day")
	case 6, 7:
		println("it is a weekend day")
	default:
		println("are you live on earth")
	}
}
func main() {
	// a := 7
	// a := 1
	a := 8
	checkWorkday(a)
}
