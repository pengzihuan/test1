package main

import "time"
import "fmt"
/*
	Tiker是持续性的（每隔d时间），timer是一次性的
*/
func main() {
	tikcer := time.NewTicker(time.Second)

	i := 0
	for {
		<-tikcer.C
		fmt.Println("i = ", i)
		i++

		if i == 5 {
			//要停掉
			tikcer.Stop()
			break
		}
	}
}
