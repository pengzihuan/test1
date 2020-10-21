package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	//加载html文档
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		//给游览器返回一个html文档
		c.HTML(200,"login.html",nil)
	})

	//login post请求
	r.POST("/login", func(c *gin.Context) {
		//第一种获取参数way,
		username := c.PostForm("username")//取得到返回值取不到返回空字符串
		password := c.PostForm("password")

		//第二种,取不到有默认值
		//username := c.DefaultPostForm("username","default")
		//password := c.DefaultPostForm("password","default")

		//第三钟返回两个值，取不到返回false
		//username,ok1 := c.GetPostForm("username")
		//password,ok2 := c.GetPostForm("password")
		//if !ok1 {}
		//if !ok2 {}

		c.HTML(200,"index.html",gin.H{
			"Username" : username,
			"Password" : password,
		})
	})

	r.Run()
}
