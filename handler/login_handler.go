package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"fmt"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		ResData := ResFlgCreate(0,"fail", 0 )
		json.NewEncoder(w).Encode(ResData)
		return
	}
	result := database.DB.Where("email = ?", reqUserData.Email).Where("password = ?", reqUserData.Password).First(&reqUserData)
	if result.Error != nil{
		ResData := ResFlgCreate(0,"fail", 0)
		fmt.Println("エラーです")
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var SendID model.User
	database.DB.First(&SendID, "email = ?", reqUserData.Email)
	ResData := ResFlgCreate(1, "succesful", SendID.Id)
	json.NewEncoder(w).Encode(ResData)
	fmt.Println("ログインに成功しました")
}