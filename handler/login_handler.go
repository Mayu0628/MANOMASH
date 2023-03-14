package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"fmt"
	"encoding/json"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	user := model.User {
		Email : reqUserData.Email,
		Password : reqUserData.Password,
	}
	result := database.DB.Where("email = ?", reqUserData.Email).Where("password = ?", reqUserData.Password).First(&user)
	fmt.Println(reqUserData.Email,reqUserData.Password)
	if result.Error != nil{
		ResData := ResFlgCreate(0,"fail")
		fmt.Println("エラーです")
		json.NewEncoder(w).Encode(ResData)
		return
	}
	ResData := ResFlgCreate(1,"succesful")
	fmt.Println("ログインに成功しました")
	expiration := time.Now()
    expiration = expiration.AddDate(0, 0, 1)
    cookie := http.Cookie{Name: reqUserData.Email, Value: reqUserData.Email, Expires: expiration}
	fmt.Println(cookie)
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(ResData)
}