package handler

import (
	"io"
	"net/http"
)

func MyPageHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
