package main

import (
	"net/http"
)

func checkAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is running!"))
}
