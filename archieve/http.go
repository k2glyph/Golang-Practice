package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa go is neat")
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8000", nil)
}
