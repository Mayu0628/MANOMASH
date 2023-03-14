package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"

	"net/http"
	"encoding/json"
	"strconv"
)

func MyPageHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData:= ResFlgCreate(0,"fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	type ResData struct{
		Status 		int
		Result 		string
		UserID 		int
		UserName 	string
		Introduce 	string
		OshiName[3] string
		OshiID		int
		//推しのIDを配列にして渡す
	}
	var SendID model.User
	var Response ResData
	var OshiData model.Oshi
	database.DB.First(&SendID, "user_id = ?", id)
	database.DB.First(&OshiData, "user_id = ?", id)
	for i := 1; i < 4; i++ {
		database.DB.First(&OshiData, "oshi_id = ?", i)
		Response.OshiName[i] = OshiData.OshiName
	Response.Status, Response.Result, Response.UserID, Response.UserName, Response.Introduce = 1, "succesful", id, SendID.UserName, SendID.Introduce
	json.NewEncoder(w).Encode(Response)
	}
}
