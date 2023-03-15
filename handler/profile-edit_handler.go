package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ProfileEditHandler(w http.ResponseWriter, req *http.Request) {
	var reqOshiData model.Oshi
	if err := json.NewDecoder(req.Body).Decode(&reqOshiData); err != nil {
		ResData := ResFlgCreate(0, "デコードに失敗しました。", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	oshi := model.Oshi{
		UserID:     reqOshiData.UserID,
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
	result := database.DB.Model(oshi).Where("user_id = ?", reqOshiData.UserID).Where("oshi_id = ?", reqOshiData.OshiID).Updates(map[string]interface{}{
		"oshi_name":  oshi.OshiName,
		"birthday":   oshi.Birthday,
		"oshi_meet":  oshi.OshiMeet,
		"oshi_like1": oshi.OshiLike1,
		"oshi_like2": oshi.OshiLike2,
		"oshi_like3": oshi.OshiLike3,
		"free_space": oshi.Free_Space,
		"interest":   oshi.Interest,
		"updated_at": time.Now(),
	})
	if result.Error != nil {
		ResData := ResFlgCreate(0, "アップデートに失敗しました", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	ResData := ResFlgCreate(1, "succesful", uint(oshi.OshiID))
	if err := json.NewEncoder(w).Encode(result); err != nil {
		ResData = ResFlgCreate(0, "エンコードに失敗しました", 0)
		json.NewEncoder(w).Encode(ResData)
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(ResData)
	return
}
