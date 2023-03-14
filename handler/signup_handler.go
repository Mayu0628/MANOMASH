package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"io"
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

	result := database.DB.First(&reqUserData, "email = ?", reqUserData.Email)
	if result.Error != nil {
		result = database.DB.Create(&reqUserData)

		ResData := ResFlgCreate(1, "succesful")
		if err := json.NewEncoder(w).Encode(ResData); err != nil {
			ResData := ResFlgCreate(0, "fail")
			json.NewEncoder(w).Encode(ResData)
			fmt.Println(err)
			return
		}
		fmt.Println(ResData)

		io.WriteString(w, "アカウントが作成されました\n")

		return
	}
	ResData := ResFlgCreate(0, "fail")
	json.NewEncoder(w).Encode(ResData)
	io.WriteString(w, "そのメールアドレスは既に使用されています\n")
	return
}
