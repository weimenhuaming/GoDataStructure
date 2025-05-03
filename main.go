package main

import (
	"fmt"
	"net/http"
)

// 处理器函数
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}

func main() {
	//设置多路复用器：
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	//搭建服务器
	err := http.ListenAndServe(":8080", mux)
	//参数为1.网络地址和2.负责处理请求的处理器。
	//如果不采用的话就是一个mux，就是nil，就会使用默认的多路复用器，其他均为http.

	if err != nil {
		fmt.Printf("failed ,err:%v\n", err)
	}
	//上一句的代码还可以使用server结构（也就是服务器的搭建啦）
	/*   其实就是一个扩展内容，除了下面的参数,还有其他参数
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
	*/
}
