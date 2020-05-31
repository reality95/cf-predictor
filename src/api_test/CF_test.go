package api_test

import (
	"github.com/reality95/cf-predictor/src/api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetComments(t *testing.T) {
	assert := assert.New(t)

	comments, err := api.GetComments(69)
	assert.True(err == nil, "Expected no error while extracting comments for `blogID` 69")
	assert.True(len(comments) == 0, "Blog with blogID 69 has no comments, received %d comments\n", len(comments))

	comments, err = api.GetComments(666)
	assert.True(err == nil, "Expected no error while extracting comments for `blogID` 666")
	assert.Truef(len(comments) == 5, "Blog with blogID 666 has 5 comments, received %d comments\n", len(comments))

	assert.Equal(comments[0].ID, 10505)
	assert.Equal(comments[0].CreationTimeSeconds, 1284399381)
	assert.Equal(comments[0].CommentatorHandle, "cerealguy")
	assert.Equal(comments[0].Locale, "en")
	assert.Equal(comments[0].Rating, 3)
	assert.Equal(comments[1].ParentCommentID, 10505)

	for i, c := range comments {
		assert.Truef(c.Text != "", "The comment number %d must not be empty\n", i)
	}

	_, err = api.GetComments(666666)
	assert.True(err != nil, "Expected an error while extracting comments for `blogID` 666666, received none\n")
	assert.Equal(err.Error(), "blogEntryId: Blog entry with id 666666 not found", "Expected a different error while extracting comments for blogID 666666\n")
}

func TestGetBlog(t *testing.T) {
	assert := assert.New(t)

	blog, err := api.GetBlog(69)
	assert.True(err == nil, "Expected no error while extracting BlogEntry for `blogID` 69")
	assert.Equal(blog.OriginalLocale, "ru")
	assert.Equal(blog.AllowViewHistory, false)
	assert.Equal(blog.CreationTimeSeconds, 1265647999)
	assert.Equal(blog.Rating, 1)
	assert.Equal(blog.AuthorHandle, "MikeMirzayanov")
	assert.Equal(blog.ModificationTimeSeconds, 1265653678)
	assert.Equal(blog.ID, 69)
	assert.Equal(blog.Title, "Back home!")
	assert.Equal(blog.Locale, "en")
	assert.Equal(blog.Tags, []string{"2010", "acm", "acm-icpc", "back home", "saratov"})

	_, err = api.GetBlog(666666)

	assert.True(err != nil, "Expected an error while extracting BlogEntry for `blogI`D 666666, received none\n")
	assert.Equal(err.Error(), "blogEntryId: Blog entry with id 666666 not found", "Expected a different error while extracting comments for `blogID` 666666\n")

	time.Sleep(time.Second)
}

func TestGetHacks(t *testing.T) {
	assert := assert.New(t)
	hacks, err := api.GetHacks(566)
	assert.True(err == nil, "Expected no error while extracting hacks from `contestID` 566")
	assert.Equal(len(hacks), 325)
	assert.Equal(hacks[0].ID, 160426)
	assert.Equal(hacks[0].CreationTimeSeconds, 1438274514)
	assert.Equal(hacks[0].Hacker, api.Party{
		ContestID: 566,
		Members: []api.Member{
			{
				Handle: "Sehnsucht",
			},
		},
		ParticipantType:  "CONTESTANT",
		Ghost:            false,
		Room:             29,
		StartTimeSeconds: 1438273200,
	})
	assert.Equal(hacks[0].Verdict, "INVALID_INPUT")
	assert.Equal(hacks[0].Defender, api.Party{
		ContestID: 566,
		Members: []api.Member{
			{
				Handle: "osama",
			},
		},
		ParticipantType:  "CONTESTANT",
		Ghost:            false,
		Room:             29,
		StartTimeSeconds: 1438273200,
	})
	assert.Equal(hacks[0].Problemv, api.Problem{
		ContestID: 566,
		Index:     "F",
		Name:      "Clique in the Divisibility Graph",
		Typev:     "PROGRAMMING",
		Points:    500.0,
		Rating:    1500,
		Tags:      []string{"dp", "math", "number theory"},
	})
	assert.Equal(hacks[0].Test, "4\r\n2 2 4 4\r\n\n")
	assert.Equal(hacks[0].JudgeProtocol, api.JProtocol{
		Protocol: "Validator \u0027validate.exe\u0027 returns exit code 3 [FAIL Integer parameter [name\u003da[1]] equals to 2, violates the range [3, 1000000] (stdin)]",
		Manual:   "true",
		Verdict:  "Invalid input",
	})
	assert.Equal(hacks[5].JudgeProtocol.Manual, "false")

	_, err = api.GetHacks(6669)
	assert.True(err != nil, "Expected error while extracting hacks from `contestID` 6669, got none")
	assert.Equal(err.Error(), "contestId: Contest with id 6669 not found", "Expected a different error while extracting Hacks from `contestID` 6669\n")
}

