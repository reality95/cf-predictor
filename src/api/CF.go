package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

/*
	All the methods are described in the following link: https://codeforces.com/apiHelp/methods
*/

//GetComments ... getting the list of comments from blog with blog ID `blogID`
func GetComments(blogID int) ([]Comment, error) {
	resp, err := http.Get("https://codeforces.com/api/blogEntry.comments?blogEntryId=" + strconv.Itoa(blogID))
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestComments
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetBlog ... getting the BlogEntry from the blog with blog ID `blogID`
func GetBlog(blogID int) (BlogEntry, error) {
	resp, err := http.Get("https://codeforces.com/api/blogEntry.view?blogEntryId=" + strconv.Itoa(blogID))
	if err != nil {
		return BlogEntry{}, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return BlogEntry{}, err
	}
	var data RequestBlog
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return BlogEntry{}, errors.New(data.Commentv)
	}
	return BlogEntry{}, errors.New("Unknown Error")
}

//GetHacks ... getting the list of hacks from the contest with contest ID `contestID`
func GetHacks(contestID int) ([]Hack, error) {
	resp, err := http.Get("https://codeforces.com/api/contest.hacks?contestId=" + strconv.Itoa(contestID))
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestHacks
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetContests ... getting the list of contests
/*
	Optional arguments:
		(*) gymIncluded - if it is not nil then will include the contests from the gym if `gymIncluded` is true, otherwise false
*/
func GetContests(gymIncluded interface{}) ([]Contest, error) {
	addr := "https://codeforces.com/api/contest.list"
	if gymIncluded != nil {
		if reflect.TypeOf(gymIncluded).Kind() == reflect.Bool {
			return nil, errors.New("`gymIncluded` must have type bool")
		}
		addr += "?gym=" + strconv.FormatBool(gymIncluded.(bool))
	}
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestContests
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetRatingChanges ... getting the list of all rating changes in the contest with contest ID `contestID`
func GetRatingChanges(contestID int) ([]RatingChange, error) {
	resp, err := http.Get("https://codeforces.com/api/contest.ratingChanges?contestId=" + strconv.Itoa(contestID))
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestRatingChange
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetContestStandings ... getting the contest standings from the contest with contest ID `contestID`
/*	Optional arguments in the following order:
	(*) handles - if it is not nil, the contest standings will be restricted only to these handles, at most 10000 handles
	(*) start - if it is not nil, the contest standings will include only participants starting at position `start`
	(*) end - if it is not nil, start must not be nil as well, the contest standings will include only participants between positions `start` and `end`
	(*) room - if it is not nil, the contest standings will include only participants from room `room`
	(*) showUnofficial - if it is not nil, the answer will include unofficial standings if `showUnofficial` is false and not included if it's true,
						by default it is false
*/
func GetContestStandings(contestID int, handles []string, start interface{}, end interface{}, room interface{}, showUnofficial interface{}) (ContestStandings, error) {
	addr := "https://codeforces.com/api/contest.standings?contestId=" + strconv.Itoa(contestID)
	if start != nil {
		if reflect.TypeOf(start).Kind() != reflect.Int {
			return ContestStandings{}, errors.New("`start` argument must be either nil or int")
		}
		addr += "&from=" + strconv.Itoa(start.(int))
	}
	if end != nil {
		if start == nil {
			return ContestStandings{}, errors.New("If `end` is not nil, `start` must not be nil as well")
		}
		if reflect.TypeOf(end).Kind() != reflect.Int {
			return ContestStandings{}, errors.New("`end` argument must be either nil or int")
		}
		addr += "&count=" + strconv.Itoa(end.(int)-start.(int)+1)
	}
	if handles != nil {
		if len(handles) > 10000 {
			return ContestStandings{}, errors.New("Expected at most 10000 handles")
		}
		addr += "&handles="
		for i, handle := range handles {
			addr += handle
			if i+1 < len(handles) {
				addr += ";"
			}
		}
	}
	if room != nil {
		if reflect.TypeOf(room).Kind() != reflect.Int {
			return ContestStandings{}, errors.New("`room` argument must be either nil or int")
		}
		addr += "&room=" + strconv.Itoa(room.(int))
	}
	if showUnofficial != nil {
		if reflect.TypeOf(showUnofficial).Kind() != reflect.Bool {
			return ContestStandings{}, errors.New("`showUnofficial` argument must be either nil or bool")
		}
		addr += "&showUnofficial=" + strconv.FormatBool(showUnofficial.(bool))
	}
	resp, err := http.Get(addr)
	if err != nil {
		return ContestStandings{}, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ContestStandings{}, err
	}
	var data RequestContestStandings
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return ContestStandings{}, errors.New(data.Commentv)
	}
	return ContestStandings{}, errors.New("Unknown Error")
}

//GetContestStatus ... getting the contest status from the contest with contest ID `contestID`
/*		Optional arguments in the following order:
		(*) handles - if it is not nil, the contest standings will be restricted only to submissions submitted by participant with handle `handle`
		(*) start - if it is not nil, the contest standings will include only participants starting at position `start`
		(*) end - if it is not nil, `start` must not be nil as well, the contest standings will include only participants between positions `start` and `end`
*/
func GetContestStatus(contestID int, handle interface{}, start interface{}, end interface{}) ([]Submission, error) {
	addr := "https://codeforces.com/api/contest.status?contestId=" + strconv.Itoa(contestID)
	if handle != nil {
		if reflect.TypeOf(handle).Kind() != reflect.String {
			return nil, errors.New("`handle` must have type string")
		}
		addr += "&handle=" + handle.(string)
	}
	if start != nil {
		if reflect.TypeOf(start).Kind() != reflect.Int {
			return nil, errors.New("`start` must have type int")
		}
		addr += "&from=" + strconv.Itoa(start.(int))
	}
	if end != nil {
		if start == nil {
			return nil, errors.New("If `end` is not nil, `start` must not be nil as well")
		}
		if reflect.TypeOf(end).Kind() != reflect.Int {
			return nil, errors.New("`end` must have type int")
		}
		addr += "&count=" + strconv.Itoa(end.(int)-start.(int)+1)
	}
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestContestStatus
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetPsetProblems ... getting the problem statistics from the problemset
/*
	Optional arguments:
		(*) tags - if it is not nil, will return only the problem statistics for problems whose tags include subset `tags`
		(*) psetName - if it is not nil, will return the problem statistics for problems from problemset with name `psetName`
*/
func GetPsetProblems(tags []string, psetName interface{}) ([]ProblemStatistics, []Problem, error) {
	addr := "https://codeforces.com/api/problemset.problems"
	if tags != nil {
		addr += "?tags="
		for i, tag := range tags {
			addr += tag
			if i+1 < len(tags) {
				addr += ";"
			}
		}
	}
	if psetName != nil {
		if reflect.TypeOf(psetName).Kind() != reflect.String {
			return nil, nil, errors.New("`psetName` must have type string")
		}
		if tags != nil {
			addr += "&"
		} else {
			addr += "?"
		}
		addr += "problemsetName=" + psetName.(string)
	}
	resp, err := http.Get(addr)
	if err != nil {
		return nil, nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	var data RequestPsetProblems
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result.PStats, data.Result.Problems, nil
	}
	if data.Commentv != "" {
		return nil, nil, errors.New(data.Commentv)
	}
	return nil, nil, errors.New("Unknown Error")
}

//GetPsetRecentStatus ... get `countv` recent submissions (1 <= `countv` <= 1000)
/*
	Optional arguments:
		(*) psetName - if it is not nil, will return the problem statistics for problems with name `psetName`
*/
func GetPsetRecentStatus(countv int, psetName interface{}) ([]Submission, error) {
	if countv > 1000 {
		return nil, errors.New("`countv` cannot be more than 1000")
	}
	addr := "https://codeforces.com/api/problemset.recentStatus?count=" + strconv.Itoa(countv)
	if psetName != nil {
		if reflect.TypeOf(psetName).Kind() != reflect.String {
			return nil, errors.New("`psetName` must have type string")
		}
		addr += "&problemsetName=" + psetName.(string)
	}
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestRecentStatus
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetPsetRecentActions ... get at most `maxCount` recent actions (1 <= `maxCount` <= 100)
func GetPsetRecentActions(maxCount int) ([]RecentAction, error) {
	if maxCount > 100 {
		return nil, errors.New("`maxCount` cannot be more than 100")
	}
	addr := "https://codeforces.com/api/recentActions?maxCount=" + strconv.Itoa(maxCount)
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestRecentActions
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetBlogEntries ... get the list of blog entries from user with handle `handle`
func GetBlogEntries(handle string) ([]BlogEntry, error) {
	addr := "https://codeforces.com/api/user.blogEntries?handle=" + handle
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestBlogEntries
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetUserInfo ... get the user info for all users in `handles`, at most 10000 handles
func GetUserInfo(handles []string) ([]User, error) {
	if handles == nil {
		return nil, errors.New("There must be at least one handle")
	}
	if len(handles) > 10000 {
		return nil, errors.New("At most 10000 handles are accepted")
	}
	addr := "https://codeforces.com/api/user.info?handles="
	for i, handle := range handles {
		addr += handle
		if i+1 < len(handles) {
			addr += ";"
		}
	}
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestUserInfo
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetRatedList ... get the list of all users in the increasing order of ranks
/*
	Optional arguments:
		(*) activeOnly - if it is not nil, will include only active users in the last 6 months if `activeOnly` is true
*/
func GetRatedList(activeOnly interface{}) ([]User, error) {
	addr := "https://codeforces.com/api/user.ratedList"
	if activeOnly != nil {
		if reflect.TypeOf(activeOnly).Kind() != reflect.Bool {
			return nil, errors.New("`activeOnly` must have type bool")
		}
		addr += "?activeOnly=" + strconv.FormatBool(activeOnly.(bool))
	}
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestRatedList
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetUserRatings ... get the list of all rating changes for user with handle `handle`
func GetUserRatings(handle string) ([]RatingChange, error) {
	addr := "https://codeforces.com/api/user.rating?handle=" + handle
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestUserRating
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}

//GetUserStatus ... get the list of submissions from user with handle `handle`
/*
	Optional arguments:
		(*) start - if it is not nil, will include only the submissions starting with index `start`
		(*) end - if it is not nil, `start` must not be nil as well, will include only the submissions with indices between `start` and `end`
*/
func GetUserStatus(handle string, start interface{}, end interface{}) ([]Submission, error) {
	addr := "https://codeforces.com/api/user.status?handle=" + handle
	if start != nil {
		if reflect.TypeOf(start).Kind() != reflect.Int {
			return nil, errors.New("`start` must have type int")
		}
		addr += "&from=" + strconv.Itoa(start.(int))
	}
	if end != nil {
		if start == nil {
			return nil, errors.New("If `end` is not nil then from should not be nil as well")
		}
		if reflect.TypeOf(end).Kind() != reflect.Int {
			return nil, errors.New("`end` must have type int")
		}
		addr += "&count=" + strconv.Itoa(end.(int)-start.(int)+1)
	}
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data RequestUserStatus
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	}
	if data.Commentv != "" {
		return nil, errors.New(data.Commentv)
	}
	return nil, errors.New("Unknown Error")
}
