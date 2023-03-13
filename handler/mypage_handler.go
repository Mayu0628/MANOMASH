package handler

import (
	"MANOMASH/database"

	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

func MyPageHandler(w http.ResponseWriter, req *http.Request) {
	cookies := req.Cookies()
	type User struct{
		Id int
		UserName string
		Email string
		Password string
		Introduce string
	}
	if cookies != nil{
    for _, c := range cookies {
		var result User
        email := c.Value
		log.Print("Name:", c.Name, "Value:", c.Value)
		database.DB.Raw("SELECT id, user_name, email, password, introduce FROM users WHERE email = ?", email).Scan(&result) // (*sql.Row)
		fmt.Println(result)
		fmt.Println(result.Email)
		ResData:= ResFlgCreate(1,"succesful")
		json.NewEncoder(w).Encode(ResData)
		json.NewEncoder(w).Encode(result)
		return
    }
	}
	fmt.Print("クッキーがないです")
	ResData:= ResFlgCreate(0,"fail")
	json.NewEncoder(w).Encode(ResData)
}
