package main

import (
	"github.com/berryhill/gf-api/api/db"
	"github.com/berryhill/gf-api/api/server"
    //"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// TODO: Implement logging

	db.Connect()
	server.SetupScrapers()

	 //retailer := models.NewRetailer(
	 //	"cabelas",
	 //	"https://www.cabelas.com",
	 //	"http://www.cabelas.com/catalog/browse/_/" +
		//	"N-1104841?CQ_view=list&CQ_ztype=GNP&CQ_pagesize=40",
	 //	"fly_rods",
	 //)
	 //retailer.Create()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.POST("/backcountry/scrape", server.ScrapeBackcountry)
	e.POST("/cabelas/scrape", server.ScrapeCabelas)

	e.GET("/product-types", server.GetProductTypes)

	e.GET("/products/:product", server.GetProducts)

	e.Logger.Fatal(e.Start(":8080"))
}
