package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"
	"fmt"
	"time"

	"encoding/json"
	"net/http"
)

// type UpdataUser struct {
// 	UserName  string
// 	Email     string
// 	Password  string
// 	Introduce string
// }

func MyPageEditHandler(w http.ResponseWriter, req *http.Request) {
	// id, err := strconv.Atoi(req.URL.Query().Get("id"))
	// if err != nil {
	// 	ResData := ResFlgCreate(0, "ID変換失敗", 0)
	// 	json.NewEncoder(w).Encode(ResData)
	// 	return
	// }
	var reqUserData model.User
	//var checkData UpdataUser
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		ResData := ResFlgCreate(0, "デコードに失敗しました", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	//checkData.UserName, checkData.Email, checkData.Password, checkData.Introduce = reqUserData.UserName, reqUserData.Email, reqUserData.Password, reqUserData.Introduce
	// errMsg := Checkvariable(checkData)//リクエストに空のデータが入ってないか確認する
	// if errMsg != ""{
	// 	ResData := ResFlgCreate(0,"fail",0)
	// 	json.NewEncoder(w).Encode(ResData)
	// 	return
	// }
	var SendID model.User
	database.DB.First(&SendID, "user_id = ?", reqUserData.User_ID)
	user := &model.User{
		UserName:  reqUserData.UserName,
		Email:     reqUserData.Email,
		Password:  reqUserData.Password,
		Introduce: reqUserData.Introduce,
	}
	fmt.Println(reqUserData)
	result := database.DB.Model(user).Where("user_id = ?", reqUserData.User_ID).Updates(map[string]interface{}{
		"user_name":  user.UserName,
		"email":      user.Email,
		"password":   user.Password,
		"introduce":  user.Introduce,
		"updated_at": time.Now(),
	})
	if result.Error != nil {
		ResData := ResFlgCreate(0, "データを更新できませんでした", 0)

		json.NewEncoder(w).Encode(ResData)
		return
	}

	ResData := ResFlgCreate(1, "succesful", reqUserData.User_ID)
	json.NewEncoder(w).Encode(ResData)
	return
}

//unc Checkvariable(UsDt UpdataUser) (err string){
//	fmt.Println(UsDt)
//	VarNum := reflect.ValueOf(UsDt)
//	VarType := reflect.TypeOf(UsDt)
//	for i := 0; i < 4; i++ {
//		Field := VarType.Field(i)
//		VarInfo := VarNum.FieldByName(Field.Name).Interface()
//       if VarInfo == "" {
//			fmt.Printf("%s is null \n", Field.Name)
//			err := "no value"
//			return err
//       }
//	}
//	return ""
//
