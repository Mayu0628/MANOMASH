package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func ProfileAddHandler(w http.ResponseWriter, req *http.Request) {
	var reqOshiData model.Oshi
	if err := json.NewDecoder(req.Body).Decode(&reqOshiData); err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	addData := model.Oshi{
		OshiName:   reqOshiData.OshiName,
		Birthday:   reqOshiData.Birthday,
		OshiMeet:   reqOshiData.OshiMeet,
		LikePoint1: reqOshiData.LikePoint1,
		LikePoint2: reqOshiData.LikePoint2,
		LikePoint3: reqOshiData.LikePoint3,
		Free_Space: reqOshiData.Free_Space,
		Interest:   reqOshiData.Interest,
	}
	var SendID model.Oshi
	database.DB.First(&SendID, "user_id = ?", reqOshiData.UserID)
	database.DB.Create(&addData)
	fmt.Println(addData)
	ResData := ResFlgCreate(1, "succesful", uint(SendID.OshiID))
	if err := json.NewEncoder(w).Encode(ResData); err != nil {
		fmt.Println(err)
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
}
