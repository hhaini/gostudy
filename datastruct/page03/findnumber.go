package main

import "fmt"

func find(a []int, x int) int {
	pos := -1
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			pos = i
			break
		}
		fmt.Println("i = ", i)
	}
	return pos
}

/*
查找变量i在切片s1中的位置
*/
func findnumber() {
	i := 2
	var s1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	var index1 = find(s1, i)
	fmt.Printf("%d index is %d", i, index1)
}

func main() {
	findnumber()
}
