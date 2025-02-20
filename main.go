package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseURL := os.Args[1]

	fmt.Printf("starting crawl of: %v\n", baseURL)

	var resultMap = make(map[string]int)
	crawlPage(baseURL, baseURL, resultMap)

	// fmt.Printf("%v\n", len(resultMap))

}
