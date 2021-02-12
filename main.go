package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", hello)

	http.ListenAndServe(":9090", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>My name is ...</h1>")
}
