package main

import (
	"fmt"
	"time"
)

//全局变量，管道
var ch = make(chan int)
var ch2 = make(chan string)

//打印机属于公共资源，会出现多任务资源竞争问题
func Printer(s string)  {
	for _, date := range s {
		fmt.Printf("%c",date)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func person1() {
	Printer("hello")
	ch <- 666
}

func person2() {
	<- ch
	Printer("world")
	ch2 <- "子程序结束"
}
func main() {
	defer fmt.Println("主程序结束")
	//两个协程，代表两个人，公用一个打印机
	go person1()
	go person2()
	str := <-ch2
	fmt.Println(str)

}
