package frontend

import (
	"fmt"
	"github.com/reality95/cf-predictor/src/api"
	"github.com/reality95/cf-predictor/src/lib"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

const prefixUserStatistics string = "/UserStatistics/"

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
				count := lib.SelectLanguages(submissions, true)
				lgs := make([]lang, 0)
				for name, cnt := range count {
					lgs = append(lgs, lang{
						Name:  name,
						Count: cnt,
					})
				}
				sort.SliceStable(lgs, func(i, j int) bool {
					return lgs[i].Count > lgs[j].Count
				})
				tmpl.Execute(w, userStatistics{
					UserStats: false,
					Langs:     lgs,
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

type lang struct {
	Name  string
	Count int
}

type userStatistics struct {
	UserStats bool
	Langs     []lang
}
