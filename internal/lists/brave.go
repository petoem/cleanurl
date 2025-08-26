package lists

import (
	"encoding/json"
	"io/fs"
	"strings"
)

type braveFilter struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
	Params  []string `json:"params"`
}

// parseFilterlistBrave returns the embedded filters or panics.
func parseFilterlistBrave() []braveFilter {
	data, err := fs.ReadFile(filterlists, "brave/clean-urls.json")
	if err != nil {
		panic(err)
	}
	var bfilterlists []braveFilter
	err = json.Unmarshal(data, &bfilterlists)
	if err != nil {
		panic(err)
	}
	return bfilterlists
}

func wildcardBraveURLToDomain(wurl string) string {
	// trim protocols
	wurl = strings.TrimPrefix(wurl, "*://")
	wurl = strings.TrimPrefix(wurl, "https://")
	wurl = strings.TrimPrefix(wurl, "http://")
	// wildcard at the beginning
	wurl = strings.TrimPrefix(wurl, "*")
	// discard the url path
	wurl, _, _ = strings.Cut(wurl, "/")
	wurl = strings.ToLower(wurl)
	wurl = wurl + "."
	return wurl
}
