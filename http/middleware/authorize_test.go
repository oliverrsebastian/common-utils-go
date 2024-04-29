package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Authorize(t *testing.T) {
	service := GetService()
	service = newMockService(nil)

	t.Run("[POSITIVE] user allowed", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.Set(UserCtxKey, &User{})

		h := Authorize(service, DefaultCode)(func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		})

		assert.NoError(t, h(c))
	})

	t.Run("[NEGATIVE] empty context", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		h := Authorize(service, DefaultCode)(func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		})

		result := h(c).(*echo.HTTPError)
		assert.Equal(t, echo.ErrUnauthorized, result)
	})

	t.Run("[NEGATIVE] error from service", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.Set(UserCtxKey, &User{})

		tempService := newMockService(errors.New("access not allowed"))

		h := Authorize(tempService, DefaultCode)(func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		})

		result := h(c).(*echo.HTTPError)
		assert.Equal(t, echo.ErrForbidden, result)
	})
}
