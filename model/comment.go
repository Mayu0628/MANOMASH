package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	Comment  	string    `db:"comment" form:"comment" json:"comment"`
	LikeNum   	int       `db:"like_num" form:"like_num" json:"like_num"`
	gorm.Model
}