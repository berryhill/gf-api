package server

import (
	"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
	"net/http"
)


func CreateRetailer(c echo.Context) error {

	retailer := &models.Retailer{}

	if err := c.Bind(retailer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := retailer.Create(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, retailer)
}


