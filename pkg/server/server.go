package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/opoccomaxao/wblitz-watcher/pkg/middleware"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

type Server struct {
	config Config
	engine *gin.Engine
	server *http.Server
}

type Config struct {
	Port string `env:"PORT,required" envDefault:"8080"`
}

func New(
	config Config,
	_ *telemetry.Service,
) *Server {
	res := &Server{
		config: config,
	}

	res.engine = res.initGin()

	res.server = &http.Server{
		Addr:              ":" + config.Port,
		Handler:           res.engine,
		ReadHeaderTimeout: time.Minute,
		ReadTimeout:       time.Minute,
	}

	return res
}

func (s *Server) initGin() *gin.Engine {
	res := gin.Default()
	res.Use(
		middleware.CORS(),
		middleware.TelemetryInit(),
		middleware.TelemetryExtra(),
		middleware.HandleErrors(),
		middleware.HandlePanic(),
	)

	return res
}

func (s *Server) Router() gin.IRouter {
	return s.engine
}

// Serve is blocking call.
func (s *Server) Serve() error {
	//nolint:wrapcheck
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	//nolint:wrapcheck
	return s.server.Shutdown(context.Background())
}
