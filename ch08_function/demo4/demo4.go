package main

import (
	"fmt"
	"net/http"
)

/*
函数也可以被显式转型
典型的示例:标准库 http 包中的 HandlerFunc 这个类型
*/
func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}
func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}
