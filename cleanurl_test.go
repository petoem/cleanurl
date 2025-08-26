package cleanurl_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/petoem/cleanurl"
)

func TestCopyCleanURL(t *testing.T) {
	tests := []struct {
		name string
		arg  *url.URL
		want *url.URL
	}{
		{
			name: "Remove tracking",
			arg:  urlMustParse("https://youtu.be/dQw4w9WgXcQ?si=YxOc4tn6Bg0zqKor"),
			want: urlMustParse("https://youtu.be/dQw4w9WgXcQ"),
		},
		{
			name: "Remove tracking",
			arg:  urlMustParse("https://test.petoe.me/random_page/?utm_source=share"),
			want: urlMustParse("https://test.petoe.me/random_page/"),
		},
		{
			name: "Remove tracking, leave other params",
			arg:  urlMustParse("http://this.is.a.test.petoe.me/long/path/here/?stays=here&utm_source=social&utm_medium=paid_social&utm_campaign=sale"),
			want: urlMustParse("http://this.is.a.test.petoe.me/long/path/here/?stays=here"),
		},
		{
			name: "Remove tracking from URL with FQDN",
			arg:  urlMustParse("https://test.petoe.me./random_page/?utm_source=share"),
			want: urlMustParse("https://test.petoe.me./random_page/"),
		},
		{
			name: "Clean URL remains clean and complete",
			arg:  urlMustParse("https://username:password@test.petoe.me/random_page/"),
			want: urlMustParse("https://username:password@test.petoe.me/random_page/"),
		},
		{
			name: "Handle nil",
			arg:  nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run("Clean URL", func(t *testing.T) {
			if got := cleanurl.CopyCleanURL(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyCleanURL() = %v, want %v", got, tt.want)
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
