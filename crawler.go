package main

import (
)

type Crawler struct {
	parser *PageParser
}

func NewCrawler(pageParser *PageParser) *Crawler {
	return &Crawler{parser: pageParser}
}

func (c *Crawler) crawl(page_url string, depth int) []string {
	pageParser := c.parser
	urls := pageParser.extractURLs(page_url)

	if depth >= 1 {
		return urls
	}

	for _, u := range urls {
		new_urls := Crawl(u, depth+1)
		urls = append(urls, new_urls...)
	}

	return urls
}


func Crawl(page_url string, depth int) []string {

	pageParser := NewPageParser(fetchPage)
	urls := pageParser.extractURLs(page_url)

	if depth >= 1 {
		return urls
	}

	for _, u := range urls {
		new_urls := Crawl(u, depth+1)
		urls = append(urls, new_urls...)
	}

	return urls
}
