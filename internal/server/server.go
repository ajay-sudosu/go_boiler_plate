package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"abc/internal/config"
	"abc/internal/database"
	"abc/internal/di"
	"abc/internal/routes"

	_ "abc/docs" // import the docs generated by swag

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitServer() error {
	cfg := config.Load()
	db, err := database.ConnectMongo(cfg.MongoURI, cfg.DBName)
	if err != nil {
		return err
	}

	container, err := di.NewContainer(db)
	if err != nil {
		return err
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api/v1")
	routes.RegisterAllRoutes(api, container)
	// routes.RegisterUserRoutes(api, container.UserHandler)
	// routes.RegisterProductRoutes(api, container.ProductHandler)

	go func() {
		fmt.Println("Running on port :" + cfg.Port)
		e.Start(":" + cfg.Port)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Shutting down...")
	return e.Shutdown(ctx)
}
