package http

import (
	"common-utils-go/http/middleware"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
)

var r *Router

func GetRouter() *Router {
	if r == nil {
		r = &Router{
			echo.New(),
		}
	}
	r.HTTPErrorHandler = middleware.DefaultErrorHandler
	r.Echo.Use(middleware2.Logger())
	return r
}

type Router struct {
	*echo.Echo
}

func (r *Router) Start() error {
	address := fmt.Sprintf("0.0.0.0:%v", GetConfig().Port)
	config := GetConfig()
	if config.UseTLS {
		return r.StartAutoTLS(address)
	}
	return r.Echo.Start(address)
}

func (r *Router) Close(ctx context.Context) error {
	return r.Shutdown(ctx)
}

func (r *Router) Use(options ...Option) {
	for _, option := range options {
		option(r.Echo)
	}
}

type Option func(router *echo.Echo)
