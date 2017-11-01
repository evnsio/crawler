package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
    "golang.org/x/net/html"
    "strings"
    "net/url"
    "fmt"
    "strconv"
)


func FetchPage(path string) string {
    response, err := http.Get(path)

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


func GetDomain(page_url string) string {
    u, err := url.Parse(string(page_url))

    if err != nil {
        log.Fatal(err)
    }

    return u.Hostname()
}


func GenerateURL(host string, path string) string {
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

    urls := make([]string, 0)
    if depth > 2 {
        return urls
    }

    url_domain := GetDomain(page_url)
    page_contents := FetchPage(page_url)

    tokenizer := html.NewTokenizer(strings.NewReader(page_contents))
    anchorTag := []byte{'a'}

    for {
        switch tokenizer.Next() {
        case html.ErrorToken:
            return urls

        case html.StartTagToken:
            tag, hasAttr := tokenizer.TagName()
            if hasAttr && bytes.Equal(anchorTag, tag) {
                for  {
                    key, val, more := tokenizer.TagAttr()

                    if string(key) == "href" {
                        u, err := url.Parse(string(val))
                        if err != nil {
                            log.Println(err)
                            continue
                        }

                        // We only want urls matching the source domain or relative paths
                        if u.Hostname() == url_domain || !u.IsAbs() {
                            urls = append(urls, GenerateURL(page_url, string(val)))

                            fmt.Println("base: " + page_url + " -> " + GenerateURL(page_url, string(val)) + "  depth " + strconv.Itoa(depth))

                            Crawl(GenerateURL(page_url, string(val)), depth+1)
                        }
                    }

                    if !more {
                        break
                    }
                }
            }
        }
    }
}