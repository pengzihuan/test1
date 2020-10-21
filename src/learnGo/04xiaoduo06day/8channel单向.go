package main

func main() {
	//双向通道
	//双向channel能隐式转换为单向channel
	ch := make(chan int)
	//只能写进管道
	var writech chan<- int = ch
	writech <- 666
	//只能从管道读
	var readch <-chan int = ch
	<- readch
}
