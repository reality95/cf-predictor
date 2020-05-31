package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const prefixContestantsInfo string = `<td( class="(gold|bronze|silver)")?(title="Unofficial participation")?>`
const suffixContestantsInfo string = `</td>`

/*
	The `Rank`, `ScoreSum`, `Rel` and `Medal` can be found under the following regexp
	prefixContestantsInfo + `({{.Rank}}|{{.ScoreSum}}|({{.Rel}}%)|{{.Medal}})` + suffixContestantsInfo
	in the following order.
*/

//ExtractYearInfo ... extracts the contestants' info and contest's info from year `year`
/*
	If CFHandleInclude is true then in the contestant's info CFHandle will be included.
	This idea behind this parameter is that the time it takes to extract the CFHandle is 
	significantly bigger than all other information
*/
func ExtractYearInfo(year int,CFHandleInclude bool) (contestants []IOIContestant, info IOIInfo, err error) {
	if year < 1989 {
		return nil, info, errors.New("IOI started at 1989, so year must be at least 1989")
	}

	resp, err := http.Get("https://stats.ioinformatics.org/results/" + strconv.Itoa(year))
	if err != nil {
		return nil, info, err
	}

	plainText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, info, err
	}

	info.Tasks = extractTaskNames(string(plainText), year)

	prefix := prefixContestantsInfo
	suffix := suffixContestantsInfo
	re := regexp.MustCompile(prefix + `(([A-Za-z]*)|([0-9]+.?[0-9]*%?))` + suffix)
	rePrefix := regexp.MustCompile(prefix)
	reSuffix := regexp.MustCompile(suffix)

	tmp := re.FindAllIndex(plainText, -1)

	if len(tmp)%4 != 0 {
		panic("Expected 4 fields with data for every contestant")
	}

	info.ParticipantCount = len(tmp) / 4

	for i := 0; i < info.ParticipantCount; i++ {
		var contestant IOIContestant
		for j := 0; j < 4; j++ {
			l, r := tmp[4*i+j][0], tmp[4*i+j][1]
			sValue := string(reSuffix.ReplaceAll(rePrefix.ReplaceAll([]byte(plainText[l : r]), nil), nil))
			if j == 0 && sValue != "" {
				contestant.Rank, err = strconv.Atoi(sValue)
				if err != nil {
					panic(err.Error())
				}
			} else if j == 1 {
				contestant.ScoreSum, err = strconv.ParseFloat(sValue, 64)
				if err != nil {
					panic(err.Error())
				}
			} else if j == 2 {
				contestant.Rel, err = strconv.ParseFloat(sValue[:len(sValue)-1], 64) //Ignore the last '%'
				if err != nil {
					panic(err.Error())
				}
			} else if j == 3 {
				if sValue == "" {
					contestant.Medal = "No Medal"
				} else {
					contestant.Medal = sValue
				}
				switch sValue {
					case "Gold":info.Gold++
					case "Silver":info.Silver++
					case "Bronze":info.Bronze++
					case "":info.NoMedal++
				}
			}
		}
		//The slice containing the info containing info about the contestant
		plainTextSlice := string(plainText[tmp[4*i][0]:tmp[4*i+3][1]])
		contestant.ID, contestant.Name = extractIDName(plainTextSlice)
		contestant.Score = extractTaskScores(plainTextSlice)
		contestant.Country = extractCountry(plainTextSlice)
		if CFHandleInclude {
			contestant.CFHandle, _ = ExtractCFHandle(contestant.ID)
		}
		contestants = append(contestants, contestant)
	}

	return contestants, info, nil
}

