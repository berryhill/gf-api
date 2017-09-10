package server

import (
	"time"
	"net/http"

	"github.com/berryhill/gf-api/api/models"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)


type jwtCustomClaims struct {
	Name  	string 		`json:"name"`
	Admin 	bool   		`json:"admin"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.Username == "admin" && user.Password == "password" {

		// Set custom claims
		claims := &jwtCustomClaims{
			"Jon Snow",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
