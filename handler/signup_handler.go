package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reqUserData model.User
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}

	var SendID model.User
	result := database.DB.First(&SendID, "email = ?", reqUserData.Email)
	if result.Error != nil {
		result = database.DB.Create(&reqUserData)
		ResData := ResFlgCreate(1, "succesful", SendID.Id)

		if err := json.NewEncoder(w).Encode(ResData); err != nil {
			fmt.Println(err)
			ResData := ResFlgCreate(0, "fail", 0)
			json.NewEncoder(w).Encode(ResData)
			return
		}
		return
	}
	ResData := ResFlgCreate(0, "fail", 0)
	json.NewEncoder(w).Encode(ResData)
	return
}
