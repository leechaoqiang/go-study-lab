package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello Gin App")
	})

	//GET请求-path
	r.GET("/login/:name/:password", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			//{ name: "15949629528", password: "123456" }
			"name":     c.Param("name"),
			"password": c.Param("password"),
		})
	})

	//GET请求-query
	r.GET("/login", func(c *gin.Context) {
		//{ name: "张三", password: "123456" }
		c.JSON(http.StatusOK, gin.H{
			"name":     c.Query("name"),
			"password": c.Query("password"),
		})
	})

	//POST请求
	r.POST("/login", func(c *gin.Context) {
		//{"name":"张三","password":"123456"}
		c.JSON(http.StatusOK, gin.H{
			"name":     c.Query("name"),
			"password": c.Query("password"),
		})
	})

	//GET请求-默认值
	r.GET("/user", func(c *gin.Context) {
		//{ name: "张三", password: "123456" }
		c.JSON(http.StatusOK, gin.H{
			"name":     c.DefaultQuery("name", "默认张三"),
			"password": c.DefaultQuery("password", "默认密码"),
		})
	})
	//POST请求-默认值
	r.POST("/user", func(c *gin.Context) {
		//{"name":"张三","password":"默认密码"}
		c.JSON(http.StatusOK, gin.H{
			"name":     c.DefaultQuery("name", "默认张三"),
			"password": c.DefaultQuery("password", "默认密码"),
		})
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
