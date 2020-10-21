package main

import (
	"fmt"
	"time"
)
func main() {
	//给定以一个缓冲区容量>0，通道就是异步的。满了才能暂时同步
	ch := make(chan int, 3)//有缓冲区的channel，能存3个值
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n",len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("子协程[%d]，len(ch) = %d, cap(ch) = %d\n",i,len(ch), cap(ch))
		}
	}()

	time.Sleep(2*time.Second)
	for i := 0; i < 3; i++ {
		m := <-ch
		fmt.Println("m = ", m)
	}
}
