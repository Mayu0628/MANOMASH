package handler

import (
//	"MANOMASH/database"
//	"MANOMASH/model"
//
//	"fmt"
//	"encoding/json"
	"net/http"
)

func ProfileEditHandler(w http.ResponseWriter, req *http.Request) {
//	var reqOshiData model.Oshi
//	oshi := model.Oshi {
//		OshiID : reqOshiData.OshiID,
//		OshiName : reqOshiData.OshiName,
//		Tag : reqOshiData.Tag,
//		Birthday : reqOshiData.Birthday,
//		OshiMeet : reqOshiData.OshiMeet,
//		LikePoint : reqOshiData.LikePoint,
//		FreeSpace : reqOshiData.FreeSpace,
//		Interest : reqOshiData.Interest,
//	}
//	cookies := req.Cookies()
//	for _, c := range cookies {
//		birthday := c.Value
//		result := oshi
//		database.DB.Raw("SELECT oshi_id, oshi_name, tag, birthday, oshi_meet, like_point, free_space, interest FROM oshis WHERE birthday = ?", birthday).Scan(&result)
//		ResData := ResFlgCreate(1,"succesful")
//		fmt.Println(result)
//		if err := json.NewEncoder(w).Encode(result); err != nil{
//			ResData = ResFlgCreate(0,"fail")
//			json.NewEncoder(w).Encode(ResData)
//			fmt.Println(err)
//			return
//		}
//		json.NewEncoder(w).Encode(ResData)
//	}
}
