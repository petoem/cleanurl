package cleanurl

import (
	"net/url"
	"strings"
)

// copyURL returns a copy of u.
func copyURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	newurl := new(url.URL)
	*newurl = *u
	if u.User != nil {
		newurl.User = new(url.Userinfo)
		*newurl.User = *u.User
	}
	return newurl
}

// makeFQDN adds a dot to the end of the hostname.
//
// Note: It does not check if hostname is an IP address.
func makeFQDN(hostname string) string {
	if strings.HasSuffix(hostname, ".") {
		return hostname
	} else {
		return hostname + "."
	}
}
