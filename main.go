package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	port    = 8080
	timeout = 15 * time.Second
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир!"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	srv.ListenAndServe()
}
