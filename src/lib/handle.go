package lib

import (
	"net/http"
	"strings"
)

const profileURL string = "https://codeforces.com/profile/"
const homeURL string = "https://codeforces.com/"

//FindCurrentHandle... finds the present handle of the user who used to have handle `handle`
func FindCurrentHandle(handle string) string {
	req, _ := http.Get(profileURL + handle)
	redirectURL := req.Request.URL.String()
	if redirectURL != homeURL {
		return handle
	}
	return strings.TrimPrefix(redirectURL, profileURL)
}


