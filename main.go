package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/postgres", func(w http.ResponseWriter, r *http.Request) {
		postgres(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

func secret(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("SECRET")
	io.WriteString(w, fmt.Sprintf("%s\n", secret))
}

func postgres(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("value")
	db.Exec("INSERT INTO test (value) VALUES ($1)", value)
	rows, err := db.Query("SELECT value FROM test")
	if err != nil {
		panic(err)
	}
	values := []string{}
	for rows.Next() {
		value := ""
		if err := rows.Scan(&value); err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	response, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, fmt.Sprintf("%s\n", response))
}
