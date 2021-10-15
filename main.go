package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	link "github.com/Basics/src/github.com/TinStay/LinkParser"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "URL to build a sitemap for")

	flag.Parse()

	// GET webpage
	resp, err := http.Get(*urlFlag)

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

	var hrefs []string

	// Parse links on the page
	links, _ := link.Parse(resp.Body)

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, url := range hrefs {
		fmt.Println(url)
	}

	// Build proper urls with our links

	// Filter out links to different domains

	// Find all pages (BFS)

	// Print out XML

}
