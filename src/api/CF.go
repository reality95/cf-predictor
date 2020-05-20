package api

import ("net/http"
		"encoding/json"
		"strconv"
		"errors"
		"io/ioutil"
		"reflect"
)

func GetComments(blogId int) ([]Comment,error) {
	resp,err := http.Get("https://codeforces.com/api/blogEntry.comments?blogEntryId=" + strconv.Itoa(blogId))
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestComments
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}

func GetBlog(blogId int) (BlogEntry,error) {
	resp,err := http.Get("https://codeforces.com/api/blogEntry.view?blogEntryId=" + strconv.Itoa(blogId))
	if err != nil {
		return BlogEntry {},err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return BlogEntry {},err
	}
	var data RequestBlog
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return BlogEntry {},data.Comment_
		} else {
			return BlogEntry {},errors.New("Unkown Error")
		}
	}
}

func GetHacks(contestId int) ([]Hack,error) {
	resp,err := http.Get("https://codeforces.com/api/contest.hacks?contestId=" + strconv.Itoa(contestId))
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestHacks
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}
//strconv.FormatBool(v)
func GetContests(gymIncluded... bool) ([]Contest,error) {
	addr := "https://codeforces.com/api/contest.list"
	if gymIncluded != nil {
		if len(gymIncluded) != 1 {
			return nil,errors.New("gymIncluded must have only one argument")
		}
		addr += "?gym=" + strconv.FormatBool(gymIncluded[0].(bool))
	}
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestContests
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}

func GetRatingChanges(contestId int) ([]RatingChange,error) {
	resp,err := http.Get("https://codeforces.com/api/contest.ratingChanges?contestId=" + strconv.Itoa(contestId))
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestRatingChange
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}

func GetContestStandings(contestId int,handles []string,args...interface{})(ContestStandings,error) {
	if len(args) > 4 {
		return ContestStandings,errors.New("Expected at most 6 arguments")
	}
	addr := "https://codeforces.com/api/contest.standings?contestId=" + strconv.Itoa(contestId)
	if args != nil && args[0] != nil {
		if reflect.TypeOf(args[0]).Kind() != reflect.Int {
			return ContestStandings{},errors.New("start argument must be either nil or int")
		}
		addr += "&from=" + strconv.Itoa(args[0].(int))
	}
	if len(args) > 1 && args[1] != nil {
		if args[0] == nil {
			return nil,errors.New("If end is not nil, start must not be nil as well")
		}
		if reflect.TypeOf(args[1]).Kind() != reflect.Int {
			return ContestStandings{},errors.New("end argument must be either nil or int")
		}
		addr += "&count" + strconv.Itoa(args[1].(int) - args[0].(int) + 1)
	}
	if handles != nil {
		addr += "&handles="
		for i,handle := range handles {
			addr += handle
			if i+1 < len(hanldes) {
				addr += ";"
			}
		}
	}
	if len(args) > 2 && args[2] != nil {
		if reflect.TypeOf(args[2]).Kind() != reflect.Int {
			return ContestStandings{},errors.New("room argument must be either nil or int")
		}
		addr += "&room=" + strconv.Itoa(args[2].(int))
	}
	if len(args) > 3 && args[3] != nil {
		if reflect.TypeOf(args[3]).Kind() != reflect.Bool {
			return ContestStandings{},errors.New("showUnofficial argument must be either nil or bool")
		}		
		addr += "&showUnofficial=" + strconv.FormatBool(args[3].(bool))
	}
	resp,err := http.Get(addr)
	if err != nil {
		return ContestStandings{},err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ContestStandings{},err
	}
	var data RequestContestStandings
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return ContestStandings{},data.Comment_
		} else {
			return ContestStandings{},errors.New("Unkown Error")
		}
	}
}

