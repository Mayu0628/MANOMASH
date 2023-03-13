package handler

import (
	"io"
	"net/http"
)

func ProfileAddHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
