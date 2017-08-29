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
	item *goquery.Selection) (name string, err error) {

	name = item.Find(".ui-pl-name-title").Text()

	return name, nil
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

		return bc.Retailer.BaseUrl + href, nil
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

func (bc *BackcountryScraper) getDetails(
	item *goquery.Selection) (details []string, err error) {

	name := item.Find(".ui-pl-name-title").Text()
	string_array := strings.Split(name, " - ")

	if len(string_array) > 1 {
		details = append(details, string_array[1])
	}

	return details, nil
}

func (bc *BackcountryScraper) Scrape() (
	//products []*models.Product, errs []error) {
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
			selection.Each(func(i int, item *goquery.Selection) {
				product := models.NewProduct()
				product.Type = product_type
				product.Brand, _ = bc.getBrand(item)
				product.Name, _ = bc.getName(item)
				product.Title, _ = bc.getTitle(item)
				product.Price, _ = bc.getPrice(item)
				product.Url, _ = bc.getUrl(item)
				product.Retailer = bc.Retailer.Name

				product.Image, err = bc.getImg(item)
				err = errors.New("New Error")
				if err != nil {
					errs = append(errs, err)
				}

				product.Details, _ = bc.getDetails(item)

				found, _ := product.Handle(
					product.Name, product.Title, product.Brand, product.Url,
					product_type)
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
