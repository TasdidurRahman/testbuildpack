package main

import (
	"fmt"
	"net/http"
)

func user(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "USER Updated\n")
}

func main() {

	http.HandleFunc("/user", user)
	http.ListenAndServe(":8080", nil)
}
