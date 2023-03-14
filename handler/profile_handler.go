package handler

import (
	"net/http"
	"time"
)

type ResOshiData struct {
	Status    int
	Result    string
	Id        int
	OshiName  string
	Tag       string //ユーザーIDからタグのIDの最大数を取得して、その数ぶんタグをforして取得する
	Birthday  time.Time
	LikePoint string //３回forする
	FreeSpace string
	Interest  string
}

func ProfileHandler(w http.ResponseWriter, req *http.Request) {
	// id, err := strconv.Atoi(req.URL.Query().Get("id"))
	// if err != nil {
	// 	ResData:= ResFlgCreate(0,"fail", 0)
	// 	json.NewEncoder(w).Encode(ResData)
	// 	return
	// }
	// var reqUserData model.User
	// var checkData UpdataUser
	// database.DB.Where("", reqUserData.Email).Where("password = ?", reqUserData.Password).First(&reqUserData)
	// if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
	// 	ResData:= ResFlgCreate(0,"fail", 0)
	// 	json.NewEncoder(w).Encode(ResData)
	// 	return
	// }
}
