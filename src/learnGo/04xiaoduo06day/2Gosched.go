package main

import (
	"fmt"
	"runtime"
)
func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	} ()


	for i := 0; i < 3; i++ {
		//让出时间片，先让别的协程执行，他先执行完再来执行此协程
		runtime.Gosched()
		//本主协程执行的太快，如果没有Gosched，可能会导致主协程执行完了子协程还没开始执行或者执行完
		fmt.Println("hello")
	}
}
