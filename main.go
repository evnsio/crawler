package main

import (
    "fmt"
)

func main() {

    /*

      1. Fetch body of site and return as string - return body
      2. Call Crawl on body - return {urls}
      3. Print sitemap - prints string

     */
    root_url := "https://www.monzo.com"
    fmt.Println(Crawl(root_url, 0))

}
