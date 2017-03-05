package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://www.supremenewyork.com/shop/all/jackets")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".inner-article").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		fmt.Println(class, s.Text())
	})
}

func main() {
	ExampleScrape()
}
