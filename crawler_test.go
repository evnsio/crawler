package main

import (
	"testing"
	"fmt"
)

func mockMultiFetchPage(page_url string) string {
	switch page_url {
	case "http://foo.com":
		fmt.Println("base page!")
	case "http://foo.com/bar":
		fmt.Println("bar page!")
	default:
		panic("Something terrible has happened")
	}
	return ""
}

func TestCrawlerReturnsSomething(t *testing.T) {
	// foo.com -> /bar -> /baz

	urls := Crawl("http://foo.com", 0)
	fmt.Println(urls)

	//t.Errorf("Expected no urls. Got %d", len(urls))
}