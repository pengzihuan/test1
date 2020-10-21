package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	jsonBuf :=` {
 "company": "itcost",
 "subject": [
 "c",
 "c++",
 "python",
 "go"
 ],
 "isOk": true,
 "price": 666.666
 }`

	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m)//注意&m
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("m = ", m)
	fmt.Printf("m = %v\n", m)

	str := m["company"]
	fmt.Println("str = ", str)
	fmt.Printf("str 类型 %T", str)
}