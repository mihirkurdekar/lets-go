package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Hello World")

	router := chi.NewRouter()

	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed")
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
