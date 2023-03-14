package handler

import (
	"io"
	"net/http"
	"time"
	"strconv"
)

type ResOshiData struct{
	Status 		int
	Result		string
	Id			int
	OshiName	string
	Birthday	time.Time
	LikePoint1	string
	LikePoint2	string
	LikePoint3	string
	FreeSpace	string
	Interest	string
}

func ProfileHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData:= ResFlgCreate(0,"fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var reqUserData model.User
	var checkData UpdataUser
	database.DB.Where("", reqUserData.Email).Where("password = ?", reqUserData.Password).First(&reqUserData)
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		ResData:= ResFlgCreate(0,"fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
}
