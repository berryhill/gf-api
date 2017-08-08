package main

import (
	"github.com/berryhill/web-scrapper/scrapers"
	"github.com/berryhill/web-scrapper/db"
)

var Scrapers []scrapers.Scraper

func main() {
	db.Connect()

	Scrapers = append(Scrapers, scrapers.NewBackcountryScraper())
	Scrapers[0].Scrape()
}

//func init(){
//	Scrapers = append(Scrapers, scrapers.NewBackcountryScraper())
//}
