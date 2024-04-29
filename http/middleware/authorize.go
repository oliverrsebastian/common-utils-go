package middleware

import "github.com/labstack/echo/v4"

func Authorize(service Service, code ResourceCode) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			rawUser := c.Get(UserCtxKey)
			if rawUser == nil {
				return echo.ErrUnauthorized
			}

			user, ok := rawUser.(*User)
			if !ok {
				return echo.ErrUnauthorized
			}

			if err := service.GetAccess(ctx, user.ID, code); err != nil {
				return echo.ErrForbidden
			}

			return next(c)
		}
	}
}
