package main

import (
	"sort"
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
	urls := map[string]int{}

	c.max_depth = max_depth
	c.crawl(page_url, 0, urls)

	keys := make([]string, 0, len(urls))
	for k := range urls {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func (c *Crawler) crawl(page_url string, depth int, urls map[string]int) {

	// return if we've reached max depth
	if depth >= c.max_depth {
		return
	}

	// return if we've already visited this url
	if _, ok := urls[page_url]; ok {
		return
	}

	urls[page_url] = depth

	page_urls := c.parser.extractURLs(page_url)
	for _, u := range page_urls {
		c.crawl(u, depth+1, urls)
	}
}
