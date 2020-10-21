package main

import (
	"fmt"
	"encoding/json"
)

//通过结构体生产Json，成员变量首字母必须大写(小写读不到数据)，想要改变输出只能二次编码（在后面）
type iT struct {
	Company string		`json:"-"`//此字段不会输出到屏幕
	Subject []string	`json:"subject"`//通过二次编码首字母小写
	IsOk bool			`json:",string"`//通过二次编码转类型
	Price float64
}

func main() {
	s := iT{"itcost", []string{"c", "c++", "python", "go"}, true, 666.666}
	//编码，根据内容生成json文本，返回一个切片，一个错误
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, " ", "")//格式化编码，第一个内容，第二个空格或无，第三个tab缩进。与上面不同的是格式看起来好些。
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("buf = ", string(buf))
}
