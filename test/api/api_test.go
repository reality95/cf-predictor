package main

import (
	"github.com/reality95/cf-predictor/src/api"
	"github.com/stretchr/testify/assert"
	_ "log"
	"testing"
	"time"
)

func TestGetComments(t *testing.T) {
	assert := assert.New(t)

	comments, err := api.GetComments(69)
	assert.True(err == nil, "Expected no error while extracting comments for blogId 69")
	assert.True(len(comments) == 0, "Blog with blogId 69 has no comments, received %d comments\n", len(comments))

	comments, err = api.GetComments(666)
	assert.True(err == nil, "Expected no error while extracting comments for blogId 666")
	assert.Truef(len(comments) == 5, "Blog with blogId 666 has 5 comments, received %d comments\n", len(comments))

	assert.Equal(comments[0].Id, 10505)
	assert.Equal(comments[0].CreationTimeSeconds, 1284399381)
	assert.Equal(comments[0].CommentatorHandle, "cerealguy")
	assert.Equal(comments[0].Locale, "en")
	assert.Equal(comments[0].Rating, 3)
	assert.Equal(comments[1].ParentCommentId, 10505)

	for i, c := range comments {
		assert.True(c.Text != "", "The comment number %d must not be empty\n", i)
	}

	_, err = api.GetComments(666666)
	assert.True(err != nil, "Expected an error while extracting comments for blogId 666666, received none\n")
	assert.True(err.Error() == "blogEntryId: Blog entry with id 666666 not found", "Expected a different error while extracting comments for blogId 666666\n")
}

func TestGetBlog(t *testing.T) {
	assert := assert.New(t)

	blog, err := api.GetBlog(69)
	assert.True(err == nil, "Expected no error while extracting BlogEntry for blogId 69")
	assert.Equal(blog.OriginalLocale, "ru")
	assert.Equal(blog.AllowViewHistory, false)
	assert.Equal(blog.CreationTimeSeconds, 1265647999)
	assert.Equal(blog.Rating, 1)
	assert.Equal(blog.AuthorHandle, "MikeMirzayanov")
	assert.Equal(blog.ModificationTimeSeconds, 1265653678)
	assert.Equal(blog.Id, 69)
	assert.Equal(blog.Title, "Back home!")
	assert.Equal(blog.Locale, "en")
	assert.Equal(blog.Tags, []string{"2010", "acm", "acm-icpc", "back home", "saratov"})

	_, err = api.GetBlog(666666)

	assert.True(err != nil, "Expected an error while extracting BlogEntru for blogId 666666, received none\n")
	assert.True(err.Error() == ("blogEntryId: Blog entry with id 666666 not found"), "Expected a different error while extracting comments for blogId 666666\n")

}

func TestGetHacks(t *testing.T) {
	time.Sleep(1000 * time.Millisecond)
	hacks, err := api.GetHacks(566)
	assert := assert.New(t)
	assert.True(err == nil, "Expected no error while extracting hacks from contestId 566")
	assert.Equal(len(hacks), 325)
	assert.Equal(hacks[0].Id, 160426)
	assert.Equal(hacks[0].CreationTimeSeconds, 1438274514)
	assert.Equal(hacks[0].Hacker, api.Party{
		ContestId:        566,
		Members:          []api.Member{{Handle: "Sehnsucht"}},
		ParticipantType:  "CONTESTANT",
		Ghost:            false,
		Room:             29,
		StartTimeSeconds: 1438273200,
	})
	assert.Equal(hacks[0].Verdict, "INVALID_INPUT")
	assert.Equal(hacks[0].Defender, api.Party{
		ContestId:        566,
		Members:          []api.Member{{Handle: "osama"}},
		ParticipantType:  "CONTESTANT",
		Ghost:            false,
		Room:             29,
		StartTimeSeconds: 1438273200,
	})
	assert.Equal(hacks[0].Problemv, api.Problem{ContestId: 566, Index: "F", Name: "Clique in the Divisibility Graph", Typev: "PROGRAMMING", Points: 500.0, Rating: 1500, Tags: []string{"dp", "math", "number theory"}})
	assert.Equal(hacks[0].Test, "4\r\n2 2 4 4\r\n\n")
	assert.Equal(hacks[0].JudgeProtocol, api.JProtocol{
		Protocol: "Validator \u0027validate.exe\u0027 returns exit code 3 [FAIL Integer parameter [name\u003da[1]] equals to 2, violates the range [3, 1000000] (stdin)]",
		Manual:   "true",
		Verdict:  "Invalid input",
	})
	assert.Equal(hacks[5].JudgeProtocol.Manual, "false")

	hacks, err = api.GetHacks(6669)
	assert.True(err != nil, "Expected error while extracting hacks from contestId 6669, got none")
	assert.True(err.Error() == "contestId: Contest with id 6669 not found", "Expected a different error while extracting Hacks from contestId 6669\n")
}

