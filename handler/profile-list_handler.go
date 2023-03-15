package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"

	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
)

type ResData struct {
	Status    int
	Result    string
	UserID    int
	UserName  string
	Introduce string
	OshiName  []string
	OshiID    []int
	//推しのIDを配列にして渡す

}

func MyPageHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}

	var SendID model.User
	var Response ResData
	var OshiData []model.Oshi
	result := database.DB.First(&OshiData,"user_id = ?",id)
	if result.Error != nil {
		ResData := ResFlgCreate(0, "推しデータが見つかりませんでした", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	database.DB.First(&SendID, "user_id = ?", id)
	database.DB.Where("user_id = ?",id).Find(&OshiData)

	//database.DB.First(&OshiData, "oshi_id = ?", OshiData.OshiArray)
	//fmt.Println(OshiData.OshiName)
	//fmt.Println(OshiData.OshiID)
	for i := 0; i < len(OshiData); i++{
		fmt.Println(OshiData[i].OshiID)
		Response.OshiName = append(Response.OshiName, OshiData[i].OshiName)
		Response.OshiID = append(Response.OshiID, OshiData[i].OshiID)
	}
	//fmt.Println(OshiData.OshiArray)

	Response.Status, Response.Result, Response.UserID, Response.UserName, Response.Introduce= 1, "succesful", id, SendID.UserName, SendID.Introduce
	json.NewEncoder(w).Encode(Response)
	return
}