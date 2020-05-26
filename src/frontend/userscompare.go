package frontend

import (
	"fmt"
	"github.com/reality95/cf-predictor/src/api"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func HandleUsersCompare(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/frontend/userscompare.html")
	if err != nil {
		log.Printf("Error while extracting userscompare.html, %s\n", err.Error())
		fmt.Fprintf(w, "Internal error\n")
		return
	}

	handles := strings.Split(strings.TrimPrefix(r.URL.Path, "/userscompare/"), "&")
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

	p1 := make(map[string]bool)
	p2 := make(map[string]bool)

	for _, s := range submissions1 {
		if s.Verdict == "OK" {
			p1[strconv.Itoa(s.Problemv.ContestID)+"$"+s.Problemv.Index+"#"+s.Problemv.Name] = true
		}
	}
	for _, s := range submissions2 {
		if s.Verdict == "OK" {
			p2[strconv.Itoa(s.Problemv.ContestID)+"$"+s.Problemv.Index+"#"+s.Problemv.Name] = true
		}
	}

	Stats, _, err := api.GetPsetProblems(nil, nil)
	if err != nil {
		fmt.Fprintf(w, "Error while extracting the problems")
		log.Println(err.Error())
		return
	}

	solveCount := make(map[string]int)
	for _, s := range Stats {
		solveCount[strconv.Itoa(s.ContestID)+"$"+s.Index] = s.SolvedCount
	}

	var data usersCompare
	data.User1 = handles[0]
	data.User2 = handles[1]

	for prob, c := range p1 {
		d := strings.Split(prob, "#")

		if sCount, ok := solveCount[d[0]]; ok {
			tmp := strings.Split(d[0],"$")
			contestID, Index := tmp[0], tmp[1]
			if p2[prob] && c {
				data.CommonProblems = append(data.CommonProblems, problemStats{
					Name:        d[1],
					SolvedCount: sCount,
					Link : getProblemLink(contestID,Index),
				})
			} else if c {
				data.Problems1 = append(data.Problems1, problemStats{
					Name:        d[1],
					SolvedCount: sCount,
					Link : getProblemLink(contestID,Index),
				})
			}
		}
	}

	for prob, c := range p2 {
		d := strings.Split(prob, "#")

		if sCount, ok := solveCount[d[0]]; ok {
			tmp := strings.Split(d[0],"$")
			contestID, Index := tmp[0], tmp[1]
			if !p1[prob] && c {
				data.Problems2 = append(data.Problems2, problemStats{
					Name:        d[1],
					SolvedCount: sCount,
					Link : getProblemLink(contestID,Index),
				})
			}
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
}

func getProblemLink(contestID,Index string) string {
	return "https://codeforces.com/contest/" + contestID + "/problem/" + Index
}

type usersCompare struct {
	User1          string
	User2          string
	CommonProblems []problemStats
	Problems1      []problemStats
	Problems2      []problemStats
}

type problemStats struct {
	Name        string
	Link 		string
	SolvedCount int
}
