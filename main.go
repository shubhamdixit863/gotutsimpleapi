package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name string
	City string
	Age  int
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	// I will check whether the http method is post or not
	if r.Method == http.MethodPost {

		// We will read the post data
		// ioutil
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//
			fmt.Fprintf(w, err.Error())

		}

		// we can unmarshall it

		user := &User{}

		err = json.Unmarshal(body, user)
		if err != nil {
			fmt.Fprintf(w, err.Error())

		}

		fmt.Println(user)
		fmt.Fprintf(w, "success")

	} else {
		fmt.Fprintf(w, "Only Post Method is supported")
	}

}

func main() {

	// We have to attach a listerner

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")

	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Data")
		// all the business process goes here --->
		// and in the end you get a respone back

	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Data")

	})

	http.HandleFunc("/product", ProductHandler)

	// We will create an http server here
	fmt.Println("Server Running at Port 8080")
	http.ListenAndServe(":8080", nil)

}
