package asynccrawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func fetchPage(page_url string) string {
	response, err := http.Get(page_url)

	if err != nil {
		// should probably handle these
		return ""
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
