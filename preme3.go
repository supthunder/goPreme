package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://www.supremenewyork.com/shop/all/jackets")
	if err != nil {
		log.Fatal(err)
	}
	var next bool = false
	doc.Find(".name-link").Each(func(i int, s *goquery.Selection) {
		var item bool = strings.Contains(s.Text(), "Contrast Stitch")
		if next == true {
			var color bool = strings.Contains(s.Text(), "Acid Green")
			if color == true {
				fmt.Print("Found: ", s.Text(), "\n")
				str, exists := s.Attr("href")
				if exists {
					fmt.Print("Link: http://www.supremenewyork.com", str, "\n")
				}
				next = false
			}
		}

		if item == true {
			next = true
			// fmt.Print("Found: ", s.Text(), "\n")
			// str, exists := s.Attr("href")
			// if exists {
			// 	fmt.Print("Link: http://www.supremenewyork.com", str, "\n")
			// }
		}

		// str, exists := s.Attr("href")
		// if exists {
		// 	fmt.Print(str)
		// }
		// fmt.Printf("%d %s - %s \n", i+1, title, str)
	})
	// Find the review items
	// doc.Find(".inner-article").Each(func(i int, s *goquery.Selection) {
	// 	class, _ := s.Attr("class")
	// 	// fmt.Println(class, s.Text())
	// 	var item bool = strings.Contains(s.Text(), "Supreme®/Schott® Leather")
	// 	if item == true {
	// 		fmt.Print("Found!", s.Text())
	// 	}
	// })

}

func main() {
	ExampleScrape()
}