func TestGetRatingChanges(t *testing.T) {
	assert := assert.New(t)
	changes, err := api.GetRatingChanges(566)
	assert.True(err == nil, "Expected no error while extracting RatingChanges from `contestID` 566")
	assert.Equal(len(changes), 761)
	assert.Equal(changes[0], api.RatingChange{
		ContestID:               566,
		ContestName:             "VK Cup 2015 - Finals, online mirror",
		Handle:                  "rng_58",
		Rank:                    1,
		RatingUpdateTimeSeconds: 1438284000,
		OldRating:               2849,
		NewRating:               2941,
	})

	_, err = api.GetRatingChanges(6669)
	assert.True(err != nil, "Expected error while extracting hacks from `contestID` 6669, got none")
	assert.Equal(err.Error(), "contestId: Contest with id 6669 not found", "Expected a different error while extracting RatingChanges from `contestID` 6669\n")
}

func TestGetContestStandings(t *testing.T) {
	assert := assert.New(t)
	CStandings, err := api.GetContestStandings(566, nil, 1, 5, nil, nil)
	assert.True(err == nil, "Expected no error while extracting ContestStandings from `contestID` 566")
	assert.Equal(len(CStandings.Rows), 5, "Expected 5 rows when extracting ContestStandings from `contestID` 566 with `start` = 1 and `end` = 5")
	assert.Equal(CStandings.Problems, []api.Problem{
		{
			ContestID: 566,
			Index:     "A",
			Name:      "Matching Names",
			Typev:     "PROGRAMMING",
			Points:    1750.0,
			Rating:    2300,
			Tags:      []string{"dfs and similar", "strings", "trees"},
		},
		{
			ContestID: 566,
			Index:     "B",
			Name:      "Replicating Processes",
			Typev:     "PROGRAMMING",
			Points:    2500.0,
			Rating:    2600,
			Tags:      []string{"constructive algorithms", "greedy"},
		},
		{
			ContestID: 566,
			Index:     "C",
			Name:      "Logistical Questions",
			Typev:     "PROGRAMMING",
			Points:    3000.0,
			Rating:    3000,
			Tags:      []string{"dfs and similar", "divide and conquer", "trees"},
		},
		{
			ContestID: 566,
			Index:     "D",
			Name:      "Restructuring Company",
			Typev:     "PROGRAMMING",
			Points:    1000.0,
			Rating:    1900,
			Tags:      []string{"data structures", "dsu"},
		},
		{
			ContestID: 566,
			Index:     "E",
			Name:      "Restoring Map",
			Typev:     "PROGRAMMING",
			Points:    3000.0,
			Rating:    3200,
			Tags:      []string{"bitmasks", "constructive algorithms", "trees"},
		},
		{
			ContestID: 566,
			Index:     "F",
			Name:      "Clique in the Divisibility Graph",
			Typev:     "PROGRAMMING",
			Points:    500.0,
			Rating:    1500,
			Tags:      []string{"dp", "math", "number theory"},
		},
		{
			ContestID: 566,
			Index:     "G",
			Name:      "Max and Min",
			Typev:     "PROGRAMMING",
			Points:    2500.0,
			Rating:    2500,
			Tags:      []string{"geometry"},
		},
	})
	contest := CStandings.Contestv
	assert.Equal(contest.ID, 566)
	assert.Equal(contest.Name, "VK Cup 2015 - Finals, online mirror")
	assert.Equal(contest.Typev, "CF")
	assert.Equal(contest.Phase, "FINISHED")
	assert.Equal(contest.Frozen, false)
	assert.Equal(contest.DurationSeconds, 10800)
	assert.Equal(contest.StartTimeSeconds, 1438273200)

	assert.Equal(CStandings.Rows[0], api.RanklistRow{
		Partyv: api.Party{
			ContestID: 566,
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
				Points:                    1330.0,
				RejectedAttemptCount:      0,
				Typev:                     "FINAL",
				BestSubmissionTimeSeconds: 3624,
			},
			{
				Points:                    1600.0,
				RejectedAttemptCount:      0,
				Typev:                     "FINAL",
				BestSubmissionTimeSeconds: 5422,
			},
			{
				Points:                    1404.0,
				RejectedAttemptCount:      0,
				Typev:                     "FINAL",
				BestSubmissionTimeSeconds: 7991,
			},
			{
				Points:                    840.0,
				RejectedAttemptCount:      0,
				Typev:                     "FINAL",
				BestSubmissionTimeSeconds: 2447,
			},
			{
				Points:               0.0,
				RejectedAttemptCount: 0,
				Typev:                "FINAL",
			},
			{
				Points:                    490.0,
				RejectedAttemptCount:      0,
				Typev:                     "FINAL",
				BestSubmissionTimeSeconds: 339,
			},
			{
				Points:                    2210.0,
				RejectedAttemptCount:      0,
				Typev:                     "FINAL",
				BestSubmissionTimeSeconds: 1757,
			},
		},
	})

	time.Sleep(time.Second)

	CStandings, err = api.GetContestStandings(566, nil, nil, 5, nil, nil)
	assert.True(err != nil, "If `end` is not nil and `start` is nil then it should return an error\n")
	assert.Equal(err.Error(), "If `end` is not nil, `start` must not be nil as well")

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

	time.Sleep(time.Second)

	CStandings, err = api.GetContestStandings(566, nil, nil, nil, 20, nil)
	assert.True(err == nil, "Expected no error while getting contest standings from `room` 5")
	assert.Equal(len(CStandings.Rows), 40)

	CStandings, err = api.GetContestStandings(6669, nil, nil, nil, nil, nil)
	assert.True(err != nil, "Expected error while getting results from a contest with `contestID` 6669, got none")
	assert.Equal(err.Error(), "contestId: Contest with id 6669 not found")

	CStandings, err = api.GetContestStandings(566, make([]string, 10005), nil, nil, nil, nil)
	assert.True(err != nil, "Expected error while querying ContestStandings from more than 10000 handles, got none")
	assert.Equal(err.Error(), "Expected at most 10000 handles")

	CStandings, err = api.GetContestStandings(566, nil, "", nil, nil, nil)
	assert.True(err != nil, "Expected error if `start` is not int")

	CStandings, err = api.GetContestStandings(566, nil, nil, "", nil, nil)
	assert.True(err != nil, "Expected error if `end` is not int")

	time.Sleep(time.Second)

	CStandings, err = api.GetContestStandings(566, nil, nil, nil, "", nil)
	assert.True(err != nil, "Expected error if room is not int")

	CStandings, err = api.GetContestStandings(566, nil, nil, nil, nil, 0)
	assert.True(err != nil, "Expected error if `showUnofficial` is not bool")
}

