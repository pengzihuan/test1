package main

import "fmt"
import "reflect"

//1、求未知x的类型
func reflectType(x interface{})  {
	//x是未知类型，该方法返回该类型
	obj := reflect.TypeOf(x)
	//fmt.Printf("obj is type %T\n", obj) //obj is type *reflect.rtype
	fmt.Println(obj, obj.Name(), obj.Kind())//Name类型（比较具体），Kind种类（范围比较大）


}

//2、求值
func reflectValue(x interface{})  {
	//获得未知x的值
	v := reflect.ValueOf(x)
	fmt.Printf("值：%v, 类型: %T\n",v,v)//值：2.33, 类型: reflect.Value(需要转换)

	//得到传入值的类型变量
	k := v.Kind()//拿到值对应的类型
	switch k {
	case reflect.Float32:
		//
		ret := float32(v.Float())	//v.float()得到的是float64，还需要类型转换
		fmt.Printf("值：%v, 类型： %T\n",ret, ret)
	}
}
//3、设置值
func reflectSetValue(x interface{})  {
	v := reflect.ValueOf(x)
	//传递的是指针，加了Elem()，就是指针指向的种类
	k := v.Elem().Kind()
	switch k {
	case reflect.Float32:
		v.Elem().SetFloat(6.66)
	}
}



type cat struct {

}
type dog struct {

}

func main() {
	//var a float64
	//reflectType(a)//float64 float64 float64
	//var b int8
	//reflectType(b)//int8 int8 int8
	//var c cat
	//reflectType(c)//main.cat cat struct
	//var d dog
	//reflectType(d)//main.dog dog struct
	////golang中像数组，切片，map，指针等的.name返回都是空
	////一下类型只输出两个，.name为空
	//var e []int
	//reflectType(e)//[]int  slice
	//var f []string
	//reflectType(f)//[]string  slice
	var aa float32 = 2.33
	reflectValue(aa)
	
	var aaa float32 = 2.3
	reflectSetValue(&aaa)
	fmt.Println(aaa)
}
