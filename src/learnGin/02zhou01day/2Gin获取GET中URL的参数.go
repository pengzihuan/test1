package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//1、获取浏览器那边请求携带的 query string 参数

		//几种获得GET请求参数方法,多个 key=value 用&连接
		name := c.Query("query")
		age := c.Query("age")
		//name := c.DefaultQuery("query","defaultValue") //比Query()多一个功能，没有取到会选择默认的值

		/*name,ok := c.GetQuery("query") //取到返回(值，true),取不到("",false) “”空字符串
		if !ok {
			name = "没取到"
		}*/

		c.JSON(200,gin.H{
			"name":name,
			"age":age,
		})
	})
	r.Run()
}