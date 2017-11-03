package main

import (
	"fmt"
	"reflect"
	"testing"
)

func generateTestPage(urls ...string) string {
	const simple_page = "<html><body>%s</body></html>"
	const anchor = "<a href=\"%s\"></a>"

	anchors := ""
	for _, url := range urls {
		anchors += fmt.Sprintf(anchor, url)
	}

	return fmt.Sprintf(simple_page, anchors)
}

func mockSingleFetchPage(page_url string) string {
	return ""
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
	crawler := NewCrawler(
		optionTestPageParser(mockSingleFetchPage),
	)

	urls := crawler.run("http://foo.com", 1)

	expected := []string{"http://foo.com"}
	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerSingleThreadUnlimitedDepth(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPage),
	)

	urls := crawler.run("http://foo.com", 100)

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
	)

	urls := crawler.run("http://foo.com", 0)

	expected := []string{}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerSingleThreadDepthLimitedToOne(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPage),
	)

	urls := crawler.run("http://foo.com", 1)

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
	)

	urls := crawler.run("http://foo.com", 2)

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
	)

	urls := crawler.run("http://foo.com", 100)

	expected := []string{
		"http://foo.com",
		"http://foo.com/bar/a",
		"http://foo.com/bar/b",
		"http://foo.com/bar/baz",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}

func TestCrawlerFetchPageLoop(t *testing.T) {
	crawler := NewCrawler(
		optionTestPageParser(mockMultiFetchPageLoop),
	)

	urls := crawler.run("http://foo.com", 100)

	expected := []string{
		"http://foo.com",
		"http://foo.com/bar",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("\nExpected: %v\nGot: %v", expected, urls)
	}
}
