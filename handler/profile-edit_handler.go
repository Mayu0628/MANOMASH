package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ProfileEditHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var reqOshiData model.Oshi
	if err := json.NewDecoder(req.Body).Decode(&reqOshiData); err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	oshi := model.Oshi{
		OshiID:     reqOshiData.OshiID,
		OshiName:   reqOshiData.OshiName,
		Birthday:   reqOshiData.Birthday,
		OshiMeet:   reqOshiData.OshiMeet,
		OshiLike1:  reqOshiData.OshiLike1,
		OshiLike2:  reqOshiData.OshiLike2,
		OshiLike3:  reqOshiData.OshiLike3,
		Free_Space: reqOshiData.Free_Space,
		Interest:   reqOshiData.Interest,
	}
	result := database.DB.Model(&reqOshiData).Where("user_id = ?", id).Where("oshi_id = ?").Updates(oshi)
	ResData := ResFlgCreate(1, "succesful", uint(oshi.OshiID))
	if err := json.NewEncoder(w).Encode(result); err != nil {
		ResData = ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(ResData)
}
