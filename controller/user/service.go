package user

import (
	"fmt"
	"gin-sample/config/db"
	"gin-sample/entity"
)

func getUserByUserName(username string) (user entity.User, err error) {
	result := db.GetMysql().Where("username = ?", username).Find(&user)
	if result.Error != nil {
		fmt.Println("Error get user by username", result.Error)
	}
	return user, nil
}
