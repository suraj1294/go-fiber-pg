package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suraj1294/go-fiber-pg-auth/cmd/api/response"
	"github.com/suraj1294/go-fiber-pg-auth/internal/models"
)

const authHandlerMatch = "Bearer "

type MiddlewareHandler struct {
	AppStore
}

func (rh *MiddlewareHandler) CustomMiddleware(c *fiber.Ctx) error {
	fmt.Println("custom middle ware")

	return c.Next()
}

func (rh *MiddlewareHandler) VerifyAuth(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if strings.HasPrefix(authHeader, authHandlerMatch) {
		token := strings.TrimPrefix(authHeader, authHandlerMatch)

		if token == "" {
			return response.UnAuthorizedRequestHandler(c, "unauthorized")

		}

		claims := &models.Claims{}

		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(rh.AuthStore.Secret), nil
		})

		if !parsedToken.Valid {
			return response.UnAuthorizedRequestHandler(c, "unauthorized")
		}

		if err != nil {
			return response.UnAuthorizedRequestHandler(c, "unauthorized")
		}

		return c.Next()

	}

	return response.UnAuthorizedRequestHandler(c, "unauthorized")

}
