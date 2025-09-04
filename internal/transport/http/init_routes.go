package http

import "github.com/labstack/echo/v4"

func (h *Handler) InitRoutes(e *echo.Echo) error {
	e.GET("/request/:id", h.getRequest)
	e.POST("/request", h.createRequest)

	return nil
}
