package main

import (
	"context"
	"example/go-echo-stuff/webserver/internal/config"
	"example/go-echo-stuff/webserver/internal/server"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/sync/errgroup"
)

var server_instance *server.Server

func main() {
	// load config
	config := config.NewConfig()
	config.Load()

	server_instance = server.NewServer(config)

	// setup server
	setup()

	// startup server with graceful shutdown handling
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	group, groupCtx := errgroup.WithContext(mainCtx)
	group.Go(func() error {
		return server_instance.Start()
	})

	group.Go(func() error {
		<-groupCtx.Done()

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server_instance.Shutdown(timeoutCtx)
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("Server shutdown with error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Server shutdown gracefully\n")
	os.Exit(0)
}

func setup() {
	server_instance.Setup()

	e := server_instance.GetEcho()
	// Default Middleware (Request Logger, Recover on Panic)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Serve /app/static
	e.Static("/", server_instance.GetConfig().Server.RootDir)
	// Return 404 otherwise
	e.RouteNotFound("/*", func(c echo.Context) error { return c.NoContent(http.StatusNotFound) })
}
