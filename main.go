package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ungraded-challenge-2/config"
	"ungraded-challenge-2/handler"
)

func main() {
	db, err := config.GetDatabase()
	if err!= nil {
        fmt.Println(err.Error())
        log.Fatal(err)
    }
	defer db.Close()
	
	mux := http.NewServeMux()
	mux.HandleFunc("/hero", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(handler.GetHeroList(db))
	})
	mux.HandleFunc("/villain", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(handler.GetVillainList(db))
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	fmt.Println("Server is running on port 8080")

	err = server.ListenAndServe()
	if err!= nil {
        log.Fatal(err)
    }
}