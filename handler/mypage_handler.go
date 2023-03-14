package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"

	"encoding/json"
	"net/http"
	"strconv"
)

func MyPageHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	type ResData struct {
		Status    int
		Result    string
		UserID    int
		UserName  string
		Introduce string
		OshiName  string
	}
	var SendID model.User
	var Response ResData
	var OshiData model.Oshi
	database.DB.First(&SendID, "user_id = ?", id)
	database.DB.First(&OshiData, "user_id = ?", id)
	Response.Status, Response.Result, Response.UserID, Response.UserName, Response.Introduce, Response.OshiName = 1, "succesful", id, SendID.UserName, SendID.Introduce, OshiData.OshiName
	json.NewEncoder(w).Encode(Response)
}
