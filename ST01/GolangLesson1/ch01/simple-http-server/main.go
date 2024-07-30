package main

import (
	"net/http"
)

// 处理函数处理客户端发送的请求，r代表来自客户端的请求，w是用来操作返回给客户端的应答
func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))

}

func main() {
	// http.HandleFunc设置处理函数并传入模式字符串为"/"
	http.HandleFunc("/", MyHandler)
	// 通过http包提供的ListenAndServe函数，建立起一个HTTP服务，这个服务监听本地的8081端口
	http.ListenAndServe(":8081", nil)
}
