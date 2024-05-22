package handler

import (
	"net/http"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	psql domain.Datasource
}

func NewHttpHandler(psql domain.Datasource) HttpHandler {
	return HttpHandler{psql: psql}
}

func (HttpHandler) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
