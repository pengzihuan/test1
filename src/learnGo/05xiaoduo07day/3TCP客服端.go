package main

import "fmt"
import "net"

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	buf := make([]byte,1024)
	//向服务器写数据
	fmt.Println("input:")
	fmt.Scan(&buf)
	conn.Write(buf)
}
