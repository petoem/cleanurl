package lists

import (
	"embed"
	"strings"

	"github.com/petoem/cleanurl/filter"
)

//go:embed firefox brave
var filterlists embed.FS

func LoadFilters() filter.FilterList {
	filters := make(filter.FilterList)
	// Firefox
	fffilters := parseFilterlistFirefox()
	for _, fffrecord := range fffilters {
		for _, site := range fffrecord.TopLevelSites {
			if site == "*" {
				site = "."
			} else {
				site = site + "."
			}
			filters.Add(strings.ToLower(site), filter.NewFilter(fffrecord.QueryParams))
		}
	}
	// Brave
	bfilters := parseFilterlistBrave()
	for _, b := range bfilters {
		for _, url := range b.Include {
			filters.Add(wildcardBraveURLToDomain(url), filter.NewFilter(b.Params))
		}
	}
	return filters
}
