package main

import (
	"MANOMASH/database"
	"MANOMASH/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	c := cors.AllowAll()
	database.GormConnect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	fmt.Println("hallo,world")
	r.HandleFunc("/login", handler.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", handler.SignUpHandler).Methods(http.MethodPost)
	r.HandleFunc("/mypage/edit", handler.MyPageEditHandler).Methods(http.MethodPost)
	r.HandleFunc("/mypage", handler.MyPageHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile", handler.ProfileHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile/add", handler.ProfileAddHandler).Methods(http.MethodPost)
	r.HandleFunc("/profile/edit", handler.ProfileEditHandler).Methods(http.MethodPost)
	r.HandleFunc("/profile/delete", handler.ProfileDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/toppage", handler.TopPageHandler).Methods(http.MethodGet)

	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}
