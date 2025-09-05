package request

import (
	"RequestService/internal/transport/http/request/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) createRequest(c echo.Context) error {
	var req dto.Request
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := h.uc.CreateRequest(
		c.Request().Context(),
		req.Domain(),
		req.UserID,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]int64{
		"ID": id,
	})
}
