package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Code    int
	Status  string
	Message string
}

func DefaultErrorHandler(err error, c echo.Context) {
	var (
		msg string

		code = http.StatusInternalServerError
		resp = &ErrorResponse{}
	)
	switch parsed := err.(type) {
	case *echo.HTTPError:
		code = parsed.Code
		resp.Code = code
		resp.Status = http.StatusText(code)
		resp.Message = parsed.Message.(string)
		switch message := parsed.Message.(type) {
		case string:
			msg = message
			resp.Message = msg
		}
	}
	c.JSON(code, resp)
}
