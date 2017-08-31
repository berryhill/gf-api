package server

import (
	"strconv"
	"net/http"

	"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {

	page := 1
	per_page := 20
	for key, value := range c.QueryParams() {
		if key == "page" {
			page, _ = strconv.Atoi(value[0])
			delete(c.QueryParams(), "page")
 		} else if key == "per-page" {
			per_page, _ = strconv.Atoi(value[0])
			delete(c.QueryParams(), "per-page")
		}
	}

	product := c.Param("product")

	products, _ := models.GetProducts(product, c.QueryParams(), page, per_page)
	var products_paginated []*models.Product
	for k := 0; k < per_page; k++ {
		products_paginated = append(
			products_paginated, (products[(((k+1)*page)-1)]))
	}

	metadata := make(map[string]interface{})
	metadata["page"] = page
	metadata["per_page"] = per_page
	metadata["count"] = len(products)

	response := make(map[string]interface{})
	response["metadata"] = metadata
	response["results"] = products_paginated

	return c.JSON(http.StatusOK, response)
}
