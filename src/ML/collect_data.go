package main

import (
	"encoding/csv"
	"github.com/reality95/cf-predictor/src/api"
	"os"
	"strconv"
)

func getProblemCSV(Problems []api.Problem) (ans [][]string) {
	ans = make([][]string, 0)
	tags := make(map[string]int)
	tagNames := make([]string, 0)
	var nTags int = 0
	for _, p := range Problems {
		for _, tag := range p.Tags {
			if _, ok := tags[tag]; !ok {
				tags[tag] = nTags
				tagNames = append(tagNames, tag)
				nTags++
			}
		}
	}
	column := make([]string, 0)
	column = append(column, "Problem Name", "Contest ID", "Index", "Rating", "Points")
	for tag, _ := range tags {
		column = append(column, tag)
	}
	ans = append(ans, column)
	for _, p := range Problems {
		column = make([]string, 0)
		column = append(column, p.Name, strconv.Itoa(p.ContestID), p.Index, strconv.Itoa(p.Rating), strconv.FormatFloat(p.Points, 'f', -1, 64))
		cur_tags := make(map[string]bool)
		for _, tag := range p.Tags {
			cur_tags[tag] = true
		}
		for _, tag := range tagNames {
			column = append(column, strconv.FormatBool(cur_tags[tag]))
		}
		ans = append(ans, column)
	}

	return ans
}

func getScoresCSV(Rows []api.RanklistRow, Problems []api.Problem) (ans [][]string) {
	ans = make([][]string, 0)
	column := make([]string, 0)
	column = append(column, "Handle", "Rank", "Points", "Penalty")
	for _, p := range Problems {
		column = append(column, "Score for task "+p.Name)
	}
	ans = append(ans, column)
	for _, row := range Rows {
		if row.Partyv.Len() != 1 {
			continue
		}
		column = make([]string, 0)
		column = append(column, row.Partyv.Members[0].Handle, strconv.Itoa(row.Rank), strconv.FormatFloat(row.Points, 'f', -1, 64), strconv.Itoa(row.Penalty))
		for _, result := range row.ProblemResults {
			column = append(column, strconv.FormatFloat(result.Points, 'f', -1, 64))
		}
		ans = append(ans, column)
	}
	return ans
}

func getRatingsCSV(changes []api.RatingChange) (ans [][]string) {
	ans = make([][]string, 0)
	ans = append(ans, []string{"Handle", "Old Rating", "New Rating"})
	for _, c := range changes {
		ans = append(ans, []string{c.Handle, strconv.Itoa(c.OldRating), strconv.Itoa(c.NewRating)})
	}
	return ans
}

func collectSampleProblems(contestID int, room interface{}) {
	contestStats, err := api.GetContestStandings(contestID, nil, nil, nil, room, nil)
	if err != nil {
		panic(err.Error())
	}
	changes, err := api.GetRatingChanges(contestID)
	if err != nil {
		panic(err.Error())
	}
	fProblems, err := os.Create("assets/ML/SampleProblems/problems.csv")
	if err != nil {
		panic(err.Error())
	}
	wProblems := csv.NewWriter(fProblems)
	defer fProblems.Close()
	fScores, err := os.Create("assets/ML/SampleProblems/scores.csv")
	if err != nil {
		panic(err.Error())
	}
	wScores := csv.NewWriter(fScores)
	defer fScores.Close()
	fRating, err := os.Create("assets/ML/SampleProblems/rating.csv")
	if err != nil {
		panic(err.Error())
	}
	wRating := csv.NewWriter(fRating)
	defer fRating.Close()

	wProblems.WriteAll(getProblemCSV(contestStats.Problems))
	wScores.WriteAll(getScoresCSV(contestStats.Rows, contestStats.Problems))
	wRating.WriteAll(getRatingsCSV(changes))
}

func main() {
	collectSampleProblems(1329, nil)
}
