package handler

import (
	"io"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}

