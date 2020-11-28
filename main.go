package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nouzun/l2r/db"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	db.ConnectDatabase()

	handleRequests()
}
