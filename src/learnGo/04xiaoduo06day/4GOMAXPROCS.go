package main

import (
	"fmt"
	"runtime"
)

func main() {
	m := runtime.GOMAXPROCS(1) //指定几核,返回之前运行是几核
	fmt.Println("m = ", m) //

	for {
		go fmt.Println("1")
		fmt.Println(0)
	}//几核会影响切换频率
}
