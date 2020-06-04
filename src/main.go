package main

import (
	"github.com/reality95/cf-predictor/src/frontend"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/UserStatistics/", frontend.HandleUserStats)
	http.HandleFunc("/CompareUsers/", frontend.HandleCompareUsers)
	http.HandleFunc("/", frontend.HandleIndex)
	http.HandleFunc("/Home/", frontend.HandleIndex)
	http.ListenAndServe(":8080", nil)
	log.Println("Hello word")
}
