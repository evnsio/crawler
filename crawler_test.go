package main

import (
	"testing"
	"fmt"
	"reflect"
)

func generateTestPage(urls... string) string {
	const simple_page = "<html><body>%s</body></html>"
	const anchor = "<a href=\"%s\"></a>"

	anchors := ""
	for _, url := range urls {
		anchors += fmt.Sprintf(anchor, url)
	}

	return fmt.Sprintf(simple_page, anchors)
}



func mockSingleFetchPage(page_url string) string {
	//level1 := "http://foo.com"
	return  ""
}

func mockMultiFetchPage(page_url string) string {
	level1 := "http://foo.com"
	level2 := "http://foo.com/bar"
	level3 := "http://foo.com/bar/baz"

	switch page_url {
	case level1:
		return generateTestPage(level2)
	case level2:
		return generateTestPage(level3)
	case level3:
		return ""
	}
	return ""
}

func mockMultiFetchTree(page_url string) string {
	level1 := "http://foo.com"
	level2a := "http://foo.com/bar/a"
	level2b := "http://foo.com/bar/b"
	level3 := "http://foo.com/bar/baz"

	switch page_url {
	case level1:
		return generateTestPage(level2a, level2b)
	case level2a:
		return generateTestPage(level3)
	case level2b:
		return generateTestPage()
	case level3:
		return generateTestPage()
	}
	return ""
}


func mockMultiFetchPageLoop(page_url string) string {
	level1 := "http://foo.com"
	level2 := "http://foo.com/bar"

	switch page_url {
	case level1:
		return generateTestPage(level2)
	case level2:
		return generateTestPage(level1)
	}
	return ""
}

func optionTestPageParser(fetcher PageFetcher) func(*Crawler) {
	return func(c *Crawler) {
		c.parser = NewPageParser(fetcher)
	}
}


func TestCrawlerReturnsSingleURLforSinglePage(t *testing.T) {
	// foo.com

	crawler := NewCrawler(
		optionTestPageParser(mockSingleFetchPage),
		OptionMaxDepth(3),
	)

	urls := make([]string, 0)
	crawler.crawl("http://foo.com", 0, &urls)

	expected := []string{"http://foo.com"}
	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerSingleThreadUnlimitedDepth(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPage),
	)

	urls := make([]string, 0)
	crawler.crawl("http://foo.com", 0, &urls)

	expected := []string{
		"http://foo.com",
		"http://foo.com/bar",
		"http://foo.com/bar/baz",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerSingleThreadDepthLimitedToZero(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPage),
		OptionMaxDepth(0),
	)

	urls := make([]string, 0)
	crawler.crawl("http://foo.com", 0, &urls)

	expected := []string{}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerSingleThreadDepthLimitedToOne(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPage),
		OptionMaxDepth(1),
	)

	urls := make([]string, 0)
	crawler.crawl("http://foo.com", 0, &urls)

	expected := []string{
		"http://foo.com",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}


func TestCrawlerSingleThreadDepthLimitedToTwo(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPage),
		OptionMaxDepth(2),
	)

	urls := make([]string, 0)
	crawler.crawl("http://foo.com", 0, &urls)

	expected := []string{
		"http://foo.com",
		"http://foo.com/bar",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerSimpleTreeUnlimitedDepth(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchTree),
		OptionMaxDepth(100),
	)

	urls := make([]string, 0)
	crawler.crawl("http://foo.com", 0, &urls)

	expected := []string{
		"http://foo.com",
		"http://foo.com/bar/a",
		"http://foo.com/bar/baz",
		"http://foo.com/bar/b",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerFetchPageLoop(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPageLoop),
	)

	urls := make([]string, 0)
	crawler.crawl("http://foo.com", 0, &urls)

	expected := []string{
		"http://foo.com",
		"http://foo.com/bar",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}