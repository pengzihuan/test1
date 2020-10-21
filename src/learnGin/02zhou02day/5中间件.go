package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func indexHandle (c *gin.Context) {
	fmt.Println("indexHandle")
	name,_ :=c.Get("name")
	c.JSON(200,gin.H{
		"name":name,
		"msg":"ok",
	})
}

//满足参数是*Context的都是中间件
func m1(c *gin.Context)  {
	fmt.Println("m1 in...")

	c.Set("name","匿名用户")//通过c的set存储数据在上下文，跨中间件取值

	start := time.Now()
	c.Next()	//调用后续的处理函数,调用完后继续执行下面代码
	//c.Abort()	//阻止调用后续的函数
	// return   //直接结算该函数
	cost := time.Since(start)	//花费了多少时间
	fmt.Println("cost = ",cost)

	fmt.Println("m1 out...")

}

//func m2(c *gin.Context)  {
//	//判断登录
//	//if 是登录用户
//	c.Next()
//	//else
//	c.Abort()
//}

//通常不m2那样写，用m22
func m22() gin.HandlerFunc {
	//连接数据库
	//其他操作
	return func(c *gin.Context) {
			//判断登录
			//if 是登录用户
			//c.Next()
			//else
			//c.Abort()
	}
}

//有开关的，doCheck是开关,可以自行选择检查（true）或者不检查（false）
func m222(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//其他操作
	return func(c *gin.Context) {
		if doCheck{
			//判断登录
			//if 是登录用户
			//c.Next()
			//else
			//c.Abort()
		} else { c.Next() }

	}
}


func main() {
	r := gin.Default()


	//type HandlerFunc func(*Context) ,HandlerFunc就是中间件
	//r.GET("/index",m1, indexHandle)
	//r.GET("/index01",m1, indexHandle)
	//r.GET("/index02",m1, indexHandle)

	//上面这样m1冗余，可以使用r.Use(m1)全局注册m1，所有请求默认添加m1中间件
	r.Use(m1,m222(true))
	r.GET("/index",indexHandle)
	r.GET("/index01",indexHandle)
	r.GET("/index02", indexHandle)

	r.Run()
}
