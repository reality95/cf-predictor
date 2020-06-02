package api

import "strconv"

//Comment ... Comment object described here: https://codeforces.com/apiHelp/objects
type Comment struct {
	ID                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	Text                string `json:"text"`
	ParentCommentID     int    `json:"parentCommentId"`
	Rating              int    `json:"rating"`
}

//BlogEntry ... BlogEntry object described here: https://codeforces.com/apiHelp/objects
type BlogEntry struct {
	ID                      int      `json:"id"`
	OriginalLocale          string   `json:"originalLocale"`
	CreationTimeSeconds     int      `json:"creationTimeSeconds"`
	AuthorHandle            string   `json:"authorHandle"`
	Title                   string   `json:"title"`
	Content                 string   `json:"content"`
	Locale                  string   `json:"locale"`
	ModificationTimeSeconds int      `json:"modificationTimeSeconds"`
	AllowViewHistory        bool     `json:"allowViewHistory"`
	Tags                    []string `json:"tags"`
	Rating                  int      `json:"rating"`
}

//User ... User object described here: https://codeforces.com/apiHelp/objects
type User struct {
	Handle                  string `json:"handle"`
	Email                   string `json:"email"`
	VkID                    string `json:"vkId"`
	OpenID                  string `json:"openId"`
	FirstName               string `json:"firstName"`
	LastName                string `json:"lastName"`
	Country                 string `json:"country"`
	City                    string `json:"city"`
	Organization            string `json:"organization"`
	Contribution            int    `json:"contribution"`
	Rank                    string `json:"rank"`
	Rating                  int    `json:"rating"`
	MaxRank                 string `json:"maxRank"`
	MaxRating               int    `json:"maxRating"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"`
	FriendOfCount           int    `json:"friendOfCount"`
	Avatar                  string `json:"avatar"`
	TitlePhoto              string `json:"titlePhoto"`
}

//RecentAction ... RecentAction object described here: https://codeforces.com/apiHelp/objects
type RecentAction struct {
	TimeSeconds int       `json:"timeSeconds"`
	BlogEntryv  BlogEntry `json:"blogEntry"`
	Commentv    Comment   `json:"comment"`
}

//RatingChange ... RatingChange object described here: https://codeforces.com/apiHelp/objects
type RatingChange struct {
	ContestID               int    `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	Rank                    int    `json:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating"`
	NewRating               int    `json:"newRating"`
}

//Contest ... Contest object described here: https://codeforces.com/apiHelp/objects
type Contest struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Typev               string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int    `json:"durationSeconds"`
	StartTimeSeconds    int    `json:"startTimeSeconds"`
	RelativeTimeSeconds int    `json:"relativeTimeSeconds"`
	PreparedBy          string `json:"preparedBy"`
	WebsiteUrl          string `json:"websiteUrl"`
	Description         string `json:"description"`
	Difficulty          int    `json:"difficulty"`
	Kind                string `json:"kind"`
	IcpcRegion          string `json:"icpcRegion"`
	Country             string `json:"country"`
	City                string `json:"city"`
	Season              string `json:"season"`
}

//Member ... Member object described here: https://codeforces.com/apiHelp/objects
type Member struct {
	Handle string `json:"handle"`
}

//Party ... Party object described here: https://codeforces.com/apiHelp/objects
type Party struct {
	ContestID        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	TeamID           int      `json:"teamId"`
	TeamName         string   `json:"teamName"`
	Ghost            bool     `json:"ghost"`
	Room             int      `json:"room"`
	StartTimeSeconds int      `json:"startTimeSeconds"`
}

//Len ... the number of members in the party
func (p Party) Len() int {
	return len(p.Members)
}

//Problem ... Problem object described here: https://codeforces.com/apiHelp/objects
type Problem struct {
	ContestID int      `json:"contestId"`
	PsetName  string   `json:"problemsetName"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Typev     string   `json:"type"`
	Points    float64  `json:"points"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}

