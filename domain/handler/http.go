package handler

import (
	"fmt"
	"net/http"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	psql domain.Datasource
}

func NewHttpHandler(psql domain.Datasource) HttpHandler {
	return HttpHandler{psql: psql}
}

func (HttpHandler) Index(c echo.Context) error {
	fmt.Println("cookie context", c.Cookies())

	resty.New().SetDebug(true).R().Get("https://hoppscotch-api.innovasive.dev/v1/auth/providers")
	resty.New().SetDebug(true).R().Get("https://hoppscotch-api.innovasive.dev/graphql")

	return c.Render(http.StatusOK, "index.html", nil)
}
