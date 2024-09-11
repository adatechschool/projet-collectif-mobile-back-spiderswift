package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Building Rest APIs in Go 1.23.1!!")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request){

		fmt.Fprint(w, "return all comments")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id")
		fmt.Fprint(w, "return a single comment for comment with id: %s", id)
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request){

		fmt.Fprint(w, "post a new comment")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())

	}
	
}