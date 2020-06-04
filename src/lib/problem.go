package lib

import (
	"github.com/reality95/cf-predictor/src/api"
	"sort"
)

var LANGUAGES = map[string]string{
	"GNU C++0x":             "C/C++",
	"GNU C11":               "C/C++",
	"GNU C++11":             "C/C++",
	"GNU C++14":             "C/C++",
	"GNU C++17":             "C/C++",
	"GNU C++":               "C/C++",
	"Clang++17 Diagnostics": "C/C++",
	"MS C++":                "C/C++",
	"MS C++ 2017":           "C/C++",
	"Go":                    "Go",
	"Java 8":                "Java",
	"Java 11":               "Java",
	"Python 2":              "Python",
	"Python 3":              "Python",
	"Pypy 2":                "Python",
	"PyPy 3":                "Python",
	"PHP":                   "PHP",
	"FPC":                   "Pascal",
	"Delphi":                "Delphi",
	"JavaScript":            "JavaScript",
	"Node.js":               "Node.js",
	"Scala":                 "Scala",
	"Rust":                  "Rust",
	"Ruby":                  "Ruby",
	"Perl":                  "Perl",
	"OCalm":                 "OCalm",
	"Haskell":               "Haskell",
	"PascalABC.NET":         "Pascal",
	"Mono C#":               "C#",
	"D":                     "D",
}

func SelectProblems(submissions []api.Submission, OKOnly bool) (ans []api.Problem) {
	problems := make([]api.Problem, 0)
	ans = make([]api.Problem, 0)
	for _, s := range submissions {
		if !OKOnly || s.Verdict == "OK" {
			problems = append(problems, s.Problemv)
		}
	}
	SortProblems(problems)
	for i, p := range problems {
		if i+1 == len(problems) || !api.EqProblem(p, problems[i+1]) {
			ans = append(ans, p)
		}
	}

	return ans
}

func SelectLanguages(submissions []api.Submission, OKOnly bool) (ans []Lang) {
	count := make(map[string]int)
	for _, s := range submissions {
		if !OKOnly || s.Verdict == "OK" {
			if pLang, ok := LANGUAGES[s.ProgrammingLanguage]; ok {
				count[pLang] += 1
			}
		}
	}
	for name, cnt := range count {
		ans = append(ans,Lang {
			Name : name,
			Count : cnt,
		})
	}
	sort.SliceStable(ans,func (i,j int) bool {
		return ans[i].Count > ans[j].Count
	})
	return ans
}

func ProblemSetIntersection(a, b []api.Problem) (both, aOnly, bOnly []api.Problem) {
	both = make([]api.Problem, 0)
	aOnly = make([]api.Problem, 0)
	bOnly = make([]api.Problem, 0)
	SortProblems(a)
	SortProblems(b)
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if api.EqProblem(a[i], b[j]) {
			both = append(both, a[i])
			i++
			j++
		} else if api.LessProblem(a[i], b[j]) {
			aOnly = append(aOnly, a[i])
			i++
		} else {
			bOnly = append(bOnly, b[j])
			j++
		}
	}
	for i < len(a) {
		aOnly = append(aOnly, a[i])
		i++
	}
	for j < len(b) {
		bOnly = append(bOnly, b[j])
		j++
	}

	return both, aOnly, bOnly
}

func SortProblems(problems []api.Problem) {
	sort.SliceStable(problems, func(i, j int) bool {
		return api.LessProblem(problems[i], problems[j])
	})
}

type Lang struct {
	Name string
	Count int
}

