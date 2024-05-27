package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthSession(skipIfNotExists bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("access_token")
			if err != nil && err != http.ErrNoCookie {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			var accessToken string
			if cookie == nil && !skipIfNotExists {
				/* render unauthorize */
				return c.Render(http.StatusOK, "index", map[string]interface{}{
					"is_auth": false,
				})
			}

			if cookie != nil {
				accessToken = cookie.Value
			}

			c.Set("session", accessToken)
			return next(c)
		}
	}
}
