package api

type Comment struct {
	Id                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	Text                string `json:"text"`
	ParentCommentId     int    `json:"parentCommentId"`
	Rating              int    `json:"rating"`
}

type BlogEntry struct {
	Id                      int      `json:"id"`
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

type User struct {
	Handle                  string `json:"handle"`
	Email                   string `json:"email"`
	VkId                    string `json:"vkId"`
	OpenId                  string `json:"openId"`
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

type RecentAction struct {
	TimeSeconds int       `json:"timeSeconds"`
	BlogEntryv  BlogEntry `json:"blogEntry"`
	Commentv    Comment   `json:"comment"`
}

type RatingChange struct {
	ContestId               int    `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	Rank                    int    `json:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating"`
	NewRating               int    `json:"newRating"`
}

type Contest struct {
	Id                  int    `json:"id"`
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

type Member struct {
	Handle string `json:"handle"`
}

type Party struct {
	ContestId        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	TeamId           int      `json:"teamId"`
	TeamName         string   `json:"teamName"`
	Ghost            bool     `json:"ghost"`
	Room             int      `json:"room"`
	StartTimeSeconds int      `json:"startTimeSeconds"`
}

type Problem struct {
	ContestId int      `json:"contestId"`
	PsetName  string   `json:"problemsetName"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Typev     string   `json:"type"`
	Points    float64  `json:"points"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}

type ProblemStatistics struct {
	ContestId   int    `json:"contestId"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}

type PsetProblems struct {
	Problems []Problem           `json:"problems"`
	PStats   []ProblemStatistics `json:"problemStatistics"`
}

type Submission struct {
	Id                  int     `json:"id"`
	ContestId           int     `json:"contestId"`
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

type JProtocol struct {
	Manual   string `json:"manual"`
	Protocol string `json:"protocol"`
	Verdict  string `json:"verdict"`
}

type Hack struct {
	Id                  int       `json:"id"`
	CreationTimeSeconds int       `json:"creationTimeSeconds"`
	Hacker              Party     `json:"hacker"`
	Defender            Party     `json:"defender"`
	Verdict             string    `json:"verdict"`
	Problemv            Problem   `json:"problem"`
	Test                string    `json:"test"`
	JudgeProtocol       JProtocol `json:"judgeProtocol"`
}

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

type ProblemResult struct {
	Points                    float64 `json:"points"`
	Penalty                   int     `json:"penalty"`
	RejectedAttemptCount      int     `json:"rejectedAttemptCount"`
	Typev                     string  `json:"type"`
	BestSubmissionTimeSeconds int     `json:"bestSubmissionTimeSeconds"`
}

type ContestStandings struct {
	Contestv Contest       `json:"contest"`
	Problems []Problem     `json:"problems"`
	Rows     []RanklistRow `json:"rows"`
}

type ProblemsetProblems struct {
	Problems []Problem
	Stats    []ProblemStatistics
}

type RequestComments struct {
	Status   string    `json:"status"`
	Commentv string    `json:"comment"`
	Result   []Comment `json:"result"`
}

type RequestBlog struct {
	Status   string    `json:"status"`
	Commentv string    `json:"comment"`
	Result   BlogEntry `json:"result"`
}

type RequestHacks struct {
	Status   string `json:"status"`
	Commentv string `json:"comment"`
	Result   []Hack `json:"result"`
}

type RequestContests struct {
	Status   string    `json:"status"`
	Commentv string    `json:"comment"`
	Result   []Contest `json:"result"`
}

type RequestRatingChange struct {
	Status   string         `json:"status"`
	Commentv string         `json:"comment"`
	Result   []RatingChange `json:"result"`
}

type RequestContestStandings struct {
	Status   string           `json:"status"`
	Commentv string           `json:"comment"`
	Result   ContestStandings `json:"result"`
}

type RequestContestStatus struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   []Submission `json:"result"`
}

type RequestPsetProblems struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   PsetProblems `json:"result"`
}

type RequestRecentStatus struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   []Submission `json:"result"`
}

type RequestRecentActions struct {
	Status   string         `json:"status"`
	Commentv string         `json:"comment"`
	Result   []RecentAction `json:"result"`
}

type RequestBlogEntries struct {
	Status   string      `json:"status"`
	Commentv string      `json:"comment"`
	Result   []BlogEntry `json:"result"`
}

type RequestUserInfo struct {
	Status   string `json:"status"`
	Commentv string `json:"comment"`
	Result   []User `json:"result"`
}

type RequestRatedList struct {
	Status   string `json:"status"`
	Commentv string `json:"comment"`
	Result   []User `json:"result"`
}

type RequestUserRating struct {
	Status   string         `json:"status"`
	Commentv string         `json:"comment"`
	Result   []RatingChange `json:"result"`
}

type RequestUserStatus struct {
	Status   string       `json:"status"`
	Commentv string       `json:"comment"`
	Result   []Submission `json:"result"`
}
