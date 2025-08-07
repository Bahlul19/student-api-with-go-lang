package main

import (
	"fmt"
	"go_project/internal/config"
	"log"
	"net/http"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database connection

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Student API!"))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server is running %s", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server")
	}
	// fmt.Println("Server is running on")
}
