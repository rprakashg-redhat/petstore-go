package http

import (
	"net/http"
	"os"
	"time"

	"github.com/rprakashg-redhat/petstore-go/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// Server base http server
type Server interface {
	Start(logger log.Logger)
}

type server struct {
	http.Server
	tls               bool
	certFile, keyFile string
}

// NewServer creates a new instance of http server
func NewServer(addr string, tls bool, certFile, keyFile string, routes Routes) Server {
	handler := newRouter(routes)

	return &server{
		Server: http.Server{
			Addr:    addr,
			Handler: handler,
		},
		tls:      tls,
		certFile: certFile,
		keyFile:  keyFile,
	}
}

// Start starts http server with graceful shutdown
func (s *server) Start(logger log.Logger) {
	go routing.GracefulShutdown(
		&s.Server,
		10*time.Second,
		logger.Infof,
	)
	if s.tls {
		if err := s.ListenAndServeTLS(s.certFile, s.keyFile); err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
	} else {
		if err := s.ListenAndServe(); err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
	}
	logger.Infof("server is running at %v", s.Addr)
}
