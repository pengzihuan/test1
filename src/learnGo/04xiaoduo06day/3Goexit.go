package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("333")
	runtime.Goexit() //终止该协程
	fmt.Println("444")
}

func main() {
	go func() {
		fmt.Println("111")
		test()
		fmt.Println("222")
	}()

	for {

	}
}
