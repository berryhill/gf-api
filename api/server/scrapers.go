package server

import (
	"net/http"

	"github.com/berryhill/gf-api/api/scrapers"

	"github.com/labstack/echo"
)

var Scrapers []scrapers.Scraper

func SetupScrapers() {

	Scrapers = append(Scrapers, scrapers.NewBackcountryScraper())
}

func ScrapeBackcountry(c echo.Context) error {

	products_added, _ := Scrapers[0].Scrape()

	return c.JSON(http.StatusOK, products_added)
}
