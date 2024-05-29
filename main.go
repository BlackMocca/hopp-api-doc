package main

import (
	"context"
	"os"
	"strings"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/handler"
	myMiddL "github.com/Blackmocca/hopp-api-doc/domain/middleware"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	"github.com/Blackmocca/hopp-api-doc/domain/repository"
	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func registerRoute(e *echo.Echo, handler handler.HttpHandler) {
	// internal app
	e.GET("/", handler.Index, myMiddL.AuthSession(true))
	e.GET("/login", handler.Login, myMiddL.AuthSession(true))

	e.GET("/team/collections", handler.TeamCollection, myMiddL.AuthSession(false))
	e.GET("/my/collection", handler.MyCollection, myMiddL.AuthSession(false))
	e.GET("/download/:collection_id", handler.Download, myMiddL.AuthSession(false))
	e.POST("/import", handler.ImportCollection, myMiddL.AuthSession(false))
	e.DELETE("/my/collection/:collection_id", handler.DeleteMycollection, myMiddL.AuthSession(false))

	// signin provider
	e.GET("/v1/auth/signin", handler.AuthProvider)
	e.GET("/v1/auth/:provider/callback", handler.AuthProviderCallback)

	// signout
	e.POST("/signout", handler.Signout, myMiddL.AuthSession(true))

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

func getAuthMicrosoftRepository() domain.AuthProvider {
	client := models.NewAuthMicrosoftClient(
		constants.MICROSOFT_TENANT,
		constants.MICROSOFT_CLIENT_ID,
		constants.MICROSOFT_CLIENT_SECRET,
		constants.MICROSOFT_CALLBACK_URL,
		constants.MICROSOFT_SCOPE,
	)
	return repository.NewMicrosoftProviderRepository(client, true, 30)
}

func main() {
	log.SetLevel(log.INFO)
	var ctx = context.Background()
	var conn = connectDatabase(ctx, constants.DATABASE_URL)
	defer conn.Close()

	authMicrosoftRepository := getAuthMicrosoftRepository()
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

	httpHandler := handler.NewHttpHandler(repository, authMicrosoftRepository)
	registerRoute(e, httpHandler)
	e.Logger.Fatal(e.Start(":3000"))

}
