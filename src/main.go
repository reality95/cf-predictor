package main

import (
	"net/http"
	"github.com/reality95/cf-predictor/src/frontend"
)
func main() {
	http.HandleFunc("/userstats/", frontend.HandleUserStats)
	http.HandleFunc("/userscompare/", frontend.HandleUsersCompare)
	http.ListenAndServe(":8080",nil)

}