//Link ... the link of the problem if it is not from acm sgu problemset
func (p Problem) Link() string {
	return "https://codeforces.com/contest/" + strconv.Itoa(p.ContestID) + "/problem/" + p.Index
}

//Hash ... a string hash of the problem object
func (p Problem) Hash() string {
	return strconv.Itoa(p.ContestID) + "$" + p.Index
}

//LessProblem ... a comparative function for Problem object
func LessProblem(a, b Problem) bool {
	if a.ContestID != b.ContestID {
		return a.ContestID < b.ContestID
	}
	return a.Index < b.Index
}

//EqProblem ... compares if two problems are equal
func EqProblem(a, b Problem) bool {
	return a.ContestID == b.ContestID && a.Index == b.Index && a.Name == b.Name
}

//ProblemStatistics ... ProblemStatistics object described here: https://codeforces.com/apiHelp/objects
type ProblemStatistics struct {
	ContestID   int    `json:"contestId"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}

//Hash ... a string hash of ProblemStatistics object, if compared to hash of a Problem
//will be equal if it represents the problem statistics of that problem
func (pStat ProblemStatistics) Hash() string {
	return strconv.Itoa(pStat.ContestID) + "$" + pStat.Index
}

//PsetProblems ... contains the object from problemset.problems described here:
//												https://codeforces.com/apiHelp/methods
type PsetProblems struct {
	Problems []Problem           `json:"problems"`
	PStats   []ProblemStatistics `json:"problemStatistics"`
}

//Submission ... Submission object described here: https://codeforces.com/apiHelp/objects
type Submission struct {
	ID                  int     `json:"id"`
	ContestID           int     `json:"contestId"`
	CreationTimeSeconds int     `json:"creationTimeSeconds"`
	RelativeTimeSeconds int     `json:"relativeTimeSeconds"`
	Problemv            Problem `json:"problem"`
	Author              Party   `json:"author"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict"`
	Testset             string  `json:"testset"`
	PassedTestCount     int     `json:"passedTestCount"`
	TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
}

//JProtocol ... an object contained into struct hack described here: https://codeforces.com/apiHelp/objects
type JProtocol struct {
	Manual   string `json:"manual"`
	Protocol string `json:"protocol"`
	Verdict  string `json:"verdict"`
}

//Hack ... Hack object described here: https://codeforces.com/apiHelp/objects
type Hack struct {
	ID                  int       `json:"id"`
	CreationTimeSeconds int       `json:"creationTimeSeconds"`
	Hacker              Party     `json:"hacker"`
	Defender            Party     `json:"defender"`
	Verdict             string    `json:"verdict"`
	Problemv            Problem   `json:"problem"`
	Test                string    `json:"test"`
	JudgeProtocol       JProtocol `json:"judgeProtocol"`
}

//RanklistRow ... RanklistRow object described here: https://codeforces.com/apiHelp/objects
type RanklistRow struct {
	Partyv                    Party           `json:"party"`
	Rank                      int             `json:"rank"`
	Points                    float64         `json:"points"`
	Penalty                   int             `json:"penalty"`
	SuccessfulHackCount       int             `json:"successfulHackCount"`
	UnsuccessfulHackCount     int             `json:"unsuccessfulHackCount"`
	ProblemResults            []ProblemResult `json:"problemResults"`
	LastSubmissionTimeSeconds int             `json:"lastSubmissionTimeSeconds"`
}

//ProblemResult ... ProblemResult object described here: https://codeforces.com/apiHelp/objects
type ProblemResult struct {
	Points                    float64 `json:"points"`
	Penalty                   int     `json:"penalty"`
	RejectedAttemptCount      int     `json:"rejectedAttemptCount"`
	Typev                     string  `json:"type"`
	BestSubmissionTimeSeconds int     `json:"bestSubmissionTimeSeconds"`
}

//ContestStandings ... contains the object from contest.standings described here:
//													https://codeforces.com/apiHelp/methods
type ContestStandings struct {
	Contestv Contest       `json:"contest"`
	Problems []Problem     `json:"problems"`
	Rows     []RanklistRow `json:"rows"`
}

