package main

import (
	"fmt"
)

func main() {

	root_url := "https://www.monzo.com"

	crawler := NewCrawler()
	urls := crawler.run(root_url, 3)

	fmt.Println(urls)
}
