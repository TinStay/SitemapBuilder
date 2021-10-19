package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	link "github.com/Basics/src/github.com/TinStay/LinkParser"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "URL to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "Maximum number of links deep to traverse")

	flag.Parse()

	// Build proper urls with our links
	pages := getPageLinks(*urlFlag)

	for _, url := range pages {
		fmt.Println(url)
	}
	// Find all pages (BFS)

	// Print out XML

}

func getPageLinks(urlStr string) []string{
	// GET webpage
	resp, err := http.Get(urlStr)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Current page url
	reqUrl := resp.Request.URL

	// https/http + domain name
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	// Filter out links to different domains
	return filterLinks(hrefs(resp.Body, base), withPrefix(base))
}


func hrefs(r io.Reader, base string) []string{
	var ret []string

	// Parse links on the page
	links, _ := link.Parse(r)

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}

func filterLinks(links []string, keepFn func(string) bool) []string {
	var ret []string

	for _, link := range links {
		// Keep url with the same base
		if keepFn(link){
			ret = append(ret, link)
		}
	}

	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
