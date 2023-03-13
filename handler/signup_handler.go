package handler

import (
	"io"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"MANOMASH/model"
	"MANOMASH/database"
)

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	var ResData model.ResFlg
	fmt.Println("サインアップスタート")
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	user := model.User {
		UserName : reqUserData.UserName,
		Email : reqUserData.Email,
		Password : reqUserData.Password,
	}
	result := database.DB.First(&user, "email = ?", reqUserData.Email)
	if result.Error != nil{
		result = database.DB.Create(&user)
		ResData.Status = 1
		ResData.Result = "success"
		if result.Error != nil{
			log.Fatal(result.Error)
		}
		json.NewEncoder(w).Encode(reqUserData)
		fmt.Println(reqUserData)
		if err := json.NewEncoder(w).Encode(ResData); err != nil{
			fmt.Println(err)
		}
		fmt.Println( ResData)
		io.WriteString(w, "アカウントが作成されました\n")

		return
	}
	ResData.Status = 0
	ResData.Result = "fail"
	io.WriteString(w, "そのメールアドレスは既に使用されています\n")
	json.NewEncoder(w).Encode(ResData)
}
