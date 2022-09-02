package router

import (
	"fmt"
	"log"

	"net/http"
)

type root []string

// Rooter can run server
func Rooter(addr string)error {
	http.HandleFunc("/", rootHandel)
	
	println("start server port :8000...")
	err:=http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println()
		log.Fatalf("error : ",err)

	}
return err
}
func rootHandel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcom To My Website %s</h1>" ,r.RemoteAddr)
	fmt.Println("run home page")
}
