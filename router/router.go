// package router

// import(
// 	"net/http"

// 	"MANOMASH/controller"

// 	"github.com/gorilla/mux"
// )

// func NewRouter(con *controller.MyAppController) *mux.Router {
// 	r := NewRouter(con)

// 	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)

// 	r.HandleFunc("/sign-up", con.SignUpHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/login", con.LogInHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/mypage", con.MyPageHandler).Methods(http.MethodPost)
// 	r.HandleFunc("/mypage/add", con.MyPageEditHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/mypage/edit", con.MyPageEditHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/mypage/delete", con.MyPageEditHandler).Methods(http.MethodDelete)
// 	r.HandleFunc("/toppage", con.TopPageHandler).Methods(http.MethodPost)

// 	//r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

// 	return r
// }