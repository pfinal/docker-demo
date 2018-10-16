package http

import (
	"net/http"
	"log"
	"os"
)

func configRoutes() {
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("1.0.0"))
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("order ok"))
	})
}
func init() {
	configRoutes()
}

func Start() {

    addr := os.Getenv("addr")
    if addr  == "" {
        addr = ":8080"
    }
    
	s := &http.Server{
		Addr:           addr,
	}

	log.Println("listening", addr)
	log.Fatalln(s.ListenAndServe())
}
