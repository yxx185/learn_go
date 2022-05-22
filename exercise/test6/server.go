package main

import (
	"fmt"
	"net"
)
// 服务器地址：端口
var serverAddress = "localhost:1234"

func start() {
	acceptor,err := net.Listen("tcp", serverAddress)
	if err!=nil {
		fmt.Println("start server meet an error")
		return
	}

	// 延迟关闭资源，避免内存泄漏
	defer acceptor.Close()

	// 死循环监听处理
	for{
		conn,er := acceptor.Accept()
		if er!=nil {
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	fmt.Println("receive conn "+ conn.RemoteAddr().String())
	// 缓冲区
	buffer:=make([]byte, 1024)
	// 死循环读写数据
	for{
		n,err:=conn.Read(buffer)
		if err!=nil {
			fmt.Println("server read data meet an error")
			return
		}
		fmt.Println("receive client send "+ string(buffer[:n]))

		// 回复
		words:="hello,I'm server"
		conn.Write([]byte(words))

	}
}

func main() {
	start()
}