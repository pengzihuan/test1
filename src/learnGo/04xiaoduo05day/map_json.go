package main

import "fmt"
import "encoding/json"

func main()  {
	m := make(map[string]interface{}, 4)
	m["company"] = "itcost"
	m["subject"] = []string{"c", "c++", "python", "go"}
	m["isOk"] = true
	m["price"] = 666.666

	buf, err := json.MarshalIndent(m, " ", "")//格式化编码json格式
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}