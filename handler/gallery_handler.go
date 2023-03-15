package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"io"
	"fmt"
	"net/http"
	"encoding/json"
)

func GalleryHandler(w http.ResponseWriter, req *http.Request) {
	var oshidata model.Oshi
	database.DB.Table("oshis").Select([]string{"oshi_id", "oshi_name"}).Find(&oshidata)
	fmt.Println(oshidata)
	json.NewEncoder(w).Encode(oshidata)
	io.WriteString(w, "FetchProfileHandler\n")
}