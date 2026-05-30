package main

func main() {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1
	// find first even number of the interger slice
	// for i := 0; i < len(sl); i++ {
	// switch sl[i] % 2 {
	// case 0:
	// 	firstEven = sl[i]
	// 	break
	// case 1:
	// 	// do nothing
	// }
	// if sl[i]%2 == 0 {
	// 	firstEven = sl[i]
	// 	break
	// }
	// }
	// println(firstEven)//12
loop:
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			break loop
		case 1:
			// do nothing
		}
	}
	println(firstEven) //6
}
