package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Hero struct {
	ID int
	Name string
	Universe string
	Skill string
	ImageURL string
}

type Villain struct {
	ID int
	Name string
	Universe string
	ImageURL string
}

func scanHeroes(db *sql.DB) []Hero {
	var heroes []Hero

	rows, err := db.Query("SELECT id, name, universe, skill, image_url FROM hero")
	if err!= nil {
        fmt.Println(err.Error())
        return nil
    }
	defer rows.Close()

	for rows.Next() {
		var hero Hero
        err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
        if err!= nil {
            fmt.Println(err.Error())
            return nil
        }
        heroes = append(heroes, hero)
	}

	return heroes
}

func scanVillain(db *sql.DB) []Villain {
	var villains []Villain

	rows, err := db.Query("SELECT id, name, universe, image_url FROM villain")
	if err!= nil {
        fmt.Println(err.Error())
        return nil
    }
	defer rows.Close()

	for rows.Next() {
		var villain Villain
        err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
        if err!= nil {
            fmt.Println(err.Error())
            return nil
        }
        villains = append(villains, villain)
	}

	return villains
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME")))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	
	mux := http.NewServeMux()
	mux.HandleFunc("/hero", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(scanHeroes(db))
	})
	mux.HandleFunc("/villain", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(scanVillain(db))
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	fmt.Println("Server is running on port 8080")

	err = server.ListenAndServe()
	if err!= nil {
        panic(err)
    }
}