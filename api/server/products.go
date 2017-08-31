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

	// TODO: Fix internal server error when page is out of range

	last_page := len(products) / per_page + 1; var page_count int
	if page == last_page {
		page_count = len(products) % per_page
	} else {
		page_count = per_page
	}
	for k := 0; k < page_count; k++ {
		products_paginated = append(
			products_paginated, (products[k + ((page - 1) * per_page)]))
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
