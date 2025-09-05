package request

import (
	"RequestService/internal/domain/model"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) getRequest(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing id")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	req, err := h.uc.GetRequest(c.Request().Context(), id)
	if errors.Is(err, model.ErrRequestNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Request not found")
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, req)
}
