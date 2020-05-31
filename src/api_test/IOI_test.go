package api_test

import (
	"github.com/reality95/cf-predictor/src/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractYearInfo(t *testing.T) {
	assert := assert.New(t)
	contestants, info, err := api.ExtractYearInfo(2019, false)
	assert.Equal(err, nil, "Excepted no error while extracting info from year 2019")
	assert.Equal(info, api.IOIInfo{
		Gold:             28,
		Silver:           54,
		Bronze:           81,
		NoMedal:          168,
		ParticipantCount: 331,
		Tasks:            []string{"shoes", "split", "rect", "line", "vision", "walk"},
	})
	assert.Equal(contestants[0], api.IOIContestant{
		Name:     "Benjamin Qi",
		ID:       6716,
		ScoreSum: 547.09,
		Rank:     1,
		Medal:    "Gold",
		Rel:      91.18,
		Country:  "United States of America",
		Score:    []float64{100, 100, 100, 90.09, 100, 57},
		CFHandle: "",
	})

	benHandle, err := api.ExtractCFHandle(6716)
	assert.Equal(err, nil, "Excepted no error while extracting handle of Benjamin Qi")
	assert.Equalf(benHandle, "Benq", "Excepted `Benq` CFHandle for Benjamin Qi, got `%s`", benHandle)

	_, _, err = api.ExtractYearInfo(1969, false)
	assert.Truef(err != nil, "Excepted error when `year` is smaller than 1989")
	assert.Equal(err.Error(), "IOI started at 1989, so year must be at least 1989", "Excepted a different error when `year` is smaller than 1989")
}

func TestExtractContestantsInfo(t *testing.T) {
	assert := assert.New(t)
	contestantInfo, err := api.ExtractContestantInfo(5968)
	assert.Equal(err, nil, "Excepted no error while extracting info from contestant with `ID` 5968")
	assert.Equal(len(contestantInfo), 4, "Participant with `ID` 5968 has 4 year participations as a contestant")
	assert.Equal(contestantInfo[2017], api.IOIContestant{
		Name:     "Gabriel Cojocaru",
		ID:       5968,
		ScoreSum: 271.70,
		Rank:     65,
		Medal:    "Silver",
		Rel:      45.28,
		Country:  "Moldova",
		Score:    []float64{26.56, 100, 5, 98.14, 30, 12},
		CFHandle: "I_Love_Tina",
	})

	contestantInfo, err = api.ExtractContestantInfo(681)
	assert.Equal(err, nil, "Excepted no error while extracting info from contestant with `ID` 681")
	assert.Equal(len(contestantInfo), 4, "Participant with `ID` 681 has 4 year participations as a contestant")
	assert.Equal(contestantInfo[2011], api.IOIContestant{
		Name:     "Eduards Kaļiņičenko",
		ID:       681,
		ScoreSum: 491,
		Rank:     21,
		Medal:    "Gold",
		Rel:      81.83,
		Country:  "Latvia",
		Score:    []float64{100, 43, 100, 100, 50, 98},
		CFHandle: "eduardische",
	})
}
