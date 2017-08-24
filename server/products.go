package server

import (
	"net/http"

	"github.com/berryhill/web-scrapper/models"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {

	product := c.Param("product")

	products, _ := models.GetAllProducts(product)

	return c.JSON(http.StatusOK, products)
}
