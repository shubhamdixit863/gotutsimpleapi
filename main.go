package main

import (
	"fmt"
	"gotutsapi/handlers"
	"net/http"
)

func main() {

	// We have to attach a listerner

	http.HandleFunc("/", handlers.RootHandler)

	http.HandleFunc("/user", handlers.UserHandler)

	// We will create an http server here
	fmt.Println("Server Running at Port 8080")
	http.ListenAndServe(":8080", nil)

}
