package server

import (
	"net/http"

	"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
)

func GetProductTypes(c echo.Context) error {

	product_types, _ := models.GetProductTypes()

	return c.JSON(http.StatusOK, product_types)
}

