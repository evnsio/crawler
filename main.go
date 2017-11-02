package main

import (
	"fmt"
)

func main() {

	root_url := "https://www.monzo.com"

	urls := make([]string, 0)

	crawler := NewCrawler(OptionMaxDepth(2))
	crawler.crawl(root_url, 0, &urls)

	fmt.Println(urls)
}
