package main

import "fmt"
import "time"

func main() {
	fmt.Println("Nowtime is :", time.Now())
	timer := time.NewTimer(30 * time.Second)//延迟30s，返回时间
	//Reset重置延迟1s
	ok := timer.Reset(time.Second)
	fmt.Println("ok = ", ok)

	fmt.Println("时间到了,time = ",<-timer.C)
}

func main2()  {
	timer := time.NewTimer(5 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("时间到了，子协程可以打印了")
	}()
	//停止定时器，时间没到提前停止，没有向channel输入数据，<-timer.C会阻塞，协程无法正常执行后面
	timer.Stop()
	for{

	}
}