package main

import (
	"fmt"
	"unsafe"
)

// 判断是否为相同的数组类型
func foo(arr [5]int) {}

func main() {
	var arr1 [5]int
	// var arr2 [6]int
	// var arr3 [5]string
	foo(arr1) // ok
	// foo(arr2) // 错误：[6]int与函数foo参数的类型[5]int不是同一数组类型
	// foo(arr3) // 错误：[5]string与函数foo参数的类型[5]int不是同一数组类型
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度：", len(arr))           // 6
	fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48

	var arr11 [6]int
	fmt.Println(arr11)                         //int类型的0值[0,0,0,0,0,0]
	fmt.Println("数组长度：", len(arr11))           // 6
	fmt.Println("数组大小：", unsafe.Sizeof(arr11)) // 48

	var arr2 = [6]int{
		11, 12, 13, 14, 15, 16,
	} // [11 12 13 14 15 16]
	fmt.Println(arr2)
	var arr3 = [...]int{
		21, 22, 23,
	} // [21 22 23]
	fmt.Printf("%T\n", arr3) // [3]int

	//下标赋值进行初始化
	var arr4 = [...]int{
		99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	}
	fmt.Printf("%T\n", arr4) // [100]int

	var arr5 = [6]int{11, 12, 13, 14, 15, 16}
	fmt.Println(arr5[0], arr5[5]) // 11 16
	// fmt.Println(arr5[-1])         // 错误：下标值不能为负数
	// fmt.Println(arr5[8])          // 错误：小标值超出了arr的长度范围
}
