package model

import "time"

type Comment struct {
	LikeNum   	int       `db:"like_num" form:"like_num" json:"like_num"`
	Comment  	string    `db:"comment" form:"comment" json:"comment"`
	UserID 		int		  `db:"user_id" form:"user_id" json:"user_id"`
	CreatedAt time.Time   `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" form:"updated_at" json:"updated_at"`
}