func TestGetRatingChanges(t *testing.T) {
	assert := assert.New(t)
	changes, err := api.GetRatingChanges(566)
	assert.True(err == nil, "Expected no error while extracting RatingChanges from contestId 566")
	assert.Equal(len(changes), 761)
	assert.Equal(changes[0], api.RatingChange{
		ContestId:   566,
		ContestName: "VK Cup 2015 - Finals, online mirror",
		Handle:      "rng_58",
		Rank:        1,
		RatingUpdateTimeSeconds: 1438284000,
		OldRating:               2849,
		NewRating:               2941,
	})

	_, err = api.GetRatingChanges(6669)
	assert.True(err != nil, "Expected error while extracting hacks from contestId 6669, got none")
	assert.True(err.Error() == "contestId: Contest with id 6669 not found", "Expected a different error while extracting RatingChanges from contestId 6669\n")
}

func TestGetContestStandings(t *testing.T) {
	assert := assert.New(t)
	CStandings, err := api.GetContestStandings(566, nil, 1, 5, nil, nil)
	assert.True(err == nil, "Expected no error while extracting ContestStandings from contestId 566")
	assert.Equal(len(CStandings.Rows), 5)
	assert.Equal(CStandings.Problems, []api.Problem{
		{
			ContestId: 566,
			Index:     "A",
			Name:      "Matching Names",
			Typev:     "PROGRAMMING",
			Points:    1750.0,
			Rating:    2300,
			Tags:      []string{"dfs and similar", "strings", "trees"},
		},
		{
			ContestId: 566,
			Index:     "B",
			Name:      "Replicating Processes",
			Typev:     "PROGRAMMING",
			Points:    2500.0,
			Rating:    2600,
			Tags:      []string{"constructive algorithms", "greedy"},
		},
		{
			ContestId: 566,
			Index:     "C",
			Name:      "Logistical Questions",
			Typev:     "PROGRAMMING",
			Points:    3000.0,
			Rating:    3000,
			Tags:      []string{"dfs and similar", "divide and conquer", "trees"},
		},
		{
			ContestId: 566,
			Index:     "D",
			Name:      "Restructuring Company",
			Typev:     "PROGRAMMING",
			Points:    1000.0,
			Rating:    1900,
			Tags:      []string{"data structures", "dsu"},
		},
		{
			ContestId: 566,
			Index:     "E",
			Name:      "Restoring Map",
			Typev:     "PROGRAMMING",
			Points:    3000.0,
			Rating:    3200,
			Tags:      []string{"bitmasks", "constructive algorithms", "trees"},
		},
		{
			ContestId: 566,
			Index:     "F",
			Name:      "Clique in the Divisibility Graph",
			Typev:     "PROGRAMMING",
			Points:    500.0,
			Rating:    1500,
			Tags:      []string{"dp", "math", "number theory"},
		},
		{
			ContestId: 566,
			Index:     "G",
			Name:      "Max and Min",
			Typev:     "PROGRAMMING",
			Points:    2500.0,
			Rating:    2500,
			Tags:      []string{"geometry"},
		},
	})
	contest := CStandings.Contestv
	assert.Equal(contest.Id, 566)
	assert.Equal(contest.Name, "VK Cup 2015 - Finals, online mirror")
	assert.Equal(contest.Typev, "CF")
	assert.Equal(contest.Phase, "FINISHED")
	assert.Equal(contest.Frozen, false)
	assert.Equal(contest.DurationSeconds, 10800)
	assert.Equal(contest.StartTimeSeconds, 1438273200)

	assert.Equal(CStandings.Rows[0], api.RanklistRow{
		Partyv: api.Party{
			ContestId: 566,
			Members: []api.Member{
				{
					Handle: "rng_58",
				},
			},
			ParticipantType:  "CONTESTANT",
			Ghost:            false,
			Room:             8,
			StartTimeSeconds: 1438273200,
		},
		Rank:                  1,
		Points:                7974.0,
		Penalty:               0,
		SuccessfulHackCount:   1,
		UnsuccessfulHackCount: 0,
		ProblemResults: []api.ProblemResult{
			{
				Points:               1330.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
				BestSubmissionTimeSeconds: 3624,
			},
			{
				Points:               1600.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
				BestSubmissionTimeSeconds: 5422,
			},
			{
				Points:               1404.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
				BestSubmissionTimeSeconds: 7991,
			},
			{
				Points:               840.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
				BestSubmissionTimeSeconds: 2447,
			},
			{
				Points:               0.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
			},
			{
				Points:               490.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
				BestSubmissionTimeSeconds: 339,
			},
			{
				Points:               2210.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
				BestSubmissionTimeSeconds: 1757,
			},
		},
	})

	CStandings, err = api.GetContestStandings(566, nil, nil, 5, nil, nil)
	assert.True(err != nil, "If end is not nil and start is nil then it should return an error\n")
	assert.Equal(err.Error(), "If end is not nil, start must not be nil as well")

	time.Sleep(time.Second)

	CStandings, err = api.GetContestStandings(566, []string{"rng_58", "Errichto", "I_Love_Tina"}, nil, nil, nil, nil)
	assert.True(err == nil, "Expected no error while getting contest standings for handles rng_58;Erricho;I_Love_Tina")
	assert.Equal(len(CStandings.Rows), 3)

	CStandings, err = api.GetContestStandings(566, []string{"rng_58", "Errichto", "I_Love_Tina", "I_Love_Tina", "rng_58"}, nil, nil, nil, nil)
	assert.True(err == nil, "Expected no error while getting contest standings for handles rng_58;Erricho;I_Love_Tina with duplicates")
	assert.Equal(len(CStandings.Rows), 3)

	CStandings, err = api.GetContestStandings(566, nil, 69, nil, nil, nil)
	assert.True(err == nil, "Expected no error while getting contest standings starting at position 69")
	assert.Equal(len(CStandings.Rows), 557)

	CStandings, err = api.GetContestStandings(566, nil, 69, 666, nil, true)
	assert.True(err == nil, "Expected no error while getting contest standings between positions 69 and 666 showing unofficial standings")
	assert.Equal(len(CStandings.Rows), 666-69+1)

	CStandings, err = api.GetContestStandings(566, nil, nil, nil, 20, nil)
	assert.True(err == nil, "Expected no error while getting contest standings from room 5")
	assert.Equal(len(CStandings.Rows), 40)

	time.Sleep(time.Second)

	CStandings, err = api.GetContestStandings(6669, nil, nil, nil, nil, nil)
	assert.True(err != nil, "Expected error while getting results from a contest with Id 6669, got none")
	assert.Equal(err.Error(), "contestId: Contest with id 6669 not found")

	CStandings, err = api.GetContestStandings(566, make([]string, 10005), nil, nil, nil, nil)
	assert.True(err != nil, "Expected error while querying more than 10000 handles, got none")
	assert.Equal(err.Error(), "Expected at most 10000 handles")

	CStandings, err = api.GetContestStandings(566, nil, "", nil, nil, nil)
	assert.True(err != nil, "Expected error if start is not int")

	CStandings, err = api.GetContestStandings(566, nil, nil, "", nil, nil)
	assert.True(err != nil, "Expected error if end is not int")

	CStandings, err = api.GetContestStandings(566, nil, nil, nil, "", nil)
	assert.True(err != nil, "Expected error if room is not int")

	CStandings, err = api.GetContestStandings(566, nil, nil, nil, nil, 0)
	assert.True(err != nil, "Expected error if showUnofficial is not bool")
}

