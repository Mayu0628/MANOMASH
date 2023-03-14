package handler

import (
	"MANOMASH/database"
	"MANOMASH/model"

	"net/http"
	"fmt"
	"encoding/json"
	"reflect"
)

type UpdataUser struct{
	UserName string
	Email string
	Password string
	Introduce string
}

func MyPageEditHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	var checkData UpdataUser
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	checkData.UserName, checkData.Email, checkData.Password, checkData.Introduce = reqUserData.UserName, reqUserData.Email, reqUserData.Password, reqUserData.Introduce
	errMsg := Checkvariable(checkData)//リクエストに空のデータが入ってないか確認する
	if errMsg != ""{
		fmt.Println("エラーです")
		ResData := ResFlgCreate(0,"fail")
		json.NewEncoder(w).Encode(ResData)
		return
	}
	cookies := req.Cookies()
	if cookies != nil{
    for _, c := range cookies {
		var result UpdataUser
        email := c.Value
		database.DB.Raw("SELECT user_name, email, password, introduce FROM users WHERE email = ?", email).Scan(&result)//resultを変数として扱う
		result.UserName, result.Email, result.Password, result.Introduce = reqUserData.UserName, reqUserData.Email, reqUserData.Password, reqUserData.Introduce
		user := model.User {
			UserName : result.UserName,
			Email : result.Email,
			Password : result.Password,
			Introduce : result.Introduce,
		}
		ResData:= ResFlgCreate(1,"succesful")
		database.DB.Model(&user).Where("email = ?", email).Updates(user)
		json.NewEncoder(w).Encode(ResData)
		json.NewEncoder(w).Encode(result)
		return
    }
	}
	fmt.Print("ユーザー情報がありません")
	ResData:= ResFlgCreate(0,"fail")
	json.NewEncoder(w).Encode(ResData)
}

func Checkvariable(UsDt UpdataUser) (err string){
	fmt.Println(UsDt)
	VarNum := reflect.ValueOf(UsDt)
	VarType := reflect.TypeOf(UsDt)
	for i := 0; i < 4; i++ {
		Field := VarType.Field(i)
		VarInfo := VarNum.FieldByName(Field.Name).Interface()
        if VarInfo == "" {
			fmt.Printf("%s is null \n", Field.Name)
			err := "no value"
			return err
        }
	}
	return ""
}
