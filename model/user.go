package model

import (

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  	string    `db:"user_name" form:"user_name" json:"user_name"`
	Email     	string    `db:"email" form:"email" json:"email"`
	Password 	string    `db:"password" form:"password" json:"password"`
	Introduce	string    `db:"introduce" form:"introduce" json:"introduce"`
}

