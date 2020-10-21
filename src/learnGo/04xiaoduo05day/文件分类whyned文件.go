/*1
文件：
设备文件：
	屏幕（标准输出设备） fmt.println（）	往标准输出设备写内容
	键盘（标准输入设备） fmt.scan（）		从标准输入设备写内容
磁盘文件：
	文本文件：	以记事本打开，内容1不是乱码
	二进制文件：	记事本打开内容是乱码

whyneed文件？ 放内存，内存掉电丢失，程序结束，内存中内容丢失
			 文件放磁盘，程序结束，文件还在
*/
/*2
	文件常用方法
新建文件: Create()  NewFile()
打开文件: Open()	OpenFile()
写文件:   Write(b []byte)  WriteAt(b byte[], off int64)指定位置写入	WriteString(s string)
读文件：  Read(b []byte)读取数据到b	ReadAt(b []byte, off int64)指定位置读取
删除文件：Remove(name string)
*/
package main

import (
	"fmt"
	"os"
)
func main() {

	//os.Stdout.Close()		//关闭后无法输出
	//标准设备文件（os.Stout）默认已经打开，可直接使用。
	os.Stdout.WriteString("are you ok?\n")

	//os.Stdin.Close()	//关闭后无法输入
	var a int
	fmt.Println("请输入一个int类型")
	fmt.Scan(&a)	//从标准输入设备中读取内容，放在a中
	fmt.Println("a = ", a)

}