//ExtractContestantInfo ... finds the results of the contestant with id `ID` from all years
func ExtractContestantInfo(ID int) (ans map[int]IOIContestant,err error) {
	ans = make(map[int]IOIContestant)

	resp, err := http.Get("https://stats.ioinformatics.org/people/" + strconv.Itoa(ID))
	if err != nil {
		return ans, err
	}

	plainText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ans, err
	}

	prefix := prefixContestantsInfo
	suffix := suffixContestantsInfo
	re := regexp.MustCompile(prefix + `(([A-Za-z]*)|([0-9]+.?[0-9]*%?)|([0-9]+/[0-9]+))` + suffix)
	rePrefix := regexp.MustCompile(prefix)
	reSuffix := regexp.MustCompile(suffix)

	slices, years := extractYears(string(plainText))

	Name := extractName(string(plainText))
	CFHandle, _ := ExtractCFHandle(ID)


	for i,plainTextSlice := range slices {
		var contestant IOIContestant
		tmp := re.FindAllIndex([]byte(plainTextSlice),-1)
		for j := 0;j < 4;j++ {
			l,r := tmp[j][0],tmp[j][1]
			sValue := string(reSuffix.ReplaceAll(rePrefix.ReplaceAll([]byte(plainTextSlice[l : r]), nil), nil))
			if j == 0 && sValue != "" {
				contestant.Rank, err = strconv.Atoi(sValue)
				if err != nil {
					panic(err.Error())
				}
			} else if j == 1 {
				contestant.ScoreSum, err = strconv.ParseFloat(sValue, 64)
				if err != nil {
					panic(err.Error())
				}
			} else if j == 2 {
				contestant.Rel, err = strconv.ParseFloat(sValue[:len(sValue)-1], 64) //Ignore the last '%'
				if err != nil {
					panic(err.Error())
				}
			} else if j == 3 {
				if sValue == "" {
					contestant.Medal = "No Medal"
				} else {
					contestant.Medal = sValue
				}
			}
		}
		contestant.ID = ID
		contestant.Name = Name
		contestant.CFHandle = CFHandle
		contestant.Score = extractTaskScores(plainTextSlice)
		ans[years[i]] = contestant
	}
	return ans, nil
}

const prefixCFHandle string = `<a href="https?://codeforces.com/profile/`
const suffixCFHandle string = `">`

/*
	The CFHandle can be found under `<a href="https?://codeforces.com/profile/{{.CFHandle}}">`
*/

//ExtractCFHandle ... extracts the Codeforces Handle from contestant with id `ID` using the contestant's page
func ExtractCFHandle(ID int) (string, error) {
	resp, err := http.Get("https://stats.ioinformatics.org/people/" + strconv.Itoa(ID))
	if err != nil {
		return "", err
	}

	plainText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	prefix := prefixCFHandle
	suffix := suffixCFHandle
	re := regexp.MustCompile(prefix + `[^"]+` + suffix)
	rePrefix := regexp.MustCompile(prefix)
	reSuffix := regexp.MustCompile(suffix)

	return string(rePrefix.ReplaceAll(reSuffix.ReplaceAll(re.Find(plainText),nil),nil)), nil
}

const prefixTasks string = `href="tasks/`
const suffixTasks string = `">`

/*
	The `Tasks` can be found only under the format `<a href="tasks/{{.Year}}/{{.TaskName}}">`
	with exception of the first occurence being <a href="tasks/2019/">
*/

func extractTaskNames(plainText string, year int) (ans []string) {
	prefix := prefixTasks + strconv.Itoa(year) + `/`
	suffix := suffixTasks
	re := regexp.MustCompile((prefix + `[A-Za-z ]+` + suffix))

	rePrefix := regexp.MustCompile((prefix))
	reSuffix := regexp.MustCompile((suffix))

	tmp := re.FindAll([]byte(plainText),-1)

	for _, dirtyName := range tmp {
		ans = append(ans, string(reSuffix.ReplaceAll(rePrefix.ReplaceAll([]byte(dirtyName), nil), nil)))
	}
	return ans
}

const prefixContestants string = `href="people/`
const middleContestants string = `">`
const suffixContestants string = `</a></td>`

/*
	The `ID` and the `Name` of the contestant can be found only under the format `<a href="people/{{.ID}}">{{.Name}}</a>`
	with exception of the first occurence being `<a href="people/add">Add</a>`
*/

