package main

import "fmt"

type Position struct {
	x float64
	y float64
}

func main() {
	//使用复合类型初始化map变量
	m := map[int]string{}
	m[1] = "value1"
	m[2] = "value2"
	fmt.Println(m)

	m1 := map[int][]string{
		1: {"value1_1", "value1_2"},
		3: {"value3_1", "value3_2"},
		7: {"value7_1", "value7_2"},
	}
	fmt.Println(m1)

	m2 := map[Position]string{
		{29.52244, 65.256658}:  "school",
		{14.551122, 25.222656}: "home",
	}
	fmt.Println(m2)

	//使用 make 为 map 类型变量进行显式初始化。
	m1_1 := make(map[int]string)    // 未指定初始容量
	m2_1 := make(map[int]string, 8) // 指定初始容量为8

	m1_1[1] = "val1"
	m2_1[1] = "val2"
	fmt.Println(m1_1, m2_1)

	//获取map键值对数量
	fmt.Println(len(m1_1))

	//注意:我们不能对 map 类型变量调用 cap，来获取当前容量，这是 map 类型与切片类型的一个不同点。

	//判断key是否存在
	m1_2 := map[string]int{}
	m1_2["key1"] = 1
	v, ok := m1_2["key1"]
	if !ok {
		// "key1"不在map中
		fmt.Println("key1 is not in m1_2")
	} else {
		// "key1"在map中，v将被赋予"key1"键对应的value
		fmt.Println(v)
	}
	//delete 函数是从 map 中删除键的唯一方法
	fmt.Println(m1_2)
	delete(m1_2, "key1")
	fmt.Println(m1_2)

	//像对待切片那样通过 for range 语句对 map 数据进行遍历
	//对同一 map 做多次遍历的时候，每次遍历元素的次序都不相同。这是 Go 语言 map 类型的一个重要特点
	//所以千万要注意程序逻辑千万不要依赖遍历 map 所得到的的元素次序
	for k, v := range m1 {
		fmt.Printf("key is %d,value is %s\n", k, v)
	}

}
