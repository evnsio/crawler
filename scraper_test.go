package main

import (
    "testing"
    "reflect"
)

const test_domain = "http://test.com"

func TestCrawlReturnsNothingForEmptyPage(t *testing.T) {
    page := ""
    urls := Crawl(test_domain, page)

    if len(urls) > 0 {
        t.Errorf("Expected no urls. Got %d", len(urls))
    }
}


func TestCrawlThrowsExceptionForMalformedHTML(t *testing.T) {
    page := "3@2££1t1"
    urls := Crawl(test_domain, page)

    if len(urls) > 0 {
        t.Errorf("Expected no urls. Got %d", len(urls))
    }
}

func TestExtractURLReturnsNothingWhenNoLinks(t *testing.T) {
    page := "<html><body></body></html>"
    urls := Crawl(test_domain, page)

    if len(urls) > 0 {
        t.Errorf("Expected no urls. Got %d", len(urls))
    }
}

func TestCrawlReturnSingleValidURL(t *testing.T) {
    page := "<html><body><a href=\"/foo\">Valid Link</a></body></html>"
    urls := Crawl(test_domain, page)

    expected := []string{"http://test.com/foo"}

    if !reflect.DeepEqual(urls , expected) {
        t.Errorf("%v != %v", expected, urls)
    }
}


func TestCrawlReturnSingleValidURLRelative(t *testing.T) {
    page := "<html><body><a href=\"bar/baz/../foo\">Valid Link</a></body></html>"
    urls := Crawl(test_domain, page)

    expected := []string{"http://test.com/bar/foo"}

    if !reflect.DeepEqual(urls , expected) {
        t.Errorf("%v != %v", expected, urls)
    }
}


func TestCrawlReturnSingleValidURLExtraAttr(t *testing.T) {
    page := "<html><body><a extraattr=\"hello\" href=\"/foo\">Valid Link</a></body></html>"
    urls := Crawl(test_domain, page)

    expected := []string{"http://test.com/foo"}

    if !reflect.DeepEqual(urls , expected) {
        t.Errorf("%v != %v", expected, urls)
    }
}

func TestCrawlReturnsNothingForExternalURL(t *testing.T) {
    page := "<html><body>" +
        "<a href=\"http://www.google.com\">Valid Link</a>" +
        "</body></html>"

    urls := Crawl(test_domain, page)

    if len(urls) > 0 {
        t.Errorf("Expected no urls.  Got %v", urls)
    }
}

func TestCrawlReturnsSingleValidURLForMixed(t *testing.T) {
    page := "<html><body>" +
        "<a href=\"http://www.google.com\">Valid Link</a>" +
        "<a href=\"/bar\">Valid Link</a>" +
        "</body></html>"

    urls := Crawl(test_domain, page)

    expected := []string{"http://test.com/bar"}

    if !reflect.DeepEqual(urls , expected) {
        t.Errorf("%v != %v", expected, urls)
    }
}


func TestCrawlReturnsURLsMatchingSourceHost(t *testing.T) {
    page := "<html><body>" +
        "<a href=\"http://www.google.com\">Valid Link</a>" +
        "<a href=\"http://test.com/foo\">Valid Link</a>" +
        "<a href=\"/bar\">Valid Link</a>" +
        "</body></html>"

    urls := Crawl(test_domain, page)

    expected := []string{"http://test.com/foo", "http://test.com/bar"}

    if !reflect.DeepEqual(urls , expected) {
        t.Errorf("%v != %v", expected, urls)
    }
}
