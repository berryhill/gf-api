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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// CMS Endpoints
	e.POST("/backcountry/scrape", server.ScrapeBackcountry)
	e.POST("/cabelas/scrape", server.ScrapeCabelas)

	e.POST("/retailer", server.CreateRetailer)

	// Frontend Endpoints
	e.GET("/product-types", server.GetProductTypes)

	e.GET("/products/:product", server.GetProducts)

	//TODO: Implement search

	e.Logger.Fatal(e.Start(":8080"))
}
