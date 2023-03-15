package handler

import (
	"MANOMASH/model"

	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
)

type ResData2 struct {
	Status    int
	Result    string
	UserID    int
	OshiName  []string
	OshiID    []int
}

func ProfileListHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var Response ResData
	var OshiData []model.Oshi
	for i := 0; i < len(OshiData); i++{
		fmt.Println(OshiData[i].OshiID)
		Response.OshiName = append(Response.OshiName, OshiData[i].OshiName)
		Response.OshiID = append(Response.OshiID, OshiData[i].OshiID)
	}
	Response.Status, Response.Result, Response.UserID = 1, "succesful", id
	json.NewEncoder(w).Encode(Response)
}