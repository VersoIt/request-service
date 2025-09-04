package server

import (
	"github.com/labstack/echo/v4"
)

type handler interface {
	InitRoutes(ctx *echo.Echo) error
}
