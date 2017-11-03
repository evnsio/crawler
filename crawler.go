package main

import (
	"fmt"
	"sync"
)

type Crawler struct {
	parser *PageParser
	max_depth int
}

func NewCrawler(options ...func(*Crawler)) *Crawler {
	crawler := &Crawler{}
	crawler.parser = NewPageParser(fetchPage)
	crawler.max_depth = 100

	for _, option := range options{
		option(crawler)
	}

	return crawler
}

func (c *Crawler) run(page_url string, max_depth int) []string {
	c.max_depth = max_depth

	urls := make(chan string, 10000)
	var wg sync.WaitGroup

	wg.Add(1)
	go c.crawl(page_url, 0, urls, &wg)

	//
	//// extract the keys and sort
	//keys := make([]string, 0, len(urls))
	//for k := range urls {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)
	fmt.Println("Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Done waiting")

	fmt.Println("Closing Channel")
	close(urls)
	fmt.Println("Channel Closed")

	for elem := range urls {
		fmt.Println(elem)
	}



	return []string{"hello"}
}

func (c *Crawler) crawl(page_url string, depth int, urls chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// return if we've reached max depth
	if depth >= c.max_depth {
		return
	}

	//// return if we've already visited this url
	//if previous_depth, ok := urls[page_url]; ok {
	//	if previous_depth < depth {
	//		return
	//	}
	//}

	urls <- page_url

	page_urls := c.parser.extractURLs(page_url)
	for _, u := range page_urls {
		wg.Add(1)
		go c.crawl(u, depth+1, urls, wg)
	}
}
