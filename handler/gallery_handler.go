package handler

import (
	"MANOMASH/model"
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)

type ResGelData struct {
	Status    int
	Result    string
	UserID    int
	OshiName  []string
	OshiID    []int
}

func GalleryHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var Response ResData
	var Oshidata []model.Oshi
	//database.DB.Table("oshis").Select([]string{"oshi_id", "oshi_name"}).Find(&oshidata)
	for i := 0; i < 10; i++{
		fmt.Println(Oshidata[i].OshiID)
		Response.OshiName = append(Response.OshiName, Oshidata[i].OshiName)
		Response.OshiID = append(Response.OshiID, Oshidata[i].OshiID)
	}
	Response.Status, Response.Result, Response.UserID= 1, "succesful", id
	json.NewEncoder(w).Encode(Response)
	return
}
