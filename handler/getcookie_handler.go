package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func GetCookieHandler(w http.ResponseWriter, req *http.Request) {

	cookies := req.Cookies()
	for _, c := range cookies {
		json.NewEncoder(w).Encode(c.Name)
		fmt.Println(c.Name)
	}
	return
}