package handler

import (
	"MANOMASH/model"
	"MANOMASH/database"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

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
	addDate := model.Oshi{
		OshiName: reqOshiData.OshiName,
		Birthday: reqOshiData.Birthday,
		OshiMeet: reqOshiData.OshiMeet,
		LikePoint1:reqOshiData.LikePoint1,
		LikePoint2:reqOshiData.LikePoint2,
		LikePoint3:reqOshiData.LikePoint3,
		Free_space:reqOshiData.Free_space,
		Interest: reqOshiData.Interest,
	}
	var SendID model.User
	database.DB.First(&SendID, "user_id = ?", id)
	database.DB.Create(&addDate)
	ResData := ResFlgCreate(1, "succesful", SendID.Id)
	if err := json.NewEncoder(w).Encode(ResData); err != nil {
		fmt.Println(err)
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	return
}