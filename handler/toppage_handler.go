package handler

import (
	"io"
	"net/http"
)

func TopPageHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
 