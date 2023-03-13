package handler

import (
	"io"
	"net/http"
)

func ProfileDeleteHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
