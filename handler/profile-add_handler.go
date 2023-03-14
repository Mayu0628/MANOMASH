package handler

import (
	"MANOMASH/model"
	"MANOMASH/database"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
)

type AddOshiData struct{
		OshiID		int
		Id			uint
		OshiName	string
		Birthday	time.Time
		LikePoint1	string
		LikePoint2	string
		LikePoint3	string
		FreeSpace	string
		Interest	string
	}

func ProfileAddHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData:= ResFlgCreate(0,"fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var reqOshiData model.Oshi
	if err := json.NewDecoder(req.Body).Decode(&reqOshiData); err != nil {
		ResData := ResFlgCreate(0,"fail", 0 )
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var SetOshiData AddOshiData
	database.DB.First(&SetOshiData, "user_id = ?", id)
	database.DB.Create(&reqOshiData)
	ResData := ResFlgCreate(1, "succesful", SetOshiData.Id)
	if err := json.NewEncoder(w).Encode(ResData); err != nil {
		fmt.Println(err)
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	return
}