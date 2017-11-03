package main

import "fmt"

type Page struct {
	url      string
	depth    int
	scraped  bool
	parent   *Page
	children []*Page
}

func NewPage(page_url string, page_depth int, page_parent *Page) *Page {
	p := &Page{
		url:      page_url,
		depth:    page_depth,
		scraped:  false,
		parent:   page_parent,
		children: make([]*Page, 0),
	}

	return p
}

func (p *Page) toSiteMap() {
	fmt.Print("." + p.url + "\n")
	for index, _ := range p.children {
		fmt.Println("├──", p.children[index].url)
	}

	fmt.Println("")

	for index, _ := range p.children {
		child := p.children[index]
		if child.scraped {
			child.toSiteMap()
		}
	}
}

func (p *Page) toList(urls *[]string) {
	*urls = append(*urls, p.url)

	for index, _ := range p.children {
		child := p.children[index]
		if child.scraped {
			child.toList(urls)
		}
	}

}
