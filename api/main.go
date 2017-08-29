package main

import (
	"github.com/berryhill/gf-api/api/db"
	"github.com/berryhill/gf-api/api/server"
  // "github.com/berryhill/gf-api/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// TODO: Implement logging

	db.Connect()
	server.SetupScrapers()

	// retailer := models.NewRetailer(
	// 	"backcountry",
	// 	"https://www.backcountry.com",
	// 	"/fly-rods",
	// 	"fly_rods",
	// )
	// retailer.Create()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.POST("/backcountry/scrape", server.ScrapeBackcountry)

	e.GET("/products/:product", server.GetProducts)

	//TODO: Implement search

	e.Logger.Fatal(e.Start(":8080"))
}
