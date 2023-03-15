package model


type Oshis struct{
	OshiArray []Oshi `db:"oshi_array" form:"oshi_array" json:"oshi_array"`
}