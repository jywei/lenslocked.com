package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
	fmt.Fprint(w, "To get in touch, please send and email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":5000", nil)
}
