# Clean URL

[![Go Reference](https://pkg.go.dev/badge/github.com/petoem/cleanurl.svg)](https://pkg.go.dev/github.com/petoem/cleanurl)

`Copy Clean Link` for Go, automatically removes tracking parameters like `utm_source`, `si`, and `utm_campaign` from a URL.
Uses tracking parameter lists from Firefox and Brave. See [./internal/lists](internal/lists/) directory.

## Usage

```go
link, _ := url.Parse("https://youtu.be/dQw4w9WgXcQ?si=YxOc4tn6Bg0zqKor")
cleanlink := cleanurl.CopyCleanURL(link)
fmt.Println(cleanlink) // https://youtu.be/dQw4w9WgXcQ
```
