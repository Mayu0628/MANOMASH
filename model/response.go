package model

import "time"

type ResFlg struct {
	Status int    `json:"status"`
	Result string `json:"result"`
	Id     uint   `json:"id"`
}

type ResOshiData struct {
	Status     int
	Result     string
	Id         int
	OshiName   string
	Birthday   time.Time
	OshiMeet   string
	LikePoint1 string
	LikePoint2 string
	LikePoint3 string
	Free_Space string
	Interest   string
}
