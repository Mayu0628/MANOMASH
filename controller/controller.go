//package controller
//
//import (
//	"gorm.io/gorm"
//	"io"
//	"net/http"
//)
//
//type MyAppController struct{
//	db *gorm.DB
//}
//
//func NewMyAppController(db *gorm.DB) *MyAppController{
//	return &MyAppController{db: db}
//}
//
//func (db *MyAppController) HelloHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}
//
//func (db *MyAppController) SignUpHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}
//
//func (db *MyAppController) LogInHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}
//
//func (db *MyAppController) MyPageHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}
//
//func (db *MyAppController) MyPageAddHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}
//
//func (db *MyAppController) MyPageEditHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}
//
//func (db *MyAppController) MyPageDeleteHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}
//
//func (db *MyAppController) TopPageHandler(w http.ResponseWriter, req *http.Request) {
//	io.WriteString(w, "Hello, world!\n")
//}