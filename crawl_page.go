package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// Skip other website's url
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	// Normalized the URL
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: couldn't normalize %v: %v\n", rawCurrentURL, err)
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling url: %v\n", rawCurrentURL)

	htmlContent, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: couldn't get HTML from %v: %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(htmlContent, cfg.baseURL)
	if err != nil {
		return
	}

	// Recusive call
	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

}
