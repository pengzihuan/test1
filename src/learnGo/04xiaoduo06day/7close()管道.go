package main

import "fmt"
//channle为nil,接发方都堵塞
func main01() {
	ch := make(chan int)
	//新建一个协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i	//往管道写数据
		}
		close(ch)
	}()

	for {
		if n, ok := <- ch; ok == true {
			fmt.Println("n = ",n)
		} else {
			fmt.Println("管道关闭")
			break
		}
	}
}

func main() {
	ch := make(chan int)
	//新建一个协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i	//往管道写数据
		}
		close(ch)
	}()

	//接收方可以range迭代更简便
	for n := range ch {
		fmt.Println("n = ", n)
	}
}