func TestGetContestStatus(t *testing.T) {
	assert := assert.New(t)
	submissions, err := api.GetContestStatus(566, nil, 3, 13)
	assert.True(err == nil, "Expected no error while extracting submissions from contest 566, start = 3,end = 13")
	assert.Truef(len(submissions) == 11, "Expected 10 submissions while extracting submissions from contest 566, start = 3,end = 13, got %d", len(submissions))

	submissions, err = api.GetContestStatus(1179, "I_Love_Tina", 1, 1)
	assert.True(err == nil, "Expected no error while extracting submissions from contest 1179 with handle I_Love_Tina the most recent submission")
	assert.Truef(len(submissions) == 1, "Expected one submission while extracting the most recent submission from contest 1179 with handle I_Love_Tina, got %d", len(submissions))
	assert.Equal(submissions[0], api.Submission{
		Id:                  56038919,
		ContestId:           1179,
		CreationTimeSeconds: 1561469037,
		RelativeTimeSeconds: submissions[0].RelativeTimeSeconds,
		Problemv: api.Problem{
			ContestId: 1179,
			Index:     "E",
			Name:      "Alesya and Discrete Math",
			Typev:     "PROGRAMMING",
			Points:    2250.0,
			Rating:    3200,
			Tags:      []string{"divide and conquer", "interactive"},
		},
		Author: api.Party{
			ContestId: 1179,
			Members: []api.Member{
				{
					Handle: "I_Love_Tina",
				},
			},
			ParticipantType:  "PRACTICE",
			Ghost:            false,
			StartTimeSeconds: 1561136700,
		},
		ProgrammingLanguage: "GNU C++17",
		Verdict:             "RUNTIME_ERROR",
		Testset:             "TESTS",
		PassedTestCount:     2,
		TimeConsumedMillis:  15,
		MemoryConsumedBytes: 1638400,
	})

	_, err = api.GetContestStatus(6669, nil, nil, nil)
	assert.True(err != nil, "Expected erorr while trying to extract submissions from contest with ID 6669")
	assert.True(err.Error() == "contestId: Contest with id 6669 not found", "Expected a different erorr while trying to extract submissions from contest with ID 6669")
}

func TestGetPsetProblems(t *testing.T) {
	assert := assert.New(t)
	pStats, problems, err := api.GetPsetProblems([]string{"brute%20force", "math", "implementation", "binary%20search"}, nil)
	assert.True(err == nil, "Expected no error while getting the problems with tags 'brute force','math','implementation','binary search', got none")
	assert.Equal(len(pStats), len(problems))
	assert.True(len(problems) > 0, "Expected at least one problem while getting the problems with tags 'brute force','math','implementation','binary search'")

	nProblems := len(problems)

	time.Sleep(time.Second)

	pStats, problems, err = api.GetPsetProblems([]string{"brute%20force", "math", "math", "implementation", "binary%20search", "math"}, nil)

	assert.True(err == nil, "Expected no error while getting the problems with tags 'brute force','math','implementation','binary search' with duplicates, got none")
	assert.Truef(nProblems == len(problems), "Expected the same result if we add duplicate tags, expected %d, got %d", nProblems, len(problems))

	pStats, problems, err = api.GetPsetProblems([]string{"brute%20force", "math", "math", "implementation", "binary%20search", "math"}, "acmsguru")
	assert.True(err == nil, "Expected no error while getting the problems with tags 'brute force','math','implementation','binary search' with duplicates and psetName 'acmsguru', got none")

	pStats, problems, err = api.GetPsetProblems([]string{"hard%20problem"}, nil)
	assert.True(err == nil, "Expected no error while getting the problems with tag 'hard problem'")
	assert.True(len(problems) == 0 && len(pStats) == 0, "Expected no problem while getting the problems with tag 'hard problem'")

	_, _, err = api.GetPsetProblems(nil, 3.14)
	assert.True(err != nil, "Expected error if psetName type is not string")
	assert.True(err.Error() == "problemsetName must have type string", "Expected a different error if psetName type is not string")

}

/*
func TestGetPsetRecentStatus(t *testing.T) {
	assert := assert.New(t)

}
*/
