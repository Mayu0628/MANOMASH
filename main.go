package main

import (
	"MANOMASH/database"
	//"MANOMASH/controller"
	//"MANOMASH/router"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	db := database.GormConnect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	err = sqlDB.Ping()

	if err != nil {
		fmt.Println("エラー")
		fmt.Println(err)
		return
	} else {
		fmt.Println("データベース接続成功")
	}

	
	fmt.Println("hallo,world")
	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}

//func main() {
//	//c := cors.New(cors.Options{
//	//	AllowedOrigins:   []string{"http://localhost:8080"},
//	//	AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
//	//	AllowedHeaders:   []string{"*"},
//	//	AllowCredentials: true,
//	//})
//	database.GormConnect()
//	sqlDB := database.DB
//	db := sqlDB.DB()
//	defer db.Close()
//
//	err := db.Ping()
//
//	if err != nil {
//		fmt.Println("エラー")
//		fmt.Println(err)
//		return
//	} else {
//		fmt.Println("データベース接続成功")
//	}
//
//	fmt.Println("hallo,world")
//	//con := controller.NewMyAppController(sqlDB)
//
//	//r := router.NewRouter(con)
//	//http.ListenAndServe(":8080", r)
//}
//