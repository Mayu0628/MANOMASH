package model

import "time"

type Oshi struct {
	OshiID		 int	   `db:"oshi_id" form:"oshi_id" json:"oshi_id"`
	OshiName   	 string    `db:"oshi_name" form:"oshi_name" json:"oshi_name"`
	Tag  		 string    `db:"tag" form:"tag" json:"tag"`
	Birthdate  	 time.Date `db:"birthdate" form:"birthdate" json:"birthdate"`
	OshiMeet 	 string    `db:"oshi_meet" form:"oshi_meet" json:"oshi_meet"`
	LikePoint	 string    `db:"like_point" form:"like_point" json:"like_point"`
	Free_space	 string    `db:"free_space" form:"free_space" json:"free_space"`
	Interest	 string    `db:"interest" form:"interest" json:"interest"`
	Reaction_Num int	   `db:"reaction_num" form:"reaction_num" json:"reaction_num"`
	CreatedAt time.Time    `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at" json:"updated_at"`
}

