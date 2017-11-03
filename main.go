package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

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

	// run the crawler
	crawler := NewCrawler()
	root := crawler.run(*url, *max_depth)

	root.toSiteMap()
}
