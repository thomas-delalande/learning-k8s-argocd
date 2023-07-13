package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/secret", secret)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

func secret(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("SECRET")
	io.WriteString(w, fmt.Sprintf("%s\n", secret))
}
