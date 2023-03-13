package handler

import (
	"io"
	"net/http"
)

func MyPageEditHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
