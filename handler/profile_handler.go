package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ProfileHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("id取得失敗")
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var reqOshiData model.Oshi
	result := database.DB.First(&reqOshiData, "oshi_id = ?", id)
	if result.Error != nil {
		ResData := ResFlgCreate(0, "DBからデータが見つかりませんでした", 0)

		json.NewEncoder(w).Encode(ResData)
		return
	}

	oshi := &model.ResOshiData{
		Status:     1,
		Result:     "succes",
		Id:         reqOshiData.OshiID,
		OshiName:   reqOshiData.OshiName,
		Birthday:   reqOshiData.Birthday,
		OshiMeet:   reqOshiData.OshiMeet,
		OshiLike1:  reqOshiData.OshiLike1,
		OshiLike2:  reqOshiData.OshiLike2,
		OshiLike3:  reqOshiData.OshiLike3,
		Free_Space: reqOshiData.Free_Space,
		Interest:   reqOshiData.Interest,
	}
	json.NewEncoder(w).Encode(oshi)
	return

}
