package main

import (
	"sort"
	"sync"
)

type Crawler struct {
	parser    *PageParser
	max_depth int
}

func NewCrawler(options ...func(*Crawler)) *Crawler {
	crawler := &Crawler{}
	crawler.parser = NewPageParser(fetchPage)
	crawler.max_depth = 100

	for _, option := range options {
		option(crawler)
	}

	return crawler
}

func (c *Crawler) run(page_url string, max_depth int) []string {
	c.max_depth = max_depth
	urls := NewSafeMap()
	var wg sync.WaitGroup

	wg.Add(1)
	go c.crawl(page_url, 0, urls, &wg)

	// wait for all crawls to complete
	wg.Wait()

	// extract the keys and sort
	// (safe to access map direct here)
	keys := make([]string, 0, len(urls.v))
	for k := range urls.v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func (c *Crawler) crawl(page_url string, depth int, urls *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()

	// return if we've reached max depth
	if depth >= c.max_depth {
		return
	}

	// return if we've already visited this url
	if previous_depth, ok := urls.Value(page_url); ok {
		if previous_depth < depth {
			return
		}
	}

	urls.Set(page_url, depth)

	page_urls := c.parser.extractURLs(page_url)
	for _, u := range page_urls {
		wg.Add(1)
		go c.crawl(u, depth+1, urls, wg)
	}
}
