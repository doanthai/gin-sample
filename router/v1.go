package router

import (
	"gin-sample/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine)  {
	v1 := router.Group("/v1/")
	{
		v1.GET("health-check", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		//Users
		v1.GET("users", user.GetUsers)
		v1.GET("users/:id", user.GetUserById)
		v1.POST("users", user.CreateUser)
		v1.PUT("users/:id", user.UpdateUser)
		v1.DELETE("users/:id", user.DeleteUserById)
	}
}
