package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antchfx/htmlquery"
	"github.com/heying-haevg/linkchecker"
)

func main() {
	// 1. find all subpages

	url := "https://bitfieldconsulting.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	
	body := resp.Body
	// io.Copy(os.Stdout, body)
	doc, err := htmlquery.Parse(body)
	if err != nil {
		log.Fatal(err)
	}
	list := htmlquery.Find(doc, "//a/@href")
	for _, n := range list {
		link := htmlquery.SelectAttr(n, "href")

		statusCode, err := linkchecker.GetStatus(link)
		if err != nil {
			// fmt.Printf("not a valid absolute link: %v. %v", link)
			continue
		}

		if statusCode != http.StatusOK {
			fmt.Printf("%v got status %v\n", link, statusCode)
		}
	}

	// 2. find every link and check its status
	// parse html

	// home work: more tests, follow working links

}
