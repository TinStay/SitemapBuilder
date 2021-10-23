package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	link "github.com/Basics/src/github.com/TinStay/LinkParser"
)
const xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls []loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "URL to build a sitemap for")
	maxDepth := flag.Int("depth", 10, "Maximum number of links deep to traverse")

	flag.Parse()
	
	// Find all pages (BFS)
	pages := bfs(*urlFlag, *maxDepth)

	// Print out XML
	toXml:= urlset{
		Xmlns: xmlns,
	}

	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	// Print xml version
	fmt.Print(xml.Header)

	enc := xml.NewEncoder(os.Stdout)

	// Format xml
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil{
		panic(err)
	}

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

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})

	var q map[string]struct{}

	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++{
		q, nq = nq, make(map[string]struct{})

		// Exit if no links
		if len(q) == 0{
			break
		}

		for url, _ := range q{
			if _, ok := seen[url]; ok {
				continue
			}
			// Add url to seen map
			seen[url] = struct{}{}
			
			for _, link := range getPageLinks(urlStr) {
				if _, ok := seen[link]; !ok{
					nq[link] = struct{}{}
				}
			}
		}
	}
	
	ret := make([]string, 0, len(seen))

	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}
