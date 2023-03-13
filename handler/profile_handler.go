package handler

import (
	"io"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
