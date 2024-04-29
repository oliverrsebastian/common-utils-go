package middleware

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func Authenticate(service Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if authHeader == "" {
				return echo.ErrUnauthorized
			}

			token := strings.TrimSpace(strings.ReplaceAll(authHeader, Bearer, ""))
			user, err := service.Check(ctx, token)
			if err != nil || user == nil {
				return echo.ErrUnauthorized
			}

			c.Set(UserCtxKey, user)
			c.Set(TokenCtxKey, token)

			return next(c)
		}
	}
}
