package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)       // 设置根路径的处理函数
	err := http.ListenAndServe(":8080", nil) // 监听并且提供服务于8080端口
	if err != nil {
		fmt.Println("服务器启动失败: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // 返回响应
}
