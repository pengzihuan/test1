package main

import "fmt"
import "reflect"

type student struct {
	Name string		`json:"name"`
	Score int		`json:"score"`
}

func main() {
	stu1 := student{"tom", 99}
	//通过反射获取结构体所有字段
	t := reflect.TypeOf(stu1)
	fmt.Printf("类型：%v, name: %v, kind: %v\n",t,t.Name(),t.Kind())
	//遍历所有字段
	for i := 0; i < t.NumField(); i++ {
		//根据结构体的字段的索引去获取值
		sField := t.Field(i)
		fmt.Printf("name: %v, type: %v, tag: %v", sField.Name, sField.Type, sField.Tag)
	}
}
