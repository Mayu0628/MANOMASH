package main

import (
	"MANOMASH/database"
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
	database.GormConnect()
	sqlDB, err := database.DB.DB()
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
