package handler

import (
	"io"
	"net/http"
)

func ProfileEditHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
