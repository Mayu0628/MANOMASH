package model

import "time"

type Comment struct {
	UserID 		int		  `db:"user_id" form:"user_id" json:"user_id"`
	Comment  	string    `db:"comment" form:"comment" json:"comment"`
	LikeNum   	int       `db:"like_num" form:"like_num" json:"like_num"`
	CreatedAt time.Time   `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" form:"updated_at" json:"updated_at"`
}