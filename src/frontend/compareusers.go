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

const prefixCompareUsers string = "/CompareUsers/"

func HandleCompareUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl, err := template.ParseFiles("assets/frontend/userscompare.html")
		if err != nil {
			log.Printf("Error while extracting compareUsers.html, %s\n", err.Error())
			fmt.Fprintf(w, "Internal error\n")
			return
		}

		handles := strings.Split(strings.TrimPrefix(r.URL.Path, prefixCompareUsers), "&")
		if r.URL.Path != prefixCompareUsers {
			if len(handles) != 2 {
				fmt.Fprintf(w, "Expected only 2 handles, got %d\n", len(handles))
				return
			}

			submissions1, err := api.GetUserStatus(handles[0], nil, nil)
			if err != nil {
				fmt.Fprintf(w, "Error while extracting info from handle %s\n", handles[0])
				return
			}

			submissions2, err := api.GetUserStatus(handles[1], nil, nil)
			if err != nil {
				fmt.Fprintf(w, "Error while extracting info from handle %s\n", handles[1])
				return
			}

			both, aOnly, bOnly := lib.ProblemSetIntersection(lib.SelectProblems(submissions1, true), lib.SelectProblems(submissions2, true))

			Stats, _, err := api.GetPsetProblems(nil, nil)
			if err != nil {
				fmt.Fprintf(w, "Error while extracting the problems")
				log.Println(err.Error())
				return
			}

			solveCount := make(map[string]int)
			for _, s := range Stats {
				solveCount[s.Hash()] = s.SolvedCount
			}

			var data compareUsers
			data.User1 = handles[0]
			data.User2 = handles[1]
			data.CompareUsers = false

			for _, p := range both {
				if cnt, ok := solveCount[p.Hash()]; ok {
					data.CommonProblems = append(data.CommonProblems, problemStats{
						Name:        p.Name,
						Link:        p.Link(),
						SolvedCount: cnt,
					})
				}
			}

			for _, p := range aOnly {
				if cnt, ok := solveCount[p.Hash()]; ok {
					data.Problems1 = append(data.Problems1, problemStats{
						Name:        p.Name,
						Link:        p.Link(),
						SolvedCount: cnt,
					})
				}
			}

			for _, p := range bOnly {
				if cnt, ok := solveCount[p.Hash()]; ok {
					data.Problems2 = append(data.Problems2, problemStats{
						Name:        p.Name,
						Link:        p.Link(),
						SolvedCount: cnt,
					})
				}
			}

			sort.SliceStable(data.CommonProblems, func(i, j int) bool {
				return data.CommonProblems[i].SolvedCount < data.CommonProblems[j].SolvedCount
			})
			sort.SliceStable(data.Problems1, func(i, j int) bool {
				return data.Problems1[i].SolvedCount < data.Problems1[j].SolvedCount
			})
			sort.SliceStable(data.Problems2, func(i, j int) bool {
				return data.Problems2[i].SolvedCount < data.Problems2[j].SolvedCount
			})

			tmpl.Execute(w, data)
		} else {

			tmpl.Execute(w, compareUsers{
				CompareUsers: true,
			})
		}
	} else {
		handle1 := r.FormValue("handle1")
		handle2 := r.FormValue("handle2")
		http.Redirect(w, r, prefixCompareUsers+handle1+"&"+handle2, http.StatusFound)
	}
}

type compareUsers struct {
	CompareUsers   bool
	User1          string
	User2          string
	CommonProblems []problemStats
	Problems1      []problemStats
	Problems2      []problemStats
}

type problemStats struct {
	Name        string
	Link        string
	SolvedCount int
}
