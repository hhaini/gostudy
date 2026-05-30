package main

type person struct {
	name string
	age  int
}

func main() {
	p := person{"tom", 13}
	switch p {
	case person{"tony", 33}:
		println("match tony")
	case person{"tom", 13}:
		println("match tom")
	case person{"lucy", 23}:
		println("match lucy")
	default:
		println("no match")
	}
}
