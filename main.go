package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printResults(urls *[]string) {
	fmt.Println("\nCrawl Results:")

	for _, url := range *urls {
		indent := strings.Count(url, "/")
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Print(url, "\n")
	}
}

func main() {

	// parse arguments
	url := flag.String("url", "", "URL to crawl")
	max_depth := flag.Int("max-depth", 100, "Maximum depth to crawl")

	flag.Usage = func() {
		basename := filepath.Base(os.Args[0])
		fmt.Printf("Usage: %s\n", basename)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	crawler := NewCrawler()
	urls := crawler.run(*url, *max_depth)

	printResults(&urls)

}
