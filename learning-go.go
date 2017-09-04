package main

import (
	"./log"
	"./crawler"
	"github.com/jasonlvhit/gocron"
)

func main() {
	log.Path("./logs")
	
	gocron.Every(10).Seconds().Do(crawler.Crawl)

	<- gocron.Start()
}
