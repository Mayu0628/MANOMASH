package model

type ResFlg struct {
	Status int    `json:"status"`
	Result string `json:"result"`
	Id     uint   `json:"id"`
}

type ResOshiData struct {
	Status     int
	Result     string
	OshiId     int
	OshiName   string
	Birthday   string
	OshiMeet   string
	OshiLike1  string
	OshiLike2  string
	OshiLike3  string
	Free_Space string
	Interest   string
}
