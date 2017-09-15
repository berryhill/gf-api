package server

import (
	"net/http"

	"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
)

func CreateProductType(c echo.Context) error {

	product_type := &models.ProductType{}

	if err := c.Bind(product_type); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := product_type.Create(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, product_type)
}


func GetProductTypes(c echo.Context) error {

	product_types, _ := models.GetProductTypes()

	return c.JSON(http.StatusOK, product_types)
}

