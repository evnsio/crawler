package main

import (
	"flag"
	"fmt"
	"github.com/evnsio/crawler/asynccrawler"
	"os"
	"path/filepath"
	"time"
)

func main() {

	// parse arguments
	url := flag.String("url", "", "URL to crawl")
	max_depth := flag.Int("max-depth", -1, "Optional: Maximum depth to crawl")

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

	start := time.Now()

	// run the crawler
	crawler := asynccrawler.NewCrawler()
	root := crawler.Run(*url, *max_depth)

	root.PrintSiteMap()

	fmt.Printf("Crawl took %s\n", time.Since(start))
}
