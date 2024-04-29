package http

import (
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
)

var defaultOrigins = []string{"0.0.0.0"}

var DefaultSecurityOption = func(router *echo.Echo) {
	if GetConfig().CSRF {
		csrfConfig := middleware2.DefaultCSRFConfig
		csrfConfig.TokenLookup = "cookie:" + echo.HeaderXCSRFToken
		csrfConfig.CookieHTTPOnly = true
		csrfConfig.CookieSecure = true
		router.Use(middleware2.CSRFWithConfig(csrfConfig))
	}

	if GetConfig().CORS {
		router.Use(middleware2.CORSWithConfig(middleware2.CORSConfig{
			AllowOrigins: defaultOrigins,
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}))
	}

	if GetConfig().Secure {
		router.Use(middleware2.SecureWithConfig(middleware2.SecureConfig{
			XSSProtection:         "1; mode=block",
			ContentTypeNosniff:    "nosniff",
			XFrameOptions:         "SAMEORIGIN",
			HSTSMaxAge:            3600,
			ContentSecurityPolicy: "default-src 'self'",
		}))
	}
}
