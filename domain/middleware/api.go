package middleware

import (
	"net/http"

	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	"github.com/labstack/echo/v4"
)

func AuthSession(skipIfNotExists bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(constants.COOKIE_SESSION_NAME)
			if err != nil && err != http.ErrNoCookie {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			var sessionUser *models.User
			if cookie == nil && !skipIfNotExists {
				/* render unauthorize */
				return c.Render(http.StatusOK, "index", nil)
			}

			if cookie != nil {
				sessionUser = models.NewUserSession(cookie.Value)
			}

			c.Set("session", sessionUser)
			return next(c)
		}
	}
}
