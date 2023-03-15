package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"fmt"
	"time"

	"encoding/json"
	"net/http"
)

func MyPageEditHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		ResData := ResFlgCreate(0, "デコードに失敗しました", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}

	var SendID model.User
	database.DB.First(&SendID, "user_id = ?", reqUserData.User_ID)
	user := &model.User{
		UserName:  reqUserData.UserName,
		Email:     reqUserData.Email,
		Password:  reqUserData.Password,
		Introduce: reqUserData.Introduce,
	}
	fmt.Println(reqUserData)
	result := database.DB.Model(user).Where("user_id = ?", reqUserData.User_ID).Updates(map[string]interface{}{
		"user_name":  user.UserName,
		"introduce":  user.Introduce,
		"updated_at": time.Now(),
	})
	if result.Error != nil {
		ResData := ResFlgCreate(0, "データを更新できませんでした", 0)

		json.NewEncoder(w).Encode(ResData)
		return
	}

	ResData := ResFlgCreate(1, "succesful", reqUserData.User_ID)
	json.NewEncoder(w).Encode(ResData)
	return
}
