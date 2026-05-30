// 切片
package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(nums)) // 6
	nums = append(nums, 7) // 切片变为[1 2 3 4 5 6 7]
	fmt.Println(len(nums)) // 7
	//通过 make 函数来创建切片，并指定底层数组的长度
	sl := make([]byte, 6, 10)      // 其中10为cap值，即底层数组长度，6为切片的初始长度
	sl1 := make([]byte, 6)         // cap = len = 6
	fmt.Println(cap(sl), cap(sl1)) //10 6

	//采用 array[low : high : max]语法基于一个已存在的数组创建切片。这种方式被称为数组的切片化
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 := arr[3:7:9]
	fmt.Println(len(sl2), cap(sl2)) //4 6 len = high - low cap = max - low
	sl2[0] += 10
	fmt.Println("arr[3] =", arr[3]) // 14

	//基于切片创建切片
	sl3 := nums[:]
	fmt.Println(sl3)

	//切片动态扩容
	var s []int
	s = append(s, 11)
	fmt.Println(len(s), cap(s)) //1 1
	s = append(s, 12)
	fmt.Println(len(s), cap(s)) //2 2
	s = append(s, 13)
	fmt.Println(len(s), cap(s)) //3 4
	s = append(s, 14)
	fmt.Println(len(s), cap(s)) //4 4
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) //5 8

	var sl12 []int
	var sl21 = []int{}
	fmt.Println(len(sl12), cap(sl12))
	fmt.Println(len(sl21), cap(sl21))
}
