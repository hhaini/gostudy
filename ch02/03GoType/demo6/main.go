package main

import "fmt"

//常量实现枚举
const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift
	starvationThresholdNs = 1e6
)

const (
	Apple, Banana     = iota, iota + 10 // 0, 10 (iota = 0)
	Strawberry, Grape                   // 1, 11 (iota = 1)
	Pear, Watermelon                    // 2, 12 (iota = 2)
)

const (
	i = iota << 1 // 0, iota = 0
	j             // 2, iota = 1
	k             // 4, iota = 2
	l             // 6, iota = 3
	m             // 8, iota = 4
)

func main() {
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs)
	fmt.Println(Apple, Banana, Strawberry, Grape, Pear, Watermelon)
	fmt.Println(i, j, k, l, m)
}
