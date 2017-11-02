package main

type Crawler struct {
	parser *PageParser
	max_depth int
}

func OptionMaxDepth(max_depth int) func(*Crawler) {
	return func(c *Crawler) {
		c.max_depth = max_depth
	}
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

func (c *Crawler) crawl(page_url string, depth int, urls *[]string) {
	if depth >= c.max_depth {
		return
	}

	if page_url 
	*urls = append(*urls, page_url)

	page_urls := c.parser.extractURLs(page_url)
	for _, u := range page_urls {
		c.crawl(u, depth+1, urls)
	}
}
