package main

import "fmt"
import "time"

func main() {
	//创建一个定时器timer，设置3s后往time通道写内容（内容为当前时间）
	fmt.Println("Nowtime is :", time.Now())
	timer := time.NewTimer(3 * time.Second)//延迟3s，返回时间
	t := <-timer.C	//该channel没有数据会阻塞
	fmt.Println("timer.C = ", t)
}

func main0()  {
	//after底层是用timer实现的，效果差不多
	<-time.After(3 * time.Second)//定时3s，阻塞3s，3s后产生一个事件，往chennel写内容（当前时间）
	fmt.Println("时间到了")
}
func main1()  {
	//sleep也可以实现延迟
	time.Sleep(3 * time.Second)
	fmt.Println("时间到了")
}

