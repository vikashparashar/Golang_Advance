package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Greet)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome To GIC - Golang Basic Training Program  ")
}
