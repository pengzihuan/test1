package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type userinfo struct {
	Username string `form:"username" json:"username"`
	Password string	`form:"password" json:"password"`
}
func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {

	var u userinfo
	err := c.ShouldBind(&u)//绑定u,思考why是&u，传过来的参数可以是json，form表单数据，URL数据等等
		if err != nil {
			c.JSON(400,gin.H{
				"err":err.Error(),
			})
		} else {
			fmt.Println("ok")
			fmt.Printf("u = %v",u)
			c.JSON(200,gin.H{
				"status":"ok",
			})
		}


	})

	r.Run()
}
