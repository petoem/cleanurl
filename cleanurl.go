package cleanurl

import (
	"net/url"
	"strings"
	"sync"

	"github.com/petoem/cleanurl/filter"
	"github.com/petoem/cleanurl/internal/lists"
)

var loadfilters sync.Once
var filters filter.FilterList

// CleanURL returns u with tracking parameters removed.
func CleanURL(u *url.URL) {
	loadfilters.Do(func() {
		filters = lists.LoadFilters()
	})

	if u == nil {
		return
	}

	hostname := strings.ToLower(makeFQDN(u.Hostname()))
	for _, filter := range filters.Includes(hostname) {
		filter.Apply(u)
	}
}

// CopyCleanURL returns a copy of u with tracking parameters removed.
func CopyCleanURL(u *url.URL) *url.URL {
	newurl := copyURL(u)
	CleanURL(newurl)
	return newurl
}
