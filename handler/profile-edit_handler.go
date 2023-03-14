package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
)

func ProfileEditHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData:= ResFlgCreate(0,"fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var reqOshiData model.Oshi
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		ResData:= ResFlgCreate(0,"fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	oshi := model.Oshi {
		OshiID : reqOshiData.OshiID,
		OshiName : reqOshiData.OshiName,
		Birthday : reqOshiData.Birthday,
		OshiMeet : reqOshiData.OshiMeet,
		LikePoint1 : reqOshiData.LikePoint1,
		LikePoint2 : reqOshiData.LikePoint2,
		LikePoint3 : reqOshiData.LikePoint3,
		FreeSpace : reqOshiData.FreeSpace,
		Interest : reqOshiData.Interest,
	}
	database.DB.Model(&SendID).Where("user_id = ?", id).Where("oshi_id = ?", )Updates(oshi)
	ResData := ResFlgCreate(1,"succesful",id)
	if err := json.NewEncoder(w).Encode(result); err != nil{
		ResData = ResFlgCreate(0,"fail")
		json.NewEncoder(w).Encode(ResData)
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(ResData)
}