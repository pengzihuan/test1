package main

import (
	"fmt"
	"os"
)
import "net"

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	//从键盘读取数据
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
			//把输入的内容发送给服务器
			conn.Write(str[:n])
		}
	}()


	//接受服务器的数据
	buf := make([]byte, 1024)

	for {
		n, err1 :=conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1 = ", err1)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
