package router

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine)  {
	r := router.Group("/v1/")
	{
		r.GET("health-check", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

	}
}
