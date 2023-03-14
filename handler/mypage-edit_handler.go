package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"

	"encoding/json"
	"net/http"
	"strconv"
)

type UpdataUser struct {
	UserName  string
	Email     string
	Password  string
	Introduce string
}

func MyPageEditHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	var reqUserData model.User
	var checkData UpdataUser
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		ResData := ResFlgCreate(0, "fail", 0)
		json.NewEncoder(w).Encode(ResData)
		return
	}
	checkData.UserName, checkData.Email, checkData.Password, checkData.Introduce = reqUserData.UserName, reqUserData.Email, reqUserData.Password, reqUserData.Introduce
	// errMsg := Checkvariable(checkData)//リクエストに空のデータが入ってないか確認する
	// if errMsg != ""{
	// 	ResData := ResFlgCreate(0,"fail",0)
	// 	json.NewEncoder(w).Encode(ResData)
	// 	return
	// }
	var SendID model.User
	database.DB.First(&SendID, "user_id = ?", id)
	user := &model.User{
		UserName:  reqUserData.UserName,
		Email:     reqUserData.Email,
		Password:  reqUserData.Password,
		Introduce: reqUserData.Introduce,
	}
	database.DB.Model(user).Updates(map[string]interface{}{
		"name":  "user3",
		"email": "g5.taisa831@gmail.com",
	})

	ResData := ResFlgCreate(1, "succesful", SendID.Id)
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
