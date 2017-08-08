package scrapers

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func CabelasScraper() {
	url := "http://www.cabelas.com/catalog/browse/_/N-1102568" +
		"?CQ_view=list&CQ_ztype=GNP&CQ_pagesize=100"
	doc, _ := goquery.NewDocument(url)
	selection := doc.Find(".productItem")
	selection.Each(func(i int, item *goquery.Selection) {
		fmt.Println(i)

		price := item.Find(".itemPrice")
		class, _ := price.Attr("class")
		fmt.Println(class, price.Text())
	})
}
