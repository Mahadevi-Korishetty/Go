package controller

import (
	"encoding/json"
	"net/http"
	"views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Writes to the console
		//fmt.Println("request recieved")
		// Reponds with hello world for a "/" endpoint
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
