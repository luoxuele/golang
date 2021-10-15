package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.Println("hello gin")

	//1.创建一个 默认路由
	r := gin.Default()

	//2. 配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "你好gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "这是一个post，主要用于增加数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个put请求，主要用于编辑数据")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "这是一个delete请求，主要用于删除数据")
	})

	//3.启动http服务，默认0.0.0.0:8080
	r.Run(":9999")
}
