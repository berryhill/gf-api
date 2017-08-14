package main

import (
	"github.com/berryhill/web-scrapper/db"
	"github.com/berryhill/web-scrapper/server"

	"github.com/labstack/echo"
)

func main() {

	// TODO: Implement logging

	db.Connect()
	server.SetupScrapers()

	//retailer := models.NewRetailer(
	//	"backcountry",
	//	"https://www.backcountry.com",
	//	"/fly-rods",
	//	"fly_rods",
	//)
	//retailer.Create()


	e := echo.New()

	e.POST("/backcountry/scrape", server.ScrapeBackcountry)

	e.GET("/products/:product", server.GetProducts)

	e.Logger.Fatal(e.Start(":1323"))
}

