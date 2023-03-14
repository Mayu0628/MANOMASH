package model

import (
	"gorm.io/gorm"
	"time"
)

type Oshi struct {
	OshiID		 int	    `db:"oshi_id" form:"oshi_id" json:"oshi_id"`
	OshiName   	 string     `db:"oshi_name" form:"oshi_name" json:"oshi_name"`
	Tag  		 string     `db:"tag" form:"tag" json:"tag"`
	TagID  		 int	    `db:"tag_id" form:"tag_id" json:"tag_id"`
	Birthday  	 time.Time  `db:"birthday" form:"birthday" json:"birthday"`
	OshiMeet 	 string     `db:"oshi_meet" form:"oshi_meet" json:"oshi_meet"`
	LikePoint	 string     `db:"like_point" form:"like_point" json:"like_point"`
	Free_space	 string     `db:"free_space" form:"free_space" json:"free_space"`
	Interest	 string     `db:"interest" form:"interest" json:"interest"`
	Reaction_Num int	    `db:"reaction_num" form:"reaction_num" json:"reaction_num"`
	gorm.Model
}

