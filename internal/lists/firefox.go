package lists

import (
	"encoding/json"
	"io/fs"
)

// List of Firefox filters in json files
type firefoxFilters map[string]firefoxFilter

// Firefox filter for domains
type firefoxFilter struct {
	QueryParams   []string `json:"queryParams"`
	TopLevelSites []string `json:"topLevelSites"`
}

// parseFilterlistFirefox returns the embedded filters or panics.
func parseFilterlistFirefox() firefoxFilters {
	data, err := fs.ReadFile(filterlists, "firefox/LGPL/StripOnShareLGPL.json")
	if err != nil {
		panic(err)
	}
	var fffilterlists firefoxFilters
	err = json.Unmarshal(data, &fffilterlists)
	if err != nil {
		panic(err)
	}
	return fffilterlists
}
