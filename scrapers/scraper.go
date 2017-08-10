package scrapers

import (
	"github.com/berryhill/web-scrapper/models"
	"github.com/PuerkitoBio/goquery"
)


// simulating a db collection
var products_store = map[int]string {
	1: "flyrods",
}

type Scraper interface {
	getUrls() (map[string]string, error)
	getBrand(item *goquery.Selection) (string, error)
	getName(item *goquery.Selection) (string, error)
	getTitle(item *goquery.Selection) (string, error)
	getPrice(item *goquery.Selection) (string, error)
	getUrl(item *goquery.Selection) (string, error)
	getImg(item *goquery.Selection) (string, error)
	getDetails(item *goquery.Selection) ([]string, error)
  	Scrape() (products []*models.Product, errs []error)
}
