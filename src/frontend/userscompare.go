package frontend

import (
	"net/http"
	"html/template"
	"strings"
	"github.com/reality95/cf-predictor/src/api"
	"fmt"
	"log"
	"sort"
)

func HandleUsersCompare(w http.ResponseWriter,r *http.Request) {
	tmpl, err := template.ParseFiles("assets/frontend/userscompare.html")
	if err != nil {
		log.Printf("Error while extracting userscompare.html, %s\n",err.Error())
		fmt.Fprintf(w,"Internal error\n")
		return
	}
	
	handles := strings.Split(strings.TrimPrefix(r.URL.Path,"/userscompare/"),"&")
	if len(handles) != 2 {
		fmt.Fprintf(w,"Expected only 2 handles, got %d\n",len(handles))
		return
	}

	submissions1, err := api.GetUserStatus(handles[0],nil,nil)
	if err != nil {
		fmt.Fprintf(w,"Error while extracting info from handle %s\n",handles[0])
		return
	}
	
	submissions2, err := api.GetUserStatus(handles[1],nil,nil)
	if err != nil {
		fmt.Fprintf(w,"Error while extracting info from handle %s\n",handles[1])
		return
	}

	p1 := make(map[string]bool)
	p2 := make(map[string]bool)
	
	for _, s := submissions1 {
		p1[strconv.Itoa(s.Problemv.ContestId) + "$" + s.Problemv.Index + "#" + s.Problemv.Name] = true
	}
	for _, s := submissions2 {
		p2[strconv.Itoa(s.Problemv.ContestId) + "$" + s.Problemv.Index + "#" + s.Problemv.Name] = true
	}

	Stats, _, err = api.GetPsetProblems(nil,nil)
	if err != nil {
		fmt.Fprintf(w,"Error while extracting the problems")
		log.Println(err.Error())
		return
	}

	solveCount := make(map[string]int)
	for _, s := Stats {
		solveCount[strconv.Itoa(s.ContestId) + "$" + s.Index] = s.SolvedCount
	}

	var data usersCompare
	data.User1 = handles[0]
	data.User2 = handles[1]

	for prob, c := range(p1) {
		d := strings.Split(prob,"#")
		if p2[prob] && c {
			data.CommonProblems = append(data.CommonProblems, problemStats {
				Name : d[1],
				SolvedCount : solveCount[d[0]],
			})
		} else if c {
			data.Problems1 = append(data.Problems1, problemStats {
				Name : d[1],
				SolvedCount : solveCount[d[0]],
			})
		}
	}

	for prob, c := range(p2) {
		d := strings.Split(prob,"-")
		if !p1[prob] && c {
			data.Problems2 = append(data.Problems2,problemStats {
				Name : d[1],
				SolvedCount : solveCount[d[0]],
			})
		}
	}

	sort.SliceStable(data.CommonProblems,func(i,j int) {
		return data.CommonProblems[i].SolvedCount < data.CommonProblems[j].SolvedCount
	})
	sort.SliceStable(data.Problems1,func(i,j int) {
		return data.Problems1[i].SolvedCount < data.Problems1[j].SolvedCount
	})
	sort.SliceStable(data.Problems2,func(i,j int) {
		return data.Problems1[i].SolvedCount < data.Problems2[j].SolvedCount
	})

	tmpl.Execute(w, data)
}

type usersCompare struct {
	User1 string
	User2 string
	CommonProblems []problemStats
	Problems1 []problemStats
	Problems2 []problemStats
}

type problemStats struct {
	Name string
	SolvedCount int
}