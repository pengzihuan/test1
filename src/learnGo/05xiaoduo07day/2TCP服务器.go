package main

import (
	"fmt"
	"net"
)

func main() {
	//1、监听
	listenr, err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer listenr.Close()

	for {
		//2、阻塞，等待用户连接
		conn,err1 := listenr.Accept()

		if err1 != nil {
			fmt.Println("err1 = ", err1)
			continue
		}
		//3、接受用户请求
		buf := make([]byte, 1024)
		n,err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println("err2 = ", err2)
			continue
		}
		fmt.Println("buf = ", string(buf[:n]))
		conn.Close()
	}


}
