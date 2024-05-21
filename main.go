package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/handler"
	"github.com/Blackmocca/hopp-api-doc/domain/repository"
	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func registerRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		resp := map[string]interface{}{
			"hello": "world",
		}
		return c.JSON(http.StatusOK, resp)
	})

	group := e.Group("/docs")
	group.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			return strings.TrimSuffix(c.Request().URL.Path, "/") == "/docs"
		},
		Root:       "docs",
		Browse:     true,
		IgnoreBase: false,
	}))
}

func connectDatabase(ctx context.Context, databaseURL string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	return conn
}

func main() {
	log.SetLevel(log.INFO)
	var ctx = context.Background()
	var conn = connectDatabase(ctx, constants.DATABASE_URL)
	defer conn.Close(ctx)

	repository := repository.NewPsqlRepository(conn)

	cron, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	jobHandler := handler.NewCronjobHandler(cron, repository)
	jobHandler.Run()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	registerRoute(e)
	e.Logger.Fatal(e.Start(":3000"))

}
