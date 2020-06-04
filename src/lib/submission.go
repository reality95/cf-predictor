package lib

import (
	"github.com/reality95/cf-predictor/src/api"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

const prefixSubmission string = `https://codeforces.com/contest/`
const middleSubmission string = `/submission/`

const prefixCode string = `<pre id="program-source-text"[^>]*>`
const suffixCode string = `</pre>`

func GetSource(s api.Submission) (sourceText string, err error) {
	resp, err := http.Get(prefixSubmission + strconv.Itoa(s.ContestID) + middleSubmission + strconv.Itoa(s.ID))
	if err != nil {
		return sourceText, err
	}
	plainText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sourceText, err
	}
	re := regexp.MustCompile(prefixCode + `[^<]+` + suffixCode)
	rePrefix := regexp.MustCompile(prefixCode)
	reSuffix := regexp.MustCompile(suffixCode)
	sourceText = string(reSuffix.ReplaceAll(rePrefix.ReplaceAll(re.Find([]byte(plainText)), nil), nil))
	return replaceSpecialCharacters(sourceText), err
}

func replaceSpecialCharacters(sourceText string) string {
	quote := regexp.MustCompile(`&quot;`)
	lt := regexp.MustCompile(`&lt;`)
	gt := regexp.MustCompile(`&gt;`)
	and := regexp.MustCompile(`&amp;`)
	apo := regexp.MustCompile(`&#39;`)
	var ans []byte
	ans = quote.ReplaceAll([]byte(sourceText), []byte(`"`))
	ans = lt.ReplaceAll(ans, []byte(`<`))
	ans = gt.ReplaceAll(ans, []byte(`>`))
	ans = and.ReplaceAll(ans, []byte(`&`))
	ans = apo.ReplaceAll(ans, []byte(`'`))
	return string(ans)
}
