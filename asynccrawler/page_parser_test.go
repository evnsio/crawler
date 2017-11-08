package asynccrawler

import (
	"reflect"
	"testing"
)

var page_data = ""

func mockFetchPage(page_url string) string {
	return page_data
}

const test_domain = "http://foo.com"

func TestExtractorReturnsNothingForEmptyPage(t *testing.T) {
	page_data = ""

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	if len(urls) > 0 {
		t.Errorf("Expected no urls. Got %d", len(urls))
	}
}

func TestExtractorThrowsExceptionForMalformedHTML(t *testing.T) {
	page_data = "3@2££1t1"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	if len(urls) > 0 {
		t.Errorf("Expected no urls. Got %d", len(urls))
	}
}

func TestExtractURLReturnsNothingWhenNoLinks(t *testing.T) {
	page_data = "<html><body></body></html>"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	if len(urls) > 0 {
		t.Errorf("Expected no urls. Got %d", len(urls))
	}
}

func TestExtractorReturnSingleValidURL(t *testing.T) {
	page_data = "<html><body><a href=\"/foo\">Valid Link</a></body></html>"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	expected := []string{"http://foo.com/foo"}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("%v != %v", expected, urls)
	}
}

func TestExtractorReturnSingleValidURLRelative(t *testing.T) {
	page_data = "<html><body><a href=\"bar/baz/../foo\">Valid Link</a></body></html>"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	expected := []string{"http://foo.com/bar/foo"}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("%v != %v", expected, urls)
	}
}

func TestExtractorReturnSingleValidURLExtraAttr(t *testing.T) {
	page_data = "<html><body><a extraattr=\"hello\" href=\"/foo\">Valid Link</a></body></html>"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	expected := []string{"http://foo.com/foo"}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("%v != %v", expected, urls)
	}
}

func TestExtractorReturnsNothingForExternalURL(t *testing.T) {
	page_data = "<html><body>" +
		"<a href=\"http://www.google.com\">Valid Link</a>" +
		"</body></html>"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	if len(urls) > 0 {
		t.Errorf("Expected no urls.  Got %v", urls)
	}
}

func TestExtractorReturnsSingleValidURLForMixed(t *testing.T) {
	page_data = "<html><body>" +
		"<a href=\"http://www.google.com\">Valid Link</a>" +
		"<a href=\"/bar\">Valid Link</a>" +
		"</body></html>"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	expected := []string{"http://foo.com/bar"}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("%v != %v", expected, urls)
	}
}

func TestExtractorReturnsURLsMatchingSourceHost(t *testing.T) {
	page_data = "<html><body>" +
		"<a href=\"http://www.google.com\">Valid Link</a>" +
		"<a href=\"http://foo.com/foo\">Valid Link</a>" +
		"<a href=\"/bar\">Valid Link</a>" +
		"</body></html>"

	pageParser := NewPageParser(mockFetchPage)
	urls := pageParser.extractURLs(test_domain)

	expected := []string{"http://foo.com/foo", "http://foo.com/bar"}

	if !reflect.DeepEqual(urls, expected) {
		t.Errorf("%v != %v", expected, urls)
	}
}
