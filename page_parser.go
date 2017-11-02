package main

import (
	"bytes"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

type PageFetcher func(url string) string

type PageParser struct {
	get_page PageFetcher
}

func NewPageParser(fetcher PageFetcher) *PageParser {
	return &PageParser{get_page: fetcher}
}

func (p *PageParser) extractURLs(page_url string) []string {
	urls := make([]string, 0)
	url_domain := getDomain(page_url)
	page_contents := p.get_page(page_url)

	tokenizer := html.NewTokenizer(strings.NewReader(page_contents))
	anchorTag := []byte{'a'}

	for {
		switch tokenizer.Next() {
		case html.ErrorToken:
			return urls

		case html.StartTagToken:
			tag, hasAttr := tokenizer.TagName()
			if hasAttr && bytes.Equal(anchorTag, tag) {
				for {
					key, val, more := tokenizer.TagAttr()

					if string(key) == "href" {
						u, err := url.Parse(string(val))
						if err != nil {
							// should probably handle these
							continue
						}

						// We only want urls matching the source domain or relative paths
						if u.Hostname() == url_domain || !u.IsAbs() {
							urls = append(urls, generateURL(page_url, string(val)))

						}
					}

					if !more {
						break
					}
				}
			}
		}
	}

	return urls
}
