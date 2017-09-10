package scrapers

import (
	"fmt"
	"errors"
	"bytes"
	"strings"
	"strconv"

	"github.com/berryhill/gf-api/api/models"

	"github.com/PuerkitoBio/goquery"
)

// TODO: Implement error handling
// TODO: Implement logging

type htmlLink struct {
	Text string
	Href string
}

type BackcountryScraper struct {
	Retailer		*models.Retailer
}

func NewBackcountryScraper() *BackcountryScraper {

	bc := new(BackcountryScraper)
	retailer := models.Retailer{}
	bc.Retailer, _ = retailer.Get("backcountry")

	return bc
}

func (bc *BackcountryScraper) getBrand(
	scraped *goquery.Selection) (brand string, err error) {

	brand = scraped.Find(".ui-pl-name-brand.qa-brand-name").Text()

	return brand, nil
}

func (bc *BackcountryScraper) getName(
	scraped *goquery.Selection) (name string, err error) {

	name = scraped.Find(".ui-pl-name-title").Text()
	string_array := strings.Split(name, " ")

	var actual_name bytes.Buffer
	done := false
	for k, str := range string_array {
		if str == "-" {
			done = true
		} else {
			if !done {
				actual_name.WriteString(string_array[k] + " ")
			}
		}
	}

	name = TrimSuffix(actual_name.String(), " ")

	return name, nil
}

func (bc *BackcountryScraper) getTitle(
	scraped *goquery.Selection) (title string, err error) {

	title = scraped.Find(".ui-pl-name-title").Text()

	return title, nil
}

func (bc *BackcountryScraper) getPrice(
	scraped *goquery.Selection) (price string, err error) {

	price_html := scraped.Find(".ui-pl-pricing-low-price")
	if price_html.Text() != "" {
		price = price_html.Text()
	} else {
		price_html = scraped.Find(".ui-pl-pricing-high-price")
		price = price_html.Text()
	}

	return price, nil
}

func (bc *BackcountryScraper) getUrl(
	scraped *goquery.Selection) (href string, err error) {

	var ok bool
	scraped.Find("a").Each(func(_ int, link *goquery.Selection) {
		href, ok = link.Attr("href")
	})
	if ok {

		return bc.Retailer.BaseUrl + href, nil
	}

	return "", errors.New("Url not found")
}

func (bc *BackcountryScraper) getImg(
	scraped *goquery.Selection) (img string, err error) {

	var ok bool
	scraped.Find("img").Each(func(_ int, link *goquery.Selection) {
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

func (bc *BackcountryScraper) getDetails(
	scraped *goquery.Selection) (details []string, err error) {

	name := scraped.Find(".ui-pl-name-title").Text()
	string_array := strings.Split(name, " - ")

	if len(string_array) > 1 {
		details = append(details, string_array[1])
	}

	return details, nil
}

func (bc *BackcountryScraper) Scrape() (
	response map[string]int, errs []error) {

	item_count := 0
	item_added := 0

	var err error
	for product_type, url := range bc.Retailer.Products {

		doc, _ := goquery.NewDocument(
			bc.Retailer.BaseUrl + url)

		// TODO: Refactor pagination into its own method
		// TODO: Refactor to be entirely dynamic pulling data from DB

		pagination := doc.Find(".pag")
		var total_pages string
		pagination.Find(
			"li").Each(func(i int, item *goquery.Selection) {
			if item.Text() != "Next Page" {
				total_pages = item.Text()
			}
		})

		for k := 0; k <= len(total_pages); k++ {
			if k != 0 {
				url := strings.Join(
					[]string{url + "?page=",
						strconv.Itoa(k)}, "")
				doc, _ = goquery.NewDocument(bc.Retailer.BaseUrl + url)
			}
			selection := doc.Find(".product")
			selection.Each(func(i int, scraped *goquery.Selection) {
				item := models.NewItem()
				item.Type = product_type
				item.Brand, _ = bc.getBrand(scraped)
				item.Name, _ = bc.getName(scraped)
				item.Title, _ = bc.getTitle(scraped)
				item.Price, _ = bc.getPrice(scraped)
				item.Url, _ = bc.getUrl(scraped)
				item.Retailer = bc.Retailer.Name

				item.Image, err = bc.getImg(scraped)
				err = errors.New("New Error")
				if err != nil {
					errs = append(errs, err)
				}

				item.Details, _ = bc.getDetails(scraped)

				found, _ := item.Handle(
					item.Name, item.Title, item.Brand, item.Url,
					"items")
				if found {
					//products = append(products, product)
					item_added++
				}

				item_count++
			})
		}
	}

	response = make(map[string]int)
	response["items_found"] = item_count
	response["items_added"] = item_added

	fmt.Println("Items found: ", item_count)
	fmt.Println("Items added: ", item_added)

	return response, errs
}
