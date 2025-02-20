package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	// Check hostname are the same
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	// Normalized the URL
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: couldn't normalize %v: %v\n", rawCurrentURL, err)
	}

	// Increment if visited
	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1
	fmt.Printf("crawling url: %v\n", rawCurrentURL)

	htmlContent, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: couldn't get HTML from %v: %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(htmlContent, rawBaseURL)
	if err != nil {
		return
	}

	// Recusive call
	for _, url := range urls {

		// check it's not in map already
		if _, exists := pages[url]; !exists {
			pages[url] = 1
			crawlPage(rawBaseURL, url, pages)
		} else {
			pages[url]++
		}
	}

}
