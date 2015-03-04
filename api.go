package main

import (
	"net/http"
)

func blah(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.UserAgent()))
}
func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(blah))
}
