package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

func GetComments(blogId int) ([]Comment, error) {
	resp, err := http.Get("https://codeforces.com/api/blogEntry.comments?blogEntryId=" + strconv.Itoa(blogId))
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func GetBlog(blogId int) (BlogEntry, error) {
	resp, err := http.Get("https://codeforces.com/api/blogEntry.view?blogEntryId=" + strconv.Itoa(blogId))
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
	} else {
		if data.Commentv != "" {
			return BlogEntry{}, errors.New(data.Commentv)
		} else {
			return BlogEntry{}, errors.New("Unknown Error")
		}
	}
}

func GetHacks(contestId int) ([]Hack, error) {
	resp, err := http.Get("https://codeforces.com/api/contest.hacks?contestId=" + strconv.Itoa(contestId))
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func GetContests(gymIncluded interface{}) ([]Contest, error) {
	addr := "https://codeforces.com/api/contest.list"
	if gymIncluded != nil {
		if reflect.TypeOf(gymIncluded).Kind() == reflect.Bool {
			return nil, errors.New("gymIncluded must have type bool")
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func GetRatingChanges(contestId int) ([]RatingChange, error) {
	resp, err := http.Get("https://codeforces.com/api/contest.ratingChanges?contestId=" + strconv.Itoa(contestId))
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func GetContestStandings(contestId int, handles []string, start interface{}, end interface{}, room interface{}, showUnofficial interface{}) (ContestStandings, error) {
	addr := "https://codeforces.com/api/contest.standings?contestId=" + strconv.Itoa(contestId)
	if start != nil {
		if reflect.TypeOf(start).Kind() != reflect.Int {
			return ContestStandings{}, errors.New("start argument must be either nil or int")
		}
		addr += "&from=" + strconv.Itoa(start.(int))
	}
	if end != nil {
		if start == nil {
			return ContestStandings{}, errors.New("If end is not nil, start must not be nil as well")
		}
		if reflect.TypeOf(end).Kind() != reflect.Int {
			return ContestStandings{}, errors.New("end argument must be either nil or int")
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
			return ContestStandings{}, errors.New("room argument must be either nil or int")
		}
		addr += "&room=" + strconv.Itoa(room.(int))
	}
	if showUnofficial != nil {
		if reflect.TypeOf(showUnofficial).Kind() != reflect.Bool {
			return ContestStandings{}, errors.New("showUnofficial argument must be either nil or bool")
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
	} else {
		if data.Commentv != "" {
			return ContestStandings{}, errors.New(data.Commentv)
		} else {
			return ContestStandings{}, errors.New("Unknown Error")
		}
	}
}

func GetContestStatus(contestId int, handle interface{}, start interface{}, end interface{}) ([]Submission, error) {
	addr := " https://codeforces.com/api/contest.status?contestId=" + strconv.Itoa(contestId)
	if handle != nil {
		if reflect.TypeOf(handle).Kind() != reflect.String {
			return nil, errors.New("handle must have type string")
		}
		addr += "&handle=" + handle.(string)
	}
	if start != nil {
		if reflect.TypeOf(start).Kind() != reflect.Int {
			return nil, errors.New("start must have type int")
		}
		addr += "&from=" + strconv.Itoa(start.(int))
	}
	if end != nil {
		if start == nil {
			return nil, errors.New("If end is not nil, start must not be nil as well")
		}
		if reflect.TypeOf(end).Kind() != reflect.Int {
			return nil, errors.New("end must have type int")
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func GetPsetProblems(tags []string, psetName interface{}) (ProblemStatistics, error) {
	addr := "https://codeforces.com/api/problemset.problems"
	if tags != nil {
		addr += "&tags="
		for i, tag := range tags {
			addr += tag
			if i+1 < len(tags) {
				addr += ";"
			}
		}
	}
	if psetName != nil {
		if reflect.TypeOf(psetName).Kind() != reflect.String {
			return ProblemStatistics{}, errors.New("problemsetName must have type string")
		}
		addr += "&problemsetName=" + psetName.(string)
	}
	resp, err := http.Get(addr)
	if err != nil {
		return ProblemStatistics{}, err
	}
	var plainText []byte
	plainText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ProblemStatistics{}, err
	}
	var data RequestPsetProblems
	json.Unmarshal(plainText, &data)
	if data.Status == "OK" {
		return data.Result, nil
	} else {
		if data.Commentv != "" {
			return ProblemStatistics{}, errors.New(data.Commentv)
		} else {
			return ProblemStatistics{}, errors.New("Unknown Error")
		}
	}
}

func GetPsetRecentStatus(count_ int, psetName interface{}) ([]Submission, error) {
	addr := "https://codeforces.com/api/problemset.recentStatus?count=" + strconv.Itoa(count_)
	if psetName != nil {
		if reflect.TypeOf(psetName).Kind() != reflect.String {
			return nil, errors.New("problemsetName must have type string")
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func GetPsetRecentActions(maxCount int) ([]RecentAction, error) {
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func GetRatedList(activeOnly interface{}) ([]User, error) {
	addr := "https://codeforces.com/api/user.ratedList"
	if activeOnly != nil {
		if reflect.TypeOf(activeOnly).Kind() != reflect.Bool {
			return nil, errors.New("activeOnly must have type bool")
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}

func getUserStatus(handle string, start interface{}, end interface{}) ([]Submission, error) {
	addr := "https://codeforces.com/api/user.status?handle=" + handle
	if start != nil {
		if reflect.TypeOf(start).Kind() != reflect.Int {
			return nil, errors.New("start must have type int")
		}
		addr += "&from=" + strconv.Itoa(start.(int))
	}
	if end != nil {
		if start == nil {
			return nil, errors.New("If count is not nil then from should not be nil as well")
		}
		if reflect.TypeOf(end).Kind() != reflect.Int {
			return nil, errors.New("end must have type int")
		}
		addr += "&count" + strconv.Itoa(end.(int)-start.(int)+1)
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
	} else {
		if data.Commentv != "" {
			return nil, errors.New(data.Commentv)
		} else {
			return nil, errors.New("Unknown Error")
		}
	}
}
