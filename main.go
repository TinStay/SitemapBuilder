package main

import (
	"flag"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "URL to build a sitemap for")

	flag.Parse()

	// GET webpage


	// Parse links on the page

	// Build proper urls with our links

	// Filter out links to different domains

	// Find all pages (BFS) 

	// Print out XML

}