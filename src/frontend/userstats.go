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

func HandleUserStats(w http.ResponseWriter,r *http.Request) {
	tmpl, err := template.ParseFiles("assets/frontend/userstats.html")
	if err != nil {
		log.Printf("Erorr while extracting userstats.html, %s\n",err.Error())
		fmt.Fprintf(w,"Internal error\n")
		return
	}
	handle := strings.TrimPrefix(r.URL.Path,"/userstats/")
	submissions, err := api.GetUserStatus(handle,nil,nil)
	if err != nil {
		fmt.Fprintf(w,"Erorr while extracting user status %s",err.Error())
	} else {
		count := make(map[string]int)
		for _, s := range(submissions) {
			if s.Verdict == "OK" {
				switch s.ProgrammingLanguage {
					case "GNU C++0x":
						count["C/C++"] += 1
					case "GNU C11":
						count["C/C++"] += 1
					case "GNU C++11":
						count["C/C++"] += 1
					case "GNU C++14":
						count["C/C++"] += 1
					case "GNU C++17":
						count["C/C++"] += 1
					case "GNU C++":
						count["C/C++"] += 1
					case "Clang++17 Diagnostics":
						count["C/C++"] += 1
					case "MS C++":
						count["C/C++"] += 1
					case "MS C++ 2017":
						count["C/C++"] += 1
					case "Go":
						count["Go"] += 1
					case "Java 8":
						count["Java"] += 1
					case "Java 11":
						count["Java"] += 1
					case "Python 2":
						count["Python/Pypy"] += 1
					case "Python 3":
						count["Python/Pypy"] += 1
					case "Pypy 2":
						count["Python/Pypy"] += 1
					case "PyPy 3":
						count["Python/Pypy"] += 1
					case "PHP":
						count["PHP"] += 1
					case "FPC":
						count["Pascal"] += 1
					case "Delphi":
						count["Delphi"] += 1
					case "JavaScript":
						count["JavaScript"] += 1
					case "Node.js":
						count["JavaScript"] += 1
					case "Scala":
						count["Scala"] += 1
					case "Rust":
						count["Rust"] += 1
					case "Ruby":
						count["Ruby"] += 1
					case "Perl":
						count["Perl"] += 1
					case "OCalm":
						count["OCalm"] += 1
					case "Haskell":
						count["Haskell"] += 1
					case "PascalABC.NET":
						count["Pascal"] += 1
					case "Mono C#":
						count["C#"] += 1
					case "D":
						count["D"] += 1
				}
			}
		}
		var lgs langs
		lgs.Title = "The languages used by user " + handle + " are:"
		for name, cnt := range(count) {
			lgs.Langs = append(lgs.Langs,lang{
				Name : name,
				Count : cnt,
			})
		}
		sort.SliceStable(lgs.Langs,func(i,j int) bool {
			return lgs.Langs[i].Count > lgs.Langs[j].Count
		})
		log.Println(lgs)
		tmpl.Execute(w,lgs)
	}
}

type lang struct {
	Name string
	Count int
}

type langs struct {
	Title string
	Langs []lang
}

