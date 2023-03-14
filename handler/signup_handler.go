package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reqUserData model.User
	fmt.Println("サインアップスタート")
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	user := model.User{
		UserName: reqUserData.UserName,
		Email:    reqUserData.Email,
		Password: reqUserData.Password,
	}
	result := database.DB.First(&user, "email = ?", reqUserData.Email)
	if result.Error != nil {
		result = database.DB.Create(&user)

		ResData := ResFlgCreate(1, "succesful")
		log.Fatal(result.Error)
		if err := json.NewEncoder(w).Encode(ResData); err != nil {
			fmt.Println(err)
		}
		fmt.Println(ResData)

		io.WriteString(w, "アカウントが作成されました\n")

		return
	}
	ResData := ResFlgCreate(0, "fail")
	json.NewEncoder(w).Encode(ResData)
	io.WriteString(w, "そのメールアドレスは既に使用されています\n")
}
