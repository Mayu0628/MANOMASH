package model

import (
	"gorm.io/gorm"
	"time"
)

type Oshi struct {
	id			 int		`db:"id" form:"id" json:"id"`
	OshiID		 int	    `db:"oshi_id" form:"oshi_id" json:"oshi_id"`
	OshiName   	 string     `db:"oshi_name" form:"oshi_name" json:"oshi_name"`
	Birthday  	 time.Time  `db:"birthday" form:"birthday" json:"birthday"`
	OshiMeet 	 string     `db:"oshi_meet" form:"oshi_meet" json:"oshi_meet"`
	LikePoint1	 string     `db:"like_point1" form:"like_point1" json:"like_point1"`
	LikePoint2	 string     `db:"like_point2" form:"like_point2" json:"like_point2"`
	LikePoint3	 string     `db:"like_point3" form:"like_point3" json:"like_point3"`
	Free_space	 string     `db:"free_space" form:"free_space" json:"free_space"`
	Interest	 string     `db:"interest" form:"interest" json:"interest"`
	Reaction_Num int	    `db:"reaction_num" form:"reaction_num" json:"reaction_num"`
	gorm.Model
}

