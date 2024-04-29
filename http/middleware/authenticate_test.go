package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Authenticate(t *testing.T) {
	service := GetService()
	service = newMockService(nil)

	t.Run("[POSITIVE] check token", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		h := Authenticate(service)(func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		})

		req.Header.Set(echo.HeaderAuthorization, "Bearer auth-key")
		assert.NoError(t, h(c))
	})

	t.Run("[NEGATIVE] empty header", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		h := Authenticate(service)(func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		})

		req.Header.Set(echo.HeaderAuthorization, "")
		result := h(c).(*echo.HTTPError)
		assert.Equal(t, echo.ErrUnauthorized, result)
	})

	t.Run("[NEGATIVE] error from service", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		tempService := newMockService(errors.New("invalid token"))

		h := Authenticate(tempService)(func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		})

		req.Header.Set(echo.HeaderAuthorization, "")
		result := h(c).(*echo.HTTPError)
		assert.Equal(t, echo.ErrUnauthorized, result)
	})
}
