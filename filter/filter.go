package filter

import (
	"maps"
	"net/url"
	"slices"
	"strings"
)

// Filter is a set of URL query parameters that should be filtered.
type Filter map[string]struct{}

// NewFilter returns a Filter filtering params.
func NewFilter(params []string) Filter {
	filter := make(Filter)
	filter.AddQueryParams(params)
	return filter
}

// AddQueryParams adds the query params to the Filter.
func (f Filter) AddQueryParams(params []string) {
	mapFunc(params, strings.ToLower) // make sure everything is lower case
	maps.Insert(f, maps.All(toSet(params)))
}

// Merge adds all query parameters from filter to f.
func (f Filter) Merge(filter Filter) {
	maps.Insert(f, maps.All(filter))
}

// Apply removes all query parameters from URL u that are in filter f.
func (f Filter) Apply(u *url.URL) {
	q := u.Query()
	maps.DeleteFunc(q, func(param string, _ []string) bool {
		_, ok := f[strings.ToLower(param)]
		return ok
	})
	u.RawQuery = q.Encode()
}

// FilterList is map of domains or "wildcard" domains in FQDN format to a Filter.
// Wildcards are represented by a leading dot e.g. `.awesome.tld.` which matches `something.very.awesome.tld.`
//
// Note: FilterList is case-sensitive.
type FilterList map[string]Filter

// Includes checks if domain is included in the FilterList and returns the associated Filters.
func (flist FilterList) Includes(domain string) []Filter {
	var filters []Filter
	for d := range flist {
		// .awesome.tld includes something.very.awesome.tld
		if strings.HasPrefix(d, ".") {
			if strings.HasSuffix(domain, d) {
				filters = append(filters, flist[d])
			}
		} else {
			// awesome.tld exact match
			if d == domain {
				filters = append(filters, flist[d])
			}
		}
	}
	return filters
}

// Add adds filter for the domain to FilterList, merges filters if domain already exists.
//
// Note: Does not check if domain is in FQDN format.
func (flist FilterList) Add(domain string, filter Filter) {
	if _, ok := flist[domain]; !ok {
		flist[domain] = make(Filter)
	}

	flist[domain].Merge(filter)
}

// Domains returns all domains present.
func (flist FilterList) Domains() []string {
	return slices.Collect(maps.Keys(flist))
}

// Filters returns all Filter entries.
func (flist FilterList) Filters() []Filter {
	return slices.Collect(maps.Values(flist))
}
