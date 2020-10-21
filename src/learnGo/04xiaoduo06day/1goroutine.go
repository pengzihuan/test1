/*
	1、主协程结束，子协程也结束
*/
package main

import (
	"fmt"
	"time"
)

func NewTask() {
	for {
		fmt.Println("this is a new goroutine")
		time.Sleep(time.Second)		//延迟一秒
	}
}

func main() {

	go NewTask()	//只需要一个go关键字 开一个新协程

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)		//延迟一秒
	}
}
