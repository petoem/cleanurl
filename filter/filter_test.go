package filter_test

import (
	"net/url"
	"reflect"
	"slices"
	"testing"

	"github.com/petoem/cleanurl/filter"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		filter filter.Filter
		input  *url.URL
		want   *url.URL
	}{
		{
			filter: filter.NewFilter([]string{"remove1", "remove2"}),
			input:  urlMustParse("https://test.petoe.me/page/?remove1=a&remove3=c"),
			want:   urlMustParse("https://test.petoe.me/page/?remove3=c"),
		},
		{
			filter: filter.NewFilter([]string{"remove3"}),
			input:  urlMustParse("https://test.petoe.me/page/?remove1=a&remove3=c"),
			want:   urlMustParse("https://test.petoe.me/page/?remove1=a"),
		},
		{
			filter: filter.NewFilter([]string{"remove4"}),
			input:  urlMustParse("https://test.petoe.me/page/?remove1=a&remove3=c"),
			want:   urlMustParse("https://test.petoe.me/page/?remove1=a&remove3=c"),
		},
	}
	for _, tt := range tests {
		t.Run("Apply", func(t *testing.T) {
			if tt.filter.Apply(tt.input); !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("Filter.Apply() = %v, want %v", tt.input, tt.want)
			}
		})
	}

	t.Run("Merge", func(t *testing.T) {
		f1 := filter.NewFilter([]string{"remove3"})
		f2 := filter.NewFilter([]string{"remove4"})
		fw := filter.NewFilter(([]string{"remove3", "remove4"}))
		f1.Merge(f2)
		if !reflect.DeepEqual(f1, fw) {
			t.Errorf("Filter.Merge() = %v, want %v", f1, fw)
		}
	})

	t.Run("AddQueryParams", func(t *testing.T) {
		additionalparams := []string{"remove3", "removeX", "removeY"}
		f1 := filter.NewFilter(additionalparams[:1])
		fw := filter.NewFilter(additionalparams)

		f1.AddQueryParams(additionalparams[1:])
		if !reflect.DeepEqual(f1, fw) {
			t.Errorf("Filter.AddQueryParams() = %v, want %v", f1, fw)
		}
	})
}

func TestFilterList(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		// setup test
		f := filter.NewFilter([]string{"remove"})
		domainsWanted := []string{".", "test.petoe.me.", "."}
		filterlist := makeFilterListForTest(domainsWanted, f)

		// check if all domains are present
		domains := filterlist.Domains()
		for _, d := range domainsWanted {
			if !slices.Contains(domains, d) {
				t.Errorf("FilterList.Add(): missing domain %v", d)
			}
		}

		// check if all filters are present
		filters := filterlist.Filters()
		for _, fi := range filters {
			if !reflect.DeepEqual(fi, f) {
				t.Errorf("FilterList.Filters(): unkown filter %v", fi)
			}
		}
	})

	f := filter.NewFilter([]string{"remove"})
	tests := []struct {
		name   string
		flist  filter.FilterList
		domain string
		want   []filter.Filter
	}{
		{
			name:   "wildcard match",
			flist:  makeFilterListForTest([]string{".", "test.petoe.me.", "."}, f),
			domain: "global.site.com.",
			want:   []filter.Filter{f},
		},
		{
			name:   "wildcard and exact match",
			flist:  makeFilterListForTest([]string{".", "test.petoe.me.", "."}, f),
			domain: "test.petoe.me.",
			want:   []filter.Filter{f, f},
		},
		{
			name:   "two wildcard match",
			flist:  makeFilterListForTest([]string{".", ".test.petoe.me.", "."}, f),
			domain: "this.is.a.test.petoe.me.",
			want:   []filter.Filter{f, f},
		},
	}
	for _, tt := range tests {
		t.Run("Includes", func(t *testing.T) {
			if got := tt.flist.Includes(tt.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterList.Includes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func urlMustParse(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}

func makeFilterListForTest(domains []string, f filter.Filter) filter.FilterList {
	filterlist := make(filter.FilterList)
	for _, domain := range domains {
		filterlist.Add(domain, f)
	}
	return filterlist
}
