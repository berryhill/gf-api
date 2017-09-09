//func CabelasScraper() {
//	url := "http://www.cabelas.com/catalog/browse/_/N-1102568" +
//		"?CQ_view=list&CQ_ztype=GNP&CQ_pagesize=100"
//	doc, _ := goquery.NewDocument(url)
//	selection := doc.Find(".productItem")
//	selection.Each(func(i int, item *goquery.Selection) {
//		fmt.Println(i)
//
//		price := item.Find(".itemPrice")
//		class, _ := price.Attr("class")
//		fmt.Println(class, price.Text())
//	})
//}

package scrapers

import (
	"fmt"
	"errors"
	"strings"
	"strconv"

	"github.com/berryhill/gf-api/api/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
)

// TODO: Implement error handling
// TODO: Implement logging

type CabelasScraper struct {
	Retailer		*models.Retailer
}

func NewCabelasScraper() *CabelasScraper {

	cb := new(CabelasScraper)
	retailer := models.Retailer{}
	cb.Retailer, _ = retailer.Get("cabelas")

	return cb
}

func (cb *CabelasScraper) getBrand(
	item *goquery.Selection) (brand string, err error) {

	title := item.Find(
		".productContentBlock").Find(
		"a").Find("h3").Text()

	brand_array := strings.Split(title, " ")

	return brand_array[0], nil
}

func (cb *CabelasScraper) getName(
	item *goquery.Selection) (name string, err error) {

	name = item.Find(
		".productContentBlock").Find(
		"a").Find("h3").Text()

	return name, nil
}

func (cb *CabelasScraper) getTitle(
	item *goquery.Selection) (title string, err error) {

	title = item.Find(
		".productContentBlock").Find(
		"a").Find("h3").Text()

	return title, nil
}

func (cb *CabelasScraper) getPrice(
	item *goquery.Selection) (price string, err error) {

	price = item.Find(
		".pricingBlock").Find(".itemPrice").Text()

	price_formatted := ""
	for _, char := range price {
		if char != 9 {
			price_formatted = price_formatted + string(char)
		}
	}

	return price_formatted, nil
}

func (cb *CabelasScraper) getUrl(
	item *goquery.Selection) (href string, err error) {

	var ok bool
	item.Find(
		".imageBlock").Find(
		"a").Each(func(_ int, link *goquery.Selection) {

		href, ok = link.Attr("href")
	})
	if ok {

		return cb.Retailer.BaseUrl + href, nil
	}

	return "", errors.New("Url not found")
}

func (cb *CabelasScraper) getImg(
	item *goquery.Selection) (img string, err error) {

	var ok bool
	item.Find(
		".imageBlock").Find(
		"a").Find(
		"img").Each(func(_ int, link *goquery.Selection) {

		img , ok = link.Attr("src")
	})
	if ok {

		return img, nil
	}

	return "", errors.New("Image not found")
}

func (cb *CabelasScraper) getDetails(
	item *goquery.Selection) (details []string, err error) {

	name := item.Find(".ui-pl-name-title").Text()
	string_array := strings.Split(name, " - ")

	if len(string_array) > 1 {
		details = append(details, string_array[1])
	}

	return details, nil
}

func (cb *CabelasScraper) Scrape() (response map[string]int, errs []error) {

	item_count := 0
	item_added := 0

	fmt.Println(cb.Retailer.Products)

	//var err error
	for product_type, url := range cb.Retailer.Products {

		doc, _ := goquery.NewDocument(cb.Retailer.Products[product_type])

		// TODO: Refactor pagination into its own method
		// TODO: Refactor to be entirely dynamic pulling data from DB

		pagination := doc.Find(".paginationFilter")
		var page_nums []string
		pagination.Find(
			".entry").Each(func(i int, item *goquery.Selection) {

			page_nums = append(page_nums, item.Text())
		})

		var err error
		total_pages_array := page_nums[len(page_nums) - 4]
		total_pages, err := strconv.Atoi(total_pages_array)
		if err != nil {
			log.Errorf("Could not convert string to int: Total Pages")
		}
		fmt.Println("Total Pages: " + strconv.Itoa(total_pages))

		for k := 0; k <= total_pages; k++ {
			if k != 0 {
				url := strings.Join(
					[]string{url + "&CQ_page=",
						strconv.Itoa(k), "0"}, "")
				doc, _ = goquery.NewDocument(url)
			}
			selection := doc.Find(".productItem")
			selection.Each(func(i int, item *goquery.Selection) {
				product := models.NewProduct()
				product.Type = product_type
				product.Brand, _ = cb.getBrand(item)
				product.Name, _ = cb.getName(item)
				product.Title, _ = cb.getTitle(item)
				product.Price, _ = cb.getPrice(item)
				product.Url, _ = cb.getUrl(item)
				product.Retailer = cb.Retailer.Name

				product.Image, err = cb.getImg(item)
				err = errors.New("New Error")
				if err != nil {
					errs = append(errs, err)
				}

				product.Details, _ = cb.getDetails(item)

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