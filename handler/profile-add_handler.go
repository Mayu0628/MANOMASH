package handler

import (
	"MANOMASH/model"
	"MANOMASH/database"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

func ProfileAddHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData:= ResFlgCreate(0,"fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var reqOshiData model.Oshi
	if err := json.NewDecoder(req.Body).Decode(&reqOshiData); err != nil {
		ResData := ResFlgCreate(0,"fail", 0 )
		json.NewEncoder(w).Encode(ResData)
		return
	}
	database.DB.First(&reqOshiData, "user_id = ?", id)
	database.DB.Create(&reqOshiData)
	ResData := ResFlgCreate(1, "succesful", id)

	if err := json.NewEncoder(w).Encode(ResData); err != nil {
		fmt.Println(err)
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	return
	}
	ResData := ResFlgCreate(0, "fail", 0)
	json.NewEncoder(w).Encode(ResData)
	return
}