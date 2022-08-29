package main

import (
	"fmt"
	"net/http"
)

func main() {
	Rooter()
}

// Rooter can run server
func Rooter() {
	http.HandleFunc("/", rootHandel)
	println("start server port :8000...")
	http.ListenAndServe(":8000", nil)
}
func rootHandel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcom To My Website</h1>")
}
