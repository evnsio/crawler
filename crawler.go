package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func fetchPage(page_url string) string {
	response, err := http.Get(page_url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(bodyBytes)
}

func getDomain(page_url string) string {
	u, err := url.Parse(string(page_url))

	if err != nil {
		log.Fatal(err)
	}

	return u.Hostname()
}

func generateURL(host string, path string) string {
	u, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}

	base, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}

	return base.ResolveReference(u).String()
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
