package model

import (
	"time"
)

type Oshi struct {
	OshiID       int       `db:"oshi_id" form:"oshi_id" json:"oshi_id"`
	UserID       int       `db:"user_id" form:"user_id" json:"user_id"`
	OshiName     string    `db:"oshi_name" form:"oshi_name" json:"oshi_name"`
	Birthday     string    `db:"birthday" form:"birthday" json:"birthday"`
	OshiMeet     string    `db:"oshi_meet" form:"oshi_meet" json:"oshi_meet"`
	OshiLike1    string    `db:"oshi_like1" form:"oshi_like1" json:"oshi_like1"`
	OshiLike2    string    `db:"oshi_like2" form:"oshi_like2" json:"oshi_like2"`
	OshiLike3    string    `db:"oshi_like3" form:"oshi_like3" json:"oshi_like3"`
	Free_Space   string    `db:"free_space" form:"free_space" json:"free_space"`
	Interest     string    `db:"interest" form:"interest" json:"interest"`
	Reaction_Num int       `db:"reaction_num" form:"reaction_num" json:"reaction_num"`
	Created_At   time.Time `db:"created_at" form:"created_at" json:"created_at"`
	Updated_At   time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}
