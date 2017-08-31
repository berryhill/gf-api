package server

import (
	"net/http"

	"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {

	page := 1
	per_page := 20

	product := c.Param("product")

	products, _ := models.GetProducts(product, c.QueryParams())

	metadata := make(map[string]interface{})
	metadata["page"] = page
	metadata["per_page"] = per_page
	metadata["count"] = len(products)

	response := make(map[string]interface{})
	response["metadata"] = metadata
	response["results"] = products

	return c.JSON(http.StatusOK, response)
}
