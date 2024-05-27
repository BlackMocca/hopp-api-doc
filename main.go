package main

import (
	"context"
	"os"
	"strings"

	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/handler"
	myMiddL "github.com/Blackmocca/hopp-api-doc/domain/middleware"
	"github.com/Blackmocca/hopp-api-doc/domain/repository"
	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func registerRoute(e *echo.Echo, handler handler.HttpHandler) {
	e.GET("/", handler.Index, myMiddL.AuthSession(false))
	e.GET("/login", handler.Login, myMiddL.AuthSession(true))

	e.GET("/team/collections", handler.TeamCollection, myMiddL.AuthSession(false))
	e.GET("/my/collection/:user_id", handler.MyCollection, myMiddL.AuthSession(false))
	e.GET("/download/:collection_id", handler.Download, myMiddL.AuthSession(false))

	e.Static("/assets", "public/assets")
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

func connectDatabase(ctx context.Context, databaseURL string) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
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
	defer conn.Close()

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
	renderer := handler.NewTemplateRenderer("./public/*.html")
	e.Renderer = renderer

	httpHandler := handler.NewHttpHandler(repository)
	registerRoute(e, httpHandler)
	e.Logger.Fatal(e.Start(":3000"))

}
