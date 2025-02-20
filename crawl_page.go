package main

import ()

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	// Check hostname
	baseURL, err := normalizeURL(rawBaseURL)
	if err != nil {
		return
	}
	currentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}
	if baseURL != currentURL {
		return
	}

	// Get HTML Content & URLs
	htmlContent, err := getHTML(rawCurrentURL)
	if err != nil {
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
