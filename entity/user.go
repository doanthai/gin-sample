package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string         `db:"name"`
	Avatar      string         `db:"avatar"`
	Username    string         `db:"username" json:"username" gorm:"unique"`
	Password    string         `db:"password"`
}

func (User) TableName() string {
	return "user"
}