//ProblemsetProblems ... contains the object with the problems and problem statistics
//from problemset.problems described here: https://codeforces.com/apiHelp/methods
type ProblemsetProblems struct {
	Problems []Problem
	Stats    []ProblemStatistics
}

//RequestComments ... the object returned from blogEntry.comments described here:
//											https://codeforces.com/apiHelp/methods
type RequestComments struct {
	Status   string    `json:"status"`
	Commentv string    `json:"comment"`
	Result   []Comment `json:"result"`
}

//RequestBlog ... the object returned from blogEntry.view described here:
//											https://codeforces.com/apiHelp/methods
type RequestBlog struct {
	Status   string    `json:"status"`
	Commentv string    `json:"comment"`
	Result   BlogEntry `json:"result"`
}

//RequestHacks ... the object returned from contest.hacks described here:
//											https://codeforces.com/apiHelp/methods
type RequestHacks struct {
	Status   string `json:"status"`
	Commentv string `json:"comment"`
	Result   []Hack `json:"result"`
}

//RequestContests ... the object returned from contest.list described here:
//											https://codeforces.com/apiHelp/methods
type RequestContests struct {
	Status   string    `json:"status"`
	Commentv string    `json:"comment"`
	Result   []Contest `json:"result"`
}

//RequestRatingChange ... the object returned from contest.ratingChanges described here:
//											https://codeforces.com/apiHelp/methods
type RequestRatingChange struct {
	Status   string         `json:"status"`
	Commentv string         `json:"comment"`
	Result   []RatingChange `json:"result"`
}

//RequestContestStandings ... the object returned from contest.standings described here:
//											https://codeforces.com/apiHelp/methods
type RequestContestStandings struct {
	Status   string           `json:"status"`
	Commentv string           `json:"comment"`
	Result   ContestStandings `json:"result"`
}

//RequestContestStatus ... the object returned from contest.status described here:
//											https://codeforces.com/apiHelp/methods
type RequestContestStatus struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   []Submission `json:"result"`
}

//RequestPsetProblems ... the object returned from problemset.problems described here:
//											https://codeforces.com/apiHelp/methods
type RequestPsetProblems struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   PsetProblems `json:"result"`
}

//RequestRecentStatus ... the object returned from problemset.recentStatus described here:
//											https://codeforces.com/apiHelp/methods
type RequestRecentStatus struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   []Submission `json:"result"`
}

//RequestRecentActions ... the object returned from recentActions described here:
//											https://codeforces.com/apiHelp/methods
type RequestRecentActions struct {
	Status   string         `json:"status"`
	Commentv string         `json:"comment"`
	Result   []RecentAction `json:"result"`
}

//RequestBlogEntries ... the object returned from user.blogEntries described here:
//											https://codeforces.com/apiHelp/methods
type RequestBlogEntries struct {
	Status   string      `json:"status"`
	Commentv string      `json:"comment"`
	Result   []BlogEntry `json:"result"`
}

//RequestUserInfo ... the object returned from user.info described here:
//											https://codeforces.com/apiHelp/methods
type RequestUserInfo struct {
	Status   string `json:"status"`
	Commentv string `json:"comment"`
	Result   []User `json:"result"`
}

//RequestRatedList ... the object returned from user.ratedList described here:
//											https://codeforces.com/apiHelp/methods
type RequestRatedList struct {
	Status   string `json:"status"`
	Commentv string `json:"comment"`
	Result   []User `json:"result"`
}

//RequestUserRating ... the object returned from user.rating described here:
//											https://codeforces.com/apiHelp/methods
type RequestUserRating struct {
	Status   string         `json:"status"`
	Commentv string         `json:"comment"`
	Result   []RatingChange `json:"result"`
}

//RequestUserStatus ... the object returned from user.status described here:
//											https://codeforces.com/apiHelp/methods
type RequestUserStatus struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   []Submission `json:"result"`
}
