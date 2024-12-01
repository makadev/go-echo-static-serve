package server

import (
	"context"
	"example/go-echo-stuff/webserver/internal/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		echo:   echo.New(),
		config: config,
	}
}

func (srv *Server) Setup() {
	// setup everything we can without introducing circular imports

	// set debug mode
	srv.echo.Debug = srv.config.Server.Debug

	// Configure trusted proxies
	switch srv.config.Server.ProxyMode {
	case "xff":
		srv.echo.IPExtractor = echo.ExtractIPFromXFFHeader(srv.config.Server.GetTrustedProxyOptions()...)
	case "real-ip":
		srv.echo.IPExtractor = echo.ExtractIPFromRealIPHeader(srv.config.Server.GetTrustedProxyOptions()...)
	default:
		srv.echo.IPExtractor = echo.ExtractIPDirect()
	}

}

func (srv *Server) Start() error {
	// Start
	if err := srv.echo.Start(srv.config.Server.Host + ":" + strconv.Itoa(srv.config.Server.Port)); err != nil && err != http.ErrServerClosed {
		srv.echo.Logger.Fatal("shutting down the server with error: ", err)
		return err
	}
	return nil
}

func (srv *Server) GetEcho() *echo.Echo {
	return srv.echo
}

func (srv *Server) GetConfig() *config.Config {
	return srv.config
}

func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.echo.Shutdown(ctx); err != nil {
		srv.echo.Logger.Fatal(err)
		return err
	}
	return nil
}
