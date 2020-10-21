package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		//重定向跳转
		c.Redirect(http.StatusMovedPermanently,"http://www.sogo.com")//状态码301， 网址
	})

	r.GET("/pinga", func(c *gin.Context) {
		//重定向函数
		c.Request.URL.Path = "/pingb"	//把请求的URL修改
		r.HandleContext(c)	//继续后续的处理
	})

	r.GET("/pingb", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message":"pingb",
		})
	})

	r.Run()
}

