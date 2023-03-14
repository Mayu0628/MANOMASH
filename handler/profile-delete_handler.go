package handler

import (
	"net/http"
)

func ProfileDeleteHandler(w http.ResponseWriter, req *http.Request) {
	// id, err := strconv.Atoi(req.URL.Query().Get("id"))
	// if err != nil {
	// 	ResData := ResFlgCreate(0, "fail", 0)
	// 	json.NewEncoder(w).Encode(ResData)
	// 	return
	// }
	// database.DB.Where("user_id = ?", id).Where("oshi_id = ?").Delete(&model.Menulist{})
	// ResData := ResFlgCreate(0, "fail", id)
	// json.NewEncoder(w).Encode(ResData)
}
