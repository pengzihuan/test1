package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn)  {
	//函数调用完，关闭连接
	defer conn.Close()

	//获取用户客户端网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("conncet succesful：",addr)
	buf := make([]byte, 2048)

	for{
		//读取用户数据
		n, err := conn.Read(buf)//nc测试会把换行符\n也读入，自己写键盘输入会多2个 \r\n
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//使用nc的情况
		//if string(buf[:n-1]) == "exit" {
		//	fmt.Println(addr, "exit")
		//	return
		//}

		//自己编写客户端
		if string(buf[:n-2]) == "exit" {
			fmt.Println(addr, "exit")
			return
		}
		//返回客服端数据,把数据转化为大写
		fmt.Println("接受数据：",string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}
func main() {
	//1、监听
	listenr, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer listenr.Close()

	//接受多个用户
	for {
		conn, err1 := listenr.Accept()
		if err1 != nil {
			fmt.Println("err1 = ", err1)
			return
		}
		//用户请求处理,协程
		go HandleConn(conn)
	}
}