package scrapers

import (
	//"github.com/berryhill/web-scrapper/models"
	"github.com/PuerkitoBio/goquery"
)


// simulating a db collection
var products_store = map[int]string {
	1: "flyrods",
}

type Scraper interface {
	getBrand(item *goquery.Selection) (string, error)
	getName(item *goquery.Selection) (string, error)
	getTitle(item *goquery.Selection) (string, error)
	getPrice(item *goquery.Selection) (string, error)
	getUrl(item *goquery.Selection) (string, error)
	getImg(item *goquery.Selection) (string, error)
	getDetails(item *goquery.Selection) ([]string, error)
  	//Scrape() (products []*models.Product, errs []error)
  	Scrape() (response map[string]int, errs []error)
}
