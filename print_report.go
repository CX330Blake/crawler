package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	printHeader(baseURL)
	printResult(pages, baseURL)
}

func printHeader(baseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Printf("=============================\n")
}

func printResult(pages map[string]int, baseURL string) {
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if pages[keys[i]] == pages[keys[j]] {
			return keys[i] > keys[j]
		}
		return pages[keys[i]] > pages[keys[j]]
	})

	for _, k := range keys {
		fmt.Printf("Found %d internal links to %s\n", pages[k], baseURL+k)
	}
}
