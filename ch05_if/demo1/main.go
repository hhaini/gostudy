package main

import "fmt"

func main() {
	a, b := false, true
	if a && !b {
		println("(a && b) != true")
		return
	}
	fmt.Println("a && (b != true) == false")
}
