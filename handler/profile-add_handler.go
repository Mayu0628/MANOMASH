package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ProfileAddHandler(w http.ResponseWriter, req *http.Request) {
	var reqOshiData model.Oshi
	if err := json.NewDecoder(req.Body).Decode(&reqOshiData); err != nil {
		fmt.Println(err)
		ResData := ResFlgCreate(0, "デコードに失敗しました", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	addData := model.Oshi{
		UserID:     reqOshiData.UserID,
		OshiName:   reqOshiData.OshiName,
		Birthday:   reqOshiData.Birthday,
		OshiMeet:   reqOshiData.OshiMeet,
		OshiLike1:  reqOshiData.OshiLike1,
		OshiLike2:  reqOshiData.OshiLike2,
		OshiLike3:  reqOshiData.OshiLike3,
		Free_Space: reqOshiData.Free_Space,
		Interest:   reqOshiData.Interest,
		Created_At: time.Now(),
	}
	var SendID model.Oshi
	result := database.DB.Where("user_id = ?", reqOshiData.UserID).Where("oshi_name = ?", reqOshiData.OshiName).First(&SendID)
	if result.Error != nil {
		result = database.DB.Create(&addData)
		if result.Error != nil {
			fmt.Println(result)
			ResData := ResFlgCreate(0, "推しのデータを登録できませんでした", 0)

			json.NewEncoder(w).Encode(ResData)
			return
		}
		ResData := ResFlgCreate(1, "succesful", uint(SendID.OshiID))
		json.NewEncoder(w).Encode(ResData)
		return
	}

	ResData := ResFlgCreate(0, "既にその推しのプロフィールを作成済みです。", uint(SendID.OshiID))
	if err := json.NewEncoder(w).Encode(ResData); err != nil {
		fmt.Println(err)
		ResData := ResFlgCreate(0, "結果をエンコードできませんでした。", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
}
