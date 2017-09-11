package main

import (
	"github.com/berryhill/gf-api/api/db"
	"github.com/berryhill/gf-api/api/server"
    //"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
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

	r := e.Group("/admin")
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	// TODO: Implement CMS endpoints
	e.POST("/login", server.Login)

	r.POST("/product", server.CreateProduct)
	// r.GET("/product/:id", server.GetProduct)
	r.PUT("/product", server.UpdateProduct)
	r.DELETE("/product", server.DeleteProduct)

	// r.GET("/items", server.GetItems)
	// r.GET("/items/:id"", server.GetItem)

	e.POST("/backcountry/scrape", server.ScrapeBackcountry)
	e.POST("/cabelas/scrape", server.ScrapeCabelas)

	e.GET("/product-types", server.GetProductTypes)

	e.GET("/products/:product", server.GetProducts)

	e.Logger.Fatal(e.Start(":8080"))
}

type jwtCustomClaims struct {
	Name  	string 		`json:"name"`
	Admin 	bool   		`json:"admin"`
	jwt.StandardClaims
}
