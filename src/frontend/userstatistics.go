package frontend

import (
	"fmt"
	"github.com/reality95/cf-predictor/src/api"
	"github.com/reality95/cf-predictor/src/lib"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const prefixUserStatistics string = "/UserStatistics/"

//HandleUserStats ... the routing function that finds the statistics of a user
func HandleUserStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl, err := template.ParseFiles("assets/frontend/userstats.html")
		if err != nil {
			log.Printf("Error while extracting userstats.html, %s\n", err.Error())
			fmt.Fprintf(w, "Internal error\n")
			return
		}
		handle := strings.TrimPrefix(r.URL.Path, prefixUserStatistics)
		if r.URL.Path != prefixUserStatistics {
			submissions, err := api.GetUserStatus(handle, nil, nil)
			if err != nil {
				fmt.Fprintf(w, "Error while extracting user status %s", err.Error())
			} else {
				Languages := lib.SelectLanguages(submissions, true)
				tmpl.Execute(w, userStatistics{
					UserStats: false,
					Langs:     Languages,
				})
			}
		} else {
			tmpl.Execute(w, userStatistics{
				UserStats: true,
				Langs:     nil,
			})
		}
	} else {
		handle := r.FormValue("handle")
		http.Redirect(w, r, prefixUserStatistics+handle, http.StatusFound)
	}
}

type userStatistics struct {
	UserStats bool
	Langs     []lib.Lang
}
