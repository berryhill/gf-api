package scrapers

import (
	"fmt"
	"errors"
	"bytes"
	"strings"

	"github.com/berryhill/web-scrapper/models"

	"github.com/PuerkitoBio/goquery"
)

// simulating a db collection
var bc_urls_store = map[int]string {
	1: "/fly-rods",
}

type htmlLink struct {
	Text string
	Href string
}

type BackcountryScraper struct {
	BaseUrl			string
	Urls 			map[string]string
	Retailer		string
}

func NewBackcountryScraper() *BackcountryScraper {

	bc := new(BackcountryScraper)
	bc.BaseUrl = "https://www.backcountry.com"
	bc.Retailer = "backcountry"
	bc.Urls, _ = bc.getUrls()

	return bc
}

func (bc *BackcountryScraper) getUrls() (
	urls map[string]string, err error) {

	urls = make(map[string]string)
	for k := range products_store {
		urls[products_store[k]] = bc_urls_store[k]
	}

	return urls, nil
}

func (bc *BackcountryScraper) getBrand(
	item *goquery.Selection) (brand string, err error) {

	brand = item.Find(".ui-pl-name-brand.qa-brand-name").Text()

	return brand, nil
}

func (bc *BackcountryScraper) getName(
	item *goquery.Selection) (name string, err error) {

	name = item.Find(".ui-pl-name-title").Text()
	string_array := strings.Split(name, " ")

	var actual_name bytes.Buffer
	done := false
	for k, str := range string_array {
		if str == "Fly" {
			if string_array[k+1] == "Rod" {
				done = true
			}
		} else {
			if !done {
				actual_name.WriteString(string_array[k] + " ")
			}
		}
	}

	name = actual_name.String()

	return TrimSuffix(name, " "), nil
}

func (bc *BackcountryScraper) getPrice(
	item *goquery.Selection) (price string, err error) {

	price_html := item.Find(".ui-pl-pricing-low-price")
	if price_html.Text() != "" {
		price = price_html.Text()
	} else {
		price_html = item.Find(".ui-pl-pricing-high-price")
		price = price_html.Text()
	}

	return price, nil
}

func (bc *BackcountryScraper) getUrl(
	item *goquery.Selection) (href string, err error) {

	var ok bool
	item.Find("a").Each(func(_ int, link *goquery.Selection) {
		href, ok = link.Attr("href")
	})
	if ok {

		return bc.BaseUrl + href, nil
	}

	return "", errors.New("Url not found")
}

func (bc *BackcountryScraper) getImg(
	item *goquery.Selection) (img string, err error) {

	var ok bool
	item.Find("img").Each(func(_ int, link *goquery.Selection) {
		img, ok = link.Attr("data-src")
		if img == "" {
			img , ok = link.Attr("src")
		}
	})
	if ok {

		return TrimPrefix(img, 2), nil
	}

	return "", errors.New("Image not found")
}

func (bc *BackcountryScraper) Scrape() (
	products []*models.Product, errs []error) {

	item_count := 0
	item_added := 0

	var err error
	for product_type, url := range bc.Urls {
		doc, _ := goquery.NewDocument(bc.BaseUrl + url)

		selection := doc.Find(".product")
		selection.Each(func(i int, item *goquery.Selection) {
			product := models.NewProduct()
			product.Type = product_type
			product.Brand, _ = bc.getBrand(item)
			product.Name, _ = bc.getName(item)
			product.Price, _ = bc.getPrice(item)
			product.Url, _ = bc.getUrl(item)
			product.Retailer = bc.Retailer

			product.Image, err = bc.getImg(item)
			err = errors.New("New Error")
			if err != nil {
				errs = append(errs, err)
			}

			products = append(products, product)
			found, _ := product.Handle(product.Name, product.Brand)
			if found {
				item_added++
			}

			item_count++
		})
	}

	fmt.Println("Items found: ", item_count)
	fmt.Println("Items added: ", item_added)

	return products, errs
}
