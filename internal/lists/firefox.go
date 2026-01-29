package lists

import (
	"encoding/json"
	"io/fs"
)

// List of Firefox filters in json files
type firefoxFilters map[string]firefoxFilter

// Firefox filter for domains
type firefoxFilter struct {
	QueryParams []string `json:"queryParams"`
	Origins     []string `json:"origins"`
	IsGlobal    bool     `json:"isGlobal"`
}

// parseFilterlistFirefox returns the embedded filters at filepath or panics.
func parseFilterlistFirefox(filepath string) firefoxFilters {
	data, err := fs.ReadFile(listsFS, filepath)
	if err != nil {
		panic(err)
	}
	var firefoxfilters firefoxFilters
	err = json.Unmarshal(data, &firefoxfilters)
	if err != nil {
		panic(err)
	}
	return firefoxfilters
}