func GetContestStatus(contestId int,args...interface{}) ([]Submission,error) {
	if len(args) > 3 {
		return nil,erros.New("Expected at most 4 arguments")
	}
	addr := " https://codeforces.com/api/contest.status?contestId=" + strconv.Itoa(contestId)
	if args != nil && args[0] != nil {
		if reflect.TypeOf(args[0]).Kind() != reflect.String {
			return nil,errors.New("handle must have type string")
		}
		addr += "&handle=" + args[0].(string)
	}
	if len(args) > 1 && args[1] != nil {
		if reflect.TypeOf(args[1]).Kind() != reflect.Int {
			return nil,errors.New("start must have type int")
		}
		addr += "&from=" + strconv.Itoa(args[1].(int))
	}
	if len(args) > 2 && args[2] != nil {
		if args[1] == nil {
			return nil,errors.New("If end is not nil, start must not be nil as well")
		}
		if reflect.TypeOf(args[2]).Kind() != reflect.Int {
			return nil,errors.New("end must have type int")
		}
		addr += "&count" + strconv.Itoa(args[2].(int) - args[1].(int) + 1)
	}
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestContestStatus
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}

func GetPsetProblems(tags []string,psetName... string) (ProblemStatistics,error) {
	addr := "https://codeforces.com/api/problemset.problems"
	if tags != nil {
		addr += "&tags="
		for i,tag := range tags {
			addr += tag
			if i+1 < len(tags) {
				addr += ";"
			}
		}
	}
	if psetName != nil {
		if len(psetName) != 1 {
			return ProblemStatistics{},errors.New("There must be only one problemsetName")
		}
		addr += "&problemsetName=" + psetName[0].(string)
	}
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestPsetProblems
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}

func GetPsetRecentStatus(count_ int,psetName... string) ([]Submission,error) {
	addr := "https://codeforces.com/api/problemset.recentStatus?count=" + strconv.Itoa(count_)
	if psetName != nil {
		if len(psetName) != 1 {
			return nil,errors.New("problemsetName must have only one argument")
		}
		addr += "&problemsetName=" + psetName.(string)
	}
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestRecentStatus
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}

func GetPsetRecentActions(maxCount int) ([]RecentAction,error) {
	addr := "https://codeforces.com/api/recentActions?maxCount=" + strconv.Itoa(maxCount)
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestRecentActions
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}	
}

func GetBlogEntries(handle string) ([]BlogEntry,error) {
	addr := "https://codeforces.com/api/user.blogEntries?handle=" + handle
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestBlogEntries
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}	
}

func GetUserInfo(handles []string) ([]User,error) {
	if handles == nil {
		return nil,errors.New("There must be at least one handle")
	}
	if len(handles) > 10000 {
		return nil,erros.New("At most 10000 handles are accepted")
	}
	addr := "https://codeforces.com/api/user.info?handles="
	for i,handle := range handles {
		addr += handle
		if i+1 < len(handles) {
			addr += ";"
		}
	}
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestUserInfo
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}
}

func GetRatedList(activeOnly... bool) ([]User,error) {
	addr := "https://codeforces.com/api/user.ratedList"
	if activeOnly != nil {
		if len(activeOnly) != 1 {
			return nil,errors.New("activeOnly must contain only one argument")
		}
		addr += "?activeOnly=" + strconv.FormatBool(activeOnly[0])
	}
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestRatedList
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}	
}

func GetUserRatings(handle string) ([]RatingChange,error) {
	addr := "https://codeforces.com/api/user.rating?handle=" + handle
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestUserRating
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}	
}

func getUserStatus(handle string,args...int) ([]Submission,error) {
	if len(args) > 2 {
		return nil,errors.New("Expected at most 3 arguments")
	}
	addr := "https://codeforces.com/api/user.status?handle=" + handle
	if args != nil && args[0] != nil {
		addr += "&from=" + strconv.Itoa(args[0])
	}
	if len(args) > 1 && args[1] != nil {
		if args[0] == nil {
			return nil,errors.New("If count is not nil then from should not be nil as well")
		}
		addr += "&count" + strconv.Itoa(args[1] - args[0] + 1)
	}
	resp,err := http.Get(addr)
	if err != nil {
		return nil,err
	}
	var plainText []byte
	plainText,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var data RequestUserStatus
	json.Unmarshal(plainText,data)
	if data.Status == "OK" {
		return data.Result,nil
	} else {
		if data.Comment_ != nil {
			return nil,data.Comment_
		} else {
			return nil,errors.New("Unkown Error")
		}
	}	
}