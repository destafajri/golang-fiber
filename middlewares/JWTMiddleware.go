package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(c *fiber.Ctx) error {
	HeaderAuthorization := "Authorization"

	// get token value
	tokenValue := c.Get(HeaderAuthorization)

	// parsing authorization bearer
	tokenString := strings.Replace(tokenValue, "Bearer ", "", -1)

	// claims variable
	claims := &JWTClaim{}

	// parsing token jwt
	tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})

	// cek token validation
	err = TokenValidation(c, tokenJwt, err)
	if err != nil {
		return err
	}

	// assign role
	GetRole(claims)
	GetClaims(claims)

	return c.JSON(fiber.Map{"token": tokenJwt})
}
