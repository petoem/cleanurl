package lists

import (
	"embed"
	"strings"

	"github.com/petoem/cleanurl/filter"
)

const (
	FIREFOXLISTLGPL = "firefox/LGPL/StripOnShareLGPL.json"
	FIREFOXLISTMPL2 = "firefox/MPL2/StripOnShare.json"

	BRAVELIST = "brave/clean-urls.json"
)

//go:embed firefox brave
var listsFS embed.FS

func LoadFilters() filter.FilterList {
	filterlist := make(filter.FilterList)
	// Firefox
	for _, filepath := range []string{FIREFOXLISTLGPL, FIREFOXLISTMPL2} {
		filters := parseFilterlistFirefox(filepath)
		for _, record := range filters {
			if record.IsGlobal {
				filterlist.Add(strings.ToLower("."), filter.NewFilter(record.QueryParams))
			}
			for _, site := range record.Origins {
				filterlist.Add(strings.ToLower(site+"."), filter.NewFilter(record.QueryParams))
			}
		}
	}
	// Brave
	filters := parseFilterlistBrave(BRAVELIST)
	for _, b := range filters {
		for _, url := range b.Include {
			filterlist.Add(wildcardBraveURLToDomain(url), filter.NewFilter(b.Params))
		}
	}
	return filterlist
}
