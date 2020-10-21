package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	//1、创建默认引擎
	r := gin.Default()

	//2、访问操作
	r.GET("/ping", func(c *gin.Context) {

		//date := map[string]interface{}{
		//	"name":"samllking",
		//	"gae":18,
		//}

		//gin.H 就是map[string]interface{}
		date := gin.H{"name":"samllking","gae":18}
		//返回json格式的数据
		c.JSON(200,date)
	})

	type person struct {
		Name string		`json:"name"` //小写会访问不到，要小写可以用tag
		Age int			`json:"age"`
	}
	
	r.GET("/ping2", func(c *gin.Context) {
		date := person{
			Name: "samll king",
			Age:  18,
		}
		c.JSON(200,date)	//json的序列化
	})
	
	//3、运行
	r.Run()
}
