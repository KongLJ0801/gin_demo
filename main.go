/*

   @author #Kkk
   @File gin_demo
   @Description:
   @version 0.1
   @date 2022/7/27 16:08

*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get",
		})
	})
	r.GET("/name", func(c *gin.Context) {
		c.String(http.StatusOK, "hole")
	})

	r.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	routerGroup := r.Group("/user")
	{
		routerGroup.GET("/info", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Info",
			})
		})
		routerGroup.GET("/list", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "list",
			})
		})
	}

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8088")
}
