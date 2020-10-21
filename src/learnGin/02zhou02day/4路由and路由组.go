package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//1、不分请求类型，全都访问这个路径
	r.Any("/ping", func(c *gin.Context) {
		//重定向跳转
		c.JSON(http.StatusOK, gin.H{
			"status":"ok",
		})
	})

	//2、不可能写完所有的路径，访问没有的路径会执行这个
	r.NoMethod(func(c *gin.Context) {
		//c.HTML(http.StatusNotFound,"404.html",nil)	//返回一个404的html
		c.JSON(http.StatusOK, gin.H{
			"status":"404",

		})
	})


	//3、路由组（可以嵌套）
	//有很多相同的前缀的路劲，可以分为一个组
	//r.GET("/shop/index", func(c *gin.Context) { })
	//r.GET("/shop/one", func(c *gin.Context) { })
	//r.GET("/shop/two", func(c *gin.Context) { })

	shopGroup := r.Group("/shop")//路由组可以添加中间件
	{//加括号只是为了方便好看
		shopGroup.GET("/index", func(c *gin.Context) { })
		shopGroup.GET("/one", func(c *gin.Context) { })
		shopGroup.GET("/two", func(c *gin.Context) { })
	}

	r.Run()
}

