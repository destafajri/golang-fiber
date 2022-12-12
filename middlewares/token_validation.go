package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func TokenValidation(c *fiber.Ctx, token *jwt.Token, err error) error{
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			// token invalid
			return errors.New("Unauthorized")
		case jwt.ValidationErrorExpired:
			// token expired
			return errors.New("Unauthorized, Token expired!")
		default:
			return errors.New("Unauthorized")
		}
	}

	if !token.Valid {
		return errors.New("Unauthorized")
	}

	return nil
}