func TestGetContestStatus(t *testing.T) {
	assert := assert.New(t)
	submissions, err := api.GetContestStatus(566, nil, 3, 13)
	assert.True(err == nil, "Expected no error while extracting submissions from contest with `contestID` 566, `start` = 3,`end` = 13")
	assert.Equal(len(submissions), 11, "Expected 11 submissions while extracting submissions from contest with `contestID` 566, `start` = 3,`end` = 13, got %d", len(submissions))

	submissions, err = api.GetContestStatus(1179, "I_Love_Tina", 1, 1)
	assert.True(err == nil, "Expected no error while extracting submissions from contes with `contestID` 1179 with handle I_Love_Tina the most recent submission")
	assert.Equal(len(submissions), 1, "Expected one submission while extracting the most recent submission from contest with `contestID` 1179 with handle I_Love_Tina, got %d", len(submissions))
	assert.Equal(submissions[0], api.Submission{
		ID:                  56038919,
		ContestID:           1179,
		CreationTimeSeconds: 1561469037,
		RelativeTimeSeconds: submissions[0].RelativeTimeSeconds,
		Problemv: api.Problem{
			ContestID: 1179,
			Index:     "E",
			Name:      "Alesya and Discrete Math",
			Typev:     "PROGRAMMING",
			Points:    2250.0,
			Rating:    3200,
			Tags:      []string{"divide and conquer", "interactive"},
		},
		Author: api.Party{
			ContestID: 1179,
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
	assert.True(err != nil, "Expected erorr while trying to extract submissions from contest with `contestID` 6669")
	assert.Equal(err.Error(), "contestId: Contest with id 6669 not found", "Expected a different erorr while trying to extract submissions from contest with `contestID` 6669")

	time.Sleep(time.Second)
}

func TestGetPsetProblems(t *testing.T) {
	assert := assert.New(t)
	pStats, problems, err := api.GetPsetProblems([]string{"brute%20force", "math", "implementation", "binary%20search"}, nil)
	assert.True(err == nil, "Expected no error while getting the problems with tags 'brute force','math','implementation','binary search', got none")
	assert.Equal(len(pStats), len(problems), "The problemStatistics and problems must have the same size")
	assert.True(len(problems) > 0, "Expected at least one problem while getting the problems with tags 'brute force','math','implementation','binary search'")

	nProblems := len(problems)

	_, problems, err = api.GetPsetProblems([]string{"brute%20force", "math", "math", "implementation", "binary%20search", "math"}, nil)

	assert.True(err == nil, "Expected no error while getting the problems with tags 'brute force','math','implementation','binary search' with duplicates, got none")
	assert.Equal(nProblems, len(problems), "Expected the same result if we add duplicate tags")

	_, _, err = api.GetPsetProblems([]string{"brute%20force", "math", "math", "implementation", "binary%20search", "math"}, "acmsguru")
	assert.True(err == nil, "Expected no error while getting the problems with tags 'brute force','math','implementation','binary search' with duplicates and psetName 'acmsguru', got none")

	pStats, problems, err = api.GetPsetProblems([]string{"hard%20problem"}, nil)
	assert.True(err == nil, "Expected no error while getting the problems with tag 'hard problem'")
	assert.True(len(problems) == 0 && len(pStats) == 0, "Expected no problem while getting the problems with tag 'hard problem'")

	_, _, err = api.GetPsetProblems(nil, 3.14)
	assert.True(err != nil, "Expected error if `psetName` type is not string")

	time.Sleep(time.Second)
}

func TestGetPsetRecentStatus(t *testing.T) {
	assert := assert.New(t)
	submissions, err := api.GetPsetRecentStatus(69, nil)
	assert.True(err == nil, "Expected no error while extracting 69 most recent submissions")
	assert.Equalf(len(submissions), 69, "Expected 69 most recent submissions, got %d", len(submissions))

	submissions, err = api.GetPsetRecentStatus(69, "acmsguru")
	assert.True(err == nil, "Expected no error while extracting 69 most recent submissions from problemset with `psetName` 'acmsguru'")
	assert.Equalf(len(submissions), 69, "Expected 69 most recent submissions, got %d", len(submissions))

	_, err = api.GetPsetRecentStatus(69, 3.14)
	assert.True(err != nil, "Expected error if `psetName` type is not string")
	assert.Equal(err.Error(), "`psetName` must have type string", "Expected a different error if `psetName` type is not string")

	_, err = api.GetPsetRecentStatus(1001, nil)
	assert.True(err != nil, "Expected error if `countv` is bigger than 1000")
	assert.Equal(err.Error(), "`countv` cannot be more than 1000", "Expected a different error if `countv` is bigger 1000")
}

func TestGetPsetRecentActions(t *testing.T) {
	assert := assert.New(t)
	_, err := api.GetPsetRecentActions(101)
	assert.True(err != nil, "Expected error is `maxCount` is bigger than 100")
	assert.Equal(err.Error(), "`maxCount` cannot be more than 100", "Expected a different error when `maxCount` is bigger than 100")

	time.Sleep(time.Second)

	actions, err := api.GetPsetRecentActions(69)
	assert.True(err == nil, "Expected no error while extracting the 69 most recent actions")
	assert.Equal(len(actions), 69, "Expected 69 recent actions while extracting the 69 most recent actions")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)
	_, err := api.GetUserInfo(nil)
	assert.True(err != nil, "Expected error if `handles` is nil")
	assert.Equal(err.Error(), "There must be at least one handle", "Expected different error if `handles` is nil")

	_, err = api.GetUserInfo(make([]string, 10001))
	assert.True(err != nil, "Expected error if asked for more than 10000 handles")
	assert.Equal(err.Error(), "At most 10000 handles are accepted", "Expected a different error if asked for more than 10000 handles")

	users, err := api.GetUserInfo([]string{"I_Love_Tina"})
	assert.True(err == nil, "Expected no error while extracting user info for handle 'I_Love_Tina'")
	assert.Equal(len(users), 1, "Expected only one user information")
	assert.Equal(users[0], api.User{
		Handle:                  "I_Love_Tina",
		FirstName:               "Gabriel",
		LastName:                "Cojocaru",
		Country:                 "Moldova",
		City:                    "Cahul",
		Organization:            "MIT",
		Rating:                  2390,
		MaxRating:               2390,
		Rank:                    "international master",
		MaxRank:                 "international master",
		LastOnlineTimeSeconds:   users[0].LastOnlineTimeSeconds,
		RegistrationTimeSeconds: users[0].RegistrationTimeSeconds,
		FriendOfCount:           users[0].FriendOfCount,
		TitlePhoto:              "//userpic.codeforces.com/218500/title/5c8d4542cd73311a.jpg",
		Avatar:                  "//userpic.codeforces.com/218500/avatar/190d7b5f7b526d30.jpg",
		Contribution:            users[0].Contribution,
	})

	_, err = api.GetUserInfo([]string{"69"})
	assert.True(err != nil, "Expected error if the user handle is '69'")
	assert.Equal(err.Error(), "handles: User with handle 69 not found", "Expected a different if the user handle is '69'")

	time.Sleep(time.Second)
}

func TestGetBlogEntries(t *testing.T) {
	assert := assert.New(t)
	blogs, err := api.GetBlogEntries("I_Love_Tina")
	assert.True(err == nil, "Expected no error while extracting the blogs from user 'I_Love_Tina'")
	assert.Equal(len(blogs), 7, "User 'I_Love_Tina' has 7 blogs")
	assert.Equal(blogs[6], api.BlogEntry{
		OriginalLocale:          "en",
		AllowViewHistory:        true,
		CreationTimeSeconds:     blogs[6].CreationTimeSeconds,
		Rating:                  blogs[6].Rating,
		AuthorHandle:            "I_Love_Tina",
		ModificationTimeSeconds: blogs[6].ModificationTimeSeconds,
		ID:                      45896,
		Title:                   "\u003cp\u003eCodeforces Round #361 (Div. 2)\u003c/p\u003e",
		Locale:                  "en",
		Tags:                    []string{},
	})

	_, err = api.GetBlogEntries("69")

	assert.True(err != nil, "Expected error while extracting the blogs from user with handle '69'")
	assert.Equal(err.Error(), "handle: Field should contain between 3 and 24 characters, inclusive", "Expected a different error while extracting the blog from users with handle '69'")

	_, err = api.GetBlogEntries("666")
	assert.True(err != nil, "Expected error while extracting the blogs from user with handle '666'")
	assert.Equal(err.Error(), "handle: You are not allowed to read that blog", "Expected a a different error while extracting the blog from users with handle '666'")
}

func TestGetRatedList(t *testing.T) {
	assert := assert.New(t)
	tmp, _ := api.GetUserInfo([]string{"tourist"})
	tourist := tmp[0]

	users, err := api.GetRatedList(nil)
	var touristRank int = 0
	for i, user := range users {
		if user == tourist {
			touristRank = i
		}
	}
	nUsers := len(users)
	assert.True(err == nil, "Expected no error while extracting the rated list")
	assert.Truef(1 <= touristRank && touristRank <= 20, "Error tourist's rank is not in range [1,20], got %d", touristRank)

	time.Sleep(time.Second)

	tmp, _ = api.GetUserInfo([]string{"I_Love_Tina"})
	reality := tmp[0]

	users, err = api.GetRatedList(true)
	assert.True(err == nil, "Expected no error while extracting the rated list with only active users")
	for _, user := range users {
		assert.True(user != reality, "User I_Love_Tina is not active")
	}
	assert.Truef(nUsers > len(users), "The number of users while showing only showing active only must be smaller (which is %d) than total number of users (which is %d)", len(users), nUsers)

	_, err = api.GetRatedList(3.14)
	assert.True(err != nil, "Expected error when `activeOnly` type is not bool")
	assert.Equal(err.Error(), "`activeOnly` must have type bool", "Expected a different error when `activeOnly` type is not bool")
}

func TestGetUserRatings(t *testing.T) {
	assert := assert.New(t)

	changes, err := api.GetUserRatings("I_Love_Tina")
	assert.True(err == nil, "Expected no error while extracting the rating changes from user 'I_Love_Tina'")
	assert.Equal(len(changes), 112, "User 'I_Love_Tina' has 112 official contests")
	assert.Equal(changes[1], api.RatingChange{
		ContestID:               447,
		ContestName:             "Codeforces Round #FF (Div. 2)",
		Handle:                  "I_Love_Tina",
		Rank:                    1518,
		RatingUpdateTimeSeconds: 1405263600,
		OldRating:               1439,
		NewRating:               1350,
	})

	_, err = api.GetUserRatings("000")
	assert.True(err != nil, "Expected error while extracting the rating changes from user with handle '000'")
	assert.Equal(err.Error(), "handle: User 000 not found", "Expected a differnt while extracting the rating changes from user with handle '000'")

	time.Sleep(time.Second)
}

func TestGetUserStatus(t *testing.T) {
	assert := assert.New(t)

	submissions, err := api.GetUserStatus("I_Love_Tina", 3, 5)
	assert.True(err == nil, "Expected no error while extracting the submissions with indices in interval [3,5] from user 'I_Love_Tina'")
	assert.Equal(len(submissions), 5-3+1, "Expected 3 submissions while while extracting the submissions with indices in interval [3,5] from user 'I_Love_Tina'")

	submissions, err = api.GetUserStatus("I_Love_Tina", nil, nil)
	assert.True(err == nil, "Expected no error while extracting all the submissions from user 'I_Love_Tina'")
	assert.Equal(submissions[len(submissions)-1], api.Submission{
		ID:                  7009763,
		ContestID:           443,
		CreationTimeSeconds: 1404468318,
		RelativeTimeSeconds: 318,
		Problemv: api.Problem{
			ContestID: 443,
			Index:     "A",
			Name:      "Anton and Letters",
			Typev:     "PROGRAMMING",
			Points:    500.0,
			Rating:    800,
			Tags:      []string{"constructive algorithms", "implementation"},
		},
		Author: api.Party{
			ContestID: 443,
			Members: []api.Member{
				{
					Handle: "I_Love_Tina",
				},
			},
			ParticipantType:  "VIRTUAL",
			Ghost:            false,
			StartTimeSeconds: 1404468000,
		},
		ProgrammingLanguage: "GNU C++",
		Verdict:             "OK",
		Testset:             "TESTS",
		PassedTestCount:     27,
		TimeConsumedMillis:  15,
		MemoryConsumedBytes: 0,
	})

	_, err = api.GetUserStatus("I_Love_Tina", nil, 5)
	assert.True(err != nil, "Expected error while extracting the submissions from user 'I_Love_Tina' with start = nil and end = 5")
	assert.Equal(err.Error(), "If `end` is not nil then from should not be nil as well", "Expected a different error while extracting the submissions from user 'I_Love_Tina' with `start` = nil and `end` = 5")

	_, err = api.GetUserStatus("I_Love_Tina", 3.5, nil)
	assert.True(err != nil, "Expected error if `start` type is not int")

	_, err = api.GetUserStatus("I_Love_Tina", 3, "")
	assert.True(err != nil, "Expected error if `end` type is not int")

	time.Sleep(time.Second)

	_, err = api.GetUserStatus("000", nil, nil)
	assert.True(err != nil, "Expected error while extracting the submissions from user '000'")
	assert.Equal(err.Error(), "handle: User with handle 000 not found", "Expected a different error while extracting submissions from user '000'")
}
