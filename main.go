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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	database.GormConnect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	fmt.Println("hallo,world")
	r.HandleFunc("/login", handler.LoginHandler).Methods(http.MethodGet)
	r.HandleFunc("/sign-up", handler.SignUpHandler).Methods(http.MethodGet)
	r.HandleFunc("/mypage/edit",handler.MyPageEditHandler).Methods(http.MethodGet)
	r.HandleFunc("/mypage",handler.MyPageHandler).Methods(http.MethodPost)
	r.HandleFunc("/profile", handler.ProfileHandler).Methods(http.MethodPost)
	r.HandleFunc("/profile/add", handler.ProfileAddHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile/edit", handler.ProfileEditHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile/delete", handler.ProfileDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("toppage", handler.TopPageHandler).Methods(http.MethodPost)

	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}
