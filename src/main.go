package main

import (
	"github.com/reality95/cf-predictor/src/frontend"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/userstats/", frontend.HandleUserStats)
	http.HandleFunc("/userscompare/", frontend.HandleUsersCompare)
	http.ListenAndServe(":8080", nil)
	log.Println("Hello word")
}
