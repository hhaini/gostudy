// for 循环语句的经典形式
package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
	var sum1 int
	for i, j, k := 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k = i+1, j+1, k+5 {
		sum1 += (i + j + k)
		println(sum1)
	}
	sl1 := make([][]int, 5)
	for i := range sl1 {
		sl1[i] = make([]int, i+1)
		sl1[i][0] = 1
		sl1[i][i] = 1
		for j := 1; j < i; j++ {
			sl1[i][j] = sl1[i-1][j-1] + sl1[i-1][j]
		}
	}
	fmt.Println(sl1)
}
