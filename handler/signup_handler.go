package handler

import (
	"io"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
