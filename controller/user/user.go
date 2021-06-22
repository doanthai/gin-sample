package user

import (
	"gin-sample/config/db"
	"gin-sample/entity"
	"gin-sample/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserById(c *gin.Context)  {
	//TODO: Need implementation
}

func GetUsers(c *gin.Context)  {
	//TODO: Need Implementation
}

func CreateUser(c *gin.Context)  {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = util.HashPassword(user.Password)

	_, err := getUserByUserName(user.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Existed user by username"})
		return
	}

	result := db.GetMysql().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context)  {
	//TODO: Need Implementation
}

func DeleteUserById(c *gin.Context)  {
	//TODO: Need Implementation
}
