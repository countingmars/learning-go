package crawler

import (
	"../log"
	"github.com/PuerkitoBio/goquery"
)
func Crawl() {
	doc, err := goquery.NewDocument("https://blog.golang.org")
	if err != nil {
		log.Fatal(err)
	}

	titles := make([]string, 10)
	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3").Text()
		titles = append(titles, title)
	})
	log.Info(titles)	
	log.Flush()
}