func extractIDName(plainTextSlice string) (ID int, Name string) {
	prefix := prefixContestants
	middle := middleContestants
	suffix := suffixContestants
	re := regexp.MustCompile((prefix + `[0-9]+` + middle + `[^<]+` + suffix))
	tmp0 := strings.TrimPrefix(strings.TrimSuffix(string(re.Find([]byte(plainTextSlice))), suffix), prefix)
	tmp1 := strings.Split(tmp0, middle)
	ID, err := strconv.Atoi(tmp1[0])
	if err != nil {
		panic(err.Error())
	}
	Name = tmp1[1]
	return ID, Name
}

const prefixCountry string = `href="(countries|delegations)/[A-Z]+">`
const suffixCountry string = `</a>`

/*
	The `Country` can be found under `<a href="countries/[A-Z]+">{{.Country}}</a>` if we are
	extracting year info else under `<a href="delegations/[A-Z]+/{{.Year}}">{{.Country}}</a>`
*/

func extractCountry(plainTextSlice string) string {
	prefix := prefixCountry
	suffix := suffixCountry
	re := regexp.MustCompile(prefix + `[A-Za-z ]+(Delegations){0}?` + suffix)
	rePrefix := regexp.MustCompile(prefix)
	reSuffix := regexp.MustCompile(suffix)
	return string(reSuffix.ReplaceAll(rePrefix.ReplaceAll(re.Find([]byte(plainTextSlice)),nil),nil))
}

const prefixTaskScore string = `class="[a-z]* ?taskscore">(<a href="tasks/[0-9]+/[A-Za-z ]+">)?`
const suffixTaskScore string = `</td>`

/*
	The scores can be found under `class="[a-z]* ?taskscore">{{.Score}}</td>`
*/

func extractTaskScores(plainTextSlice string) (Scores []float64) {
	prefix := prefixTaskScore
	suffix := suffixTaskScore
	re := regexp.MustCompile((prefix + `[0-9]*.?[0-9]*` + suffix))
	rePrefix := regexp.MustCompile((prefix))
	reSuffix := regexp.MustCompile((suffix))
	tmp := re.FindAll([]byte(plainTextSlice),-1)
	for _, dirtyScore := range tmp {
		number, err := strconv.ParseFloat(string(reSuffix.ReplaceAll(rePrefix.ReplaceAll([]byte(dirtyScore), nil), nil)), 64)
		if err != nil {
			number = 0
		}
		Scores = append(Scores, number)
	}
	return Scores
}

const prefixYear string = `<a href="olympiads/[0-9]+">`
const suffixYear string = `</a>`

/*
	extractYears finds the slices corresponding to every year's info
	plus the years in which contestant has participated in
*/

func extractYears(plainText string) (ans []string,years []int) {
	prefix := prefixYear
	suffix := suffixYear
	re := regexp.MustCompile(prefix + `[0-9]+` + suffix)
	rePrefix := regexp.MustCompile(prefix)
	reSuffix := regexp.MustCompile(suffix)

	tmp := re.FindAllIndex([]byte(plainText), -1)
	for i, Indices := range tmp {
		l := Indices[1]
		var r int
		if i+1 < len(tmp) {
			r = tmp[i+1][0]
		} else {
			r = len(plainText)
		}
		ans = append(ans,plainText[l : r])
		year, err := strconv.Atoi(string(rePrefix.ReplaceAll(reSuffix.ReplaceAll([]byte(plainText),nil),nil)))
		if err != nil {
			panic(err.Error())
		}
		years = append(years,year)
	}
	return ans, years
}

const prefixTitle string = `<title>`
const suffixTitle string = `</title>`

func extractName(plainText string) string {
	prefix := prefixTitle
	suffix := suffixTitle
	
	re := regexp.MustCompile(prefix + `[^<]` + suffix)
	return strings.TrimPrefix(strings.TrimSuffix(string(re.Find([]byte(plainText))),suffix),prefix)
}


//IOIContestant ... struct representing the info related to a contestant's for a year
type IOIContestant struct {
	Country  string
	Name     string
	CFHandle string
	ID       int
	ScoreSum float64
	Rank     int
	Medal    string
	Rel      float64
	Score    []float64
}

//IOIInfo ... struct representing the info related to an IOI year
type IOIInfo struct {
	Tasks            []string
	Gold             int
	Silver           int
	Bronze           int
	NoMedal          int
	ParticipantCount int
}
