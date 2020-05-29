package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const prefixContestantsInfo string = `<td( class="[a-z]")?(title="Unofficial participation")?>`
const suffixContestantsInfo string = `</td>`

func ExtractYearInfo(year int) (contestants []IOIContestant, info IOIInfo, err error) {
	resp, err := http.Get("https://stats.ioinformatics.org/results/" + strconv.Itoa(year))
	if err != nil {
		return nil, nil, err
	}

	plainText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	info.Tasks = extractTaskNames(plainText, year)

	prefix := prefixContestantsInfo
	suffix := suffixContestantsInfo

	re := regexp.MustCompile([]byte(prefix + `([A-Za-z]*)|([0-9]+.?[0-9]*%?)` + suffix))

	rePrefix := regexp.MustCompile([]byte(prefix))
	reSuffix := regexp.MustCompile([]byte(suffix))

	tmp := re.FindAllIndex(plainText)

	if len(tmp)%4 != 0 {
		panic("Expected 4 fields with data for every contestant")
	}

	info.ParticipantCount = len(tmp) / 4

	for i := 0; i < info.ParticipantCount; i++ {
		for j := 0; j < 4; j++ {
			l, r := tmp[4*i+j][0], tmp[4*i+j][1]
			sValue := string(reSuffix.ReplaceAll(rePrefix.ReplaceAll([]byte(dirtyName), nil), nil))
			var contestant IOIContestant
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
		plainTextSlice := plainText[tmp[4*i][0]:tmp[4*i+3][1]]
		contestant.ID, contestant.Name = extractIDName(plainTextSlice)
		contestant.Score = extractTaskScores(plainTextSlice)
		contestant.Country = extractCountry(plainTextSlice)
	}

	return contestants, info, nil
}

const prefixCFHandle string = `https://codeforces.com/profile/`
const suffixCFHandle string = `">`

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

	re := regexp.MustCompile([]byte(prefix + "[A-Za-z0-9_-]+" + suffix))

	return string.TrimPrefix(strings.TrimSuffix(re.Find(plainText), suffix), prefix), nil
}

const prefixTasks string = `href="tasks/`
const suffixTasks string = `">`

/*
	The `Tasks` can be found only under the format `<a href="tasks/{{.Year}}/{{.TaskName}}">`
	with exception of the first occurence being <a href="tasks/2019/">
*/

func extractTaskNames(plainTextSlice string, year int) (ans []string) {
	prefix := prefixTasks + strconv.Itoa(year)
	suffix := suffixTasks
	re := regexp.MustCompile([]byte(prefix + `[A-Za-z]+` + suffix))

	rePrefix := regexp.MustCompile([]byte(prefix))
	reSuffix := regexp.MustCompile([]byte(suffix))

	tmp := re.FindAll(plainTextSlice)

	for _, dirtyName := range tmp {
		ans = append(ans, string(reSuffix.ReplaceAll(rePrefix.ReplaceAll([]byte(dirtyName), nil), nil)))
	}
	return ans
}

const prefixContestants string = `href="people/`
const middleContestants string = `">`
const suffixContestants string = `</a>`

/*
	The `ID` and the `Name` of the contestant can be found only under the format `<a href="people/{{.ID}}">{{.Name}}</a>`
	with exception of the first occurence being `<a href="people/add">Add</a>`
*/

func extractIDName(plainTextSlice string) (ID int, Name string) {
	prefix := prefixContestants
	middle := middleContestants
	suffix := suffixContestants
	re := regexp.MustCompile([]byte(prefix + `[0-9]+` + middle + `[A-Za-z]+` + suffix))
	tmp0 := strings.TrimPrefix(string.TrimSuffix(re.Find(plainTextSlice), suffix), prefix)
	tmp1 := strings.Split(tmp0, middle)
	return strconv.Atoi(tmp1[0]), tmp[1]
}

const prefixCountry string = `href="countries/[A-Z]+">`
const suffixCountry string = `</a>`

func extractCountry(plainTextSlice string) string {
	prefix := prefixCountry
	suffix := suffixCountry
	re := regexp.MustCompile([]byte(prefix + `[A-Z]+` + suffix))
	return strings.TrimPrefix(string.TrimSuffix(re.Find(plainTextSlice), suffix), prefix)
}

const prefixTaskScore string = `class="[a-z]* ?taskscore">`
const suffixTaskScore string = `</td>`

func extractTaskScores(plainTextSlice string) (Scores []float64) {
	prefix := prefixTaskScore
	suffix := suffixTaskScore
	re := regexp.MustCompile([]byte(prefix + `[0-9]*.?[0-9]*` + suffix))
	rePrefix := regexp.MustCompile([]byte(prefix))
	reSuffix := regexp.MustCompile([]byte(suffix))
	tmp := re.FindAllIndex([]byte(plainTextSlice))
	for _, dirtyScore := range tmp {
		number, err := strconv.ParseFloat(string(reSuffix.ReplaceAll(rePrefix.ReplaceAll([]byte(dirtyName), nil), nil)), 64)
		if err != nil {
			number = 0
		}
		Scores = append(Scores, number)
	}
	return Scores
}

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

type IOIInfo struct {
	Tasks            []string
	Gold             int
	Silver           int
	Bronze           int
	NoMedal          int
	ParticipantCount int
}
