package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reqUserData model.User
	var ResData model.ResFlg
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
		ResData.Status = 1
		ResData.Result = "success"
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		// json.NewEncoder(w).Encode(reqUserData)
		// fmt.Println(reqUserData)
		if err := json.NewEncoder(w).Encode(ResData); err != nil {
			fmt.Println(err)
		}
		fmt.Println(ResData)
		// io.WriteString(w, "アカウントが作成されました\n")

		return
	}
	ResData.Status = 0
	ResData.Result = "fail"
	//io.WriteString(w, "そのメールアドレスは既に使用されています\n")
	v, err := json.Marshal(ResData)
	if err != nil {
		fmt.Println(err)
	}
	// json.NewEncoder(w).Encode(ResData)
	w.Write(v)
}
