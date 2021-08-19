package main

import (
	"controller"
	"net/http"
)

func main() {
	// mux: Multiplexer
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux) //Cretaed a server listening on port 3000 and passed multiplexer or router,
	//router defines an endpoint. The defination contains a response writer and a request reader
}
