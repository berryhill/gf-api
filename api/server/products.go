package server

import (
	"net/http"

	"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {

	product := c.Param("product")

	products, _ := models.GetProducts(product, c.QueryParams())

	return c.JSON(http.StatusOK, products)
}
