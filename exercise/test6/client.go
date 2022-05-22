package main

import (
	"fmt"
	"net"
)

func connect() {
	conn,err:=net.Dial("tcp", serverAddress)
	if err!=nil{
		fmt.Println("client connect meet an error")
		return
	}
	fmt.Println(conn.RemoteAddr())
	handleIO(conn)
}

func handleIO(conn net.Conn) {
	conn.Write([]byte("Hi!"))
	buffer:=make([]byte, 1024)
	for{
		n,err:=conn.Read(buffer)
		if err!=nil {
			fmt.Println("client read data meet an error")
			return
		}
		content:=string(buffer[:n])
		fmt.Println("client received message "+content)

		// 回复
		conn.Write([]byte("Hi!"))
	}
}

func main() {
	connect()
}