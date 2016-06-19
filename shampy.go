package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var db *sql.DB

var shampoos = map[int]Shampoo{1: {"someBrand"}}

type Shampoo struct {
	Brand string
}

func DBConnection() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatal("$DATABASE_URL environment variable must be set")
	} else {
		fmt.Println("Connected to database: " + dbURL)
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT environment variable must be set")
	} else {
		fmt.Println("Running on port: " + port)
	}

	db = DBConnection()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/shampoo", ShampooIndex)
	router.HandleFunc("/shampoo/{id}", GetShampoo).Methods("GET")
	router.HandleFunc("/shampoo/{id}", PostShampoo).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there and welcome to Shampy!")
}

func ShampooIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(shampoos)
	w.Write(j)
}

func GetShampoo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shampooId, _ := strconv.ParseInt(vars["id"], 10, 0)

	var brand string

	err := db.QueryRow("SELECT brand FROM shampoo WHERE id = $1", shampooId).Scan(&brand)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print("Brand: ", brand)
	fmt.Fprintln(w, "Shampoo show:", brand)
}

func PostShampoo(w http.ResponseWriter, r *http.Request) {

}
