package handlers

import (
	"encoding/json"
	"fmt"
	"gotutsapi/db"
	"gotutsapi/dto"
	"io/ioutil"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

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

		user := &dto.User{}

		err = json.Unmarshal(body, user)
		if err != nil {
			fmt.Fprintf(w, err.Error())

		}

		// We will push it in our db as well

		db.Db = append(db.Db, *user)

		// We will create a json response by marshalling the struct

		res := dto.Response{
			Data:    "Some data",
			Message: "Success",
		}

		// Lets marshall it
		resData, err := json.Marshal(res)
		if err != nil {
			fmt.Fprintf(w, err.Error())

		}

		//fmt.Println(user)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(resData))

	} else {
		// We cand handle the get request
		w.Header().Set("Content-Type", "application/json")
		res := dto.Response{
			Data:    db.Db,
			Message: "Success",
		}

		resData, err := json.Marshal(res)
		if err != nil {
			fmt.Fprintf(w, err.Error())

		}

		//fmt.Println(user)

		fmt.Fprintf(w, string(resData))

	}

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")

}
