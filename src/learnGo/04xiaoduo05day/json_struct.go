package main

import (
	"fmt"
	"encoding/json"
)

//通过结构体生产Json，成员变量首字母必须大写(小写读不到数据)，想要改变输出只能二次编码（在后面）
type it struct {
	Company string		`json:"company"`
	Subject []string	`json:"subject"`
	IsOk bool			`json:"isOk"`
	Price float64		`json:"price"`
}
type it01 struct {
	Subject []string	`json:"subject"`
}

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

	var temp it
	err := json.Unmarshal([]byte(jsonBuf), &temp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("temp = ", temp)

	//想要得到指定的，结构体可以定义简单点
	var tmp it01
	err = json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("tmp = ", tmp)
}