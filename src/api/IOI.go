package api

import (
	"regexp"
	"strconv"
	"net/http"
	"errors"
	"io/ioutil"
	"strings"
)

const prefixTasks string = `href="tasks/`
const suffixTasks string = `">`
/*
	The `Tasks` can be found only under the format `<a href="tasks/{{.Year}}/{{.TaskName}}">`
	with exception of the first occurence being <a href="tasks/2019/">
*/

const prefixContestants string = `href="people/`
const middleContestants string = `">`
const suffixContestants string = `</`
/*
	The `ID` and the `Name` of the contestant can be found only under the format `<a href="people/{{.ID}}">{{.Name}}</a>`
	with exception of the first occurence being `<a href="people/add">Add</a>`
*/

const prefixCountry string = `href="countries/[A-Z][A-Z][A-Z]">`
const suffixCountry string = `</`

const prefixTaskScore string = `class="`
const middleTaskScore string = `taskscore">`
const suffixTaskScore string = `</`
const possibleMedals []string = {"gold ","silver ","bronze ",""}
const medalName []string = {"Gold", "Silver", "Bronze", "No Medal"}

const classTypes map[string]string = {
	"Gold" : `class="gold">`,
	"Silver" : `class="silver">`,
	"Bronze" : `class="bronze">`,
	
}

func ExtractYear(year int) (contestants []IOIContestant,info IOIInfo,err error) {
	resp, err := http.Get("https://stats.ioinformatics.org/results/" + strconv.Itoa(year))
	if err != nil {
		return nil,nil,err
	}

	plainText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,nil,err
	}

	rePrefixTasks := regexp.MustCompile(prefixTasks + strconv.Itoa(year) + "/")
	reSuffixTasks := regexp.MustCompile(suffixTasks)

	tmp := rePrefixTasks.FindAllIndices([]byte(plainText),-1)
	for _, Indices  := range tmp[1:] { //The first appearance does not represent a task but the header
		lf := Indices[1]
		rg := reSuffixTasks.FindIndex([]byte(plainText[lf:]))[0]
		rg += lf
		info.Tasks = append(info.Tasks, plainText[lf : rg])
	}

	rePrefixContestants := regexp.MustCompile(prefixContestants)
	
	tmp = rePrefixContestants.FindAllIndices([]byte(plainText))
	for i, Indices := range tmp[1:] { //Again the first appeatance does not represent a task but the header
		lf := Indices[1]
		var rg int
		if i+1 < len(tmp) {
			rg = tmp[i + 1][0]
		} else {
			rg = len(plainText)
		}
		contestant := ExtractParticipantInfo(plainText[lf : rg])
		switch contestant.Medal {
			"Gold":
				info.Gold += 1
			"Silver":
				info.Silver += 1
			"Bronze":
				info.Bronze += 1
			"No Medal":
				info.NoMedal += 1
			default:
				panic("Unknown type of Medal")
		}
		info.ParticipantCount += 1
		contestants = append(contestants,contestant)
	}
	return contestants, info, nil
}

func extractIDName (plainTextSlice string) (ID int,Name string) {
	reMiddleContestants := regexp.MustCompile(middleContestants)
	reSuffixContestants := regexp.MustCompile(suffixContestants)

	middleIndices := reMiddleContestants.FindIndex([]byte(plainTextSlice))
	ID, err := strconv.Atoi(plainTextSlice[:middleIndices[0]])
	if err != nil {
		panic(err.Error())
	}

	suffixIndices := reSuffixContestants.FindIndex([]byte(plainTextSlice))
	Name = plainTextSlice[middleIndices[1] : suffixIndices[0]]

	return ID, Name
}

func extractCountry(plainTextSlice string) string {
	rePrefixCountry := regexp.MustCompile(prefixCountry)
	reSuffixCountry := regexp.MustCompile(suffixCountry)

	preffixIndices := rePrefixCountry.FindIndex([]byte(plainTextSlice))
	suffixIndices = reSuffixCountry.FindIndex([]byte(plainTextSlice[preffixIndices[1]:]))
	
	return plainTextSlice[preffixIndices[1]:suffixIndices[0]]
}

func extractMedalTaskScore(plainTextSlice string) (Medal string,Score []float64) {
	for i,medal := range possibleMedals {
		rePrefixScore := regexp.MustCompile(prefixTaskScore + medal + middleTaskScore)
		scoresPreffixIndices := rePrefixScore.FindAllIndices([]byte(plainTextSlice),-1)
		if scoresPreffixIndices == nil {
			continue
		}
		Medal = medalName[i]
		
		reSuffixScore := regexp.MustCompile(suffixTaskScore)
		for _, preffixIndices := range scoresPreffixIndices {
			suffixIndices = reSuffixScore.FindIndex(plainTextSlice[preffixIndices[1]:])
			taskScore, err := strconv.ParseFloat(plainTextSlice[preffixIndices[1]:suffixIndices[0]], 64)
			if err != nil {
				panic(err.Error())
			}

			Score = append(Score, taskScore)
		}

		break
	}

	return Medal, Score
}

func extractScoreSumReal(plainTextSlice string,classType ) (ScoreSum float64,Rel float64) {

}

func extractParticipantInfo(plainTextSlice string) (ans IOIContestant) {
	ans.ID, ans.Name = extractIDName(plainTextSlice)	

	ans.Country = extractCountry(plainTextSlice)

	ans.Medal, ans.Score = extractMedalTaskScore(plainTextSlice)


}

type IOIContestant struct {
	Country string
	Name string
	CFHandle string
	ID int
	ScoreSum float64
	Rank int
	Medal string
	Rel float64
	Score []float64
}

type IOIInfo struct {
	Tasks []string
	Gold int
	Silver int
	Bronze int
	NoMedal int
	ParticipantCount int
}