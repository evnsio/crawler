package main

import (
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

func (c *Crawler) run(page_url string, max_depth int) *Page {
	c.max_depth = max_depth
	var wg sync.WaitGroup

	wg.Add(1)

	root := NewPage(page_url, 0, nil)
	go c.crawl(root, &wg)

	wg.Wait()

	return root
}

func (c *Crawler) crawl(page *Page, wg *sync.WaitGroup) {
	defer wg.Done()

	// return if we've reached max depth
	if page.depth == c.max_depth {
		return
	}

	page_urls := c.parser.extractURLs(page.url)
	page.scraped = true
	for _, child_url := range page_urls {
		child := NewPage(child_url, page.depth+1, page)
		page.children = append(page.children, child)

		wg.Add(1)
		go c.crawl(child, wg)
	}
}
