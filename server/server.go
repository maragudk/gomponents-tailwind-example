package server

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

// Server handles HTTP requests.
type Server struct {
	host   string
	log    *log.Logger
	mux    chi.Router
	port   int
	server *http.Server
}

// Options for New.
type Options struct {
	Host string
	Log  *log.Logger
	Port int
}

// New Server with the given Options.
func New(opts Options) *Server {
	if opts.Log == nil {
		opts.Log = log.New(ioutil.Discard, "", 0)
	}
	return &Server{
		host: opts.Host,
		log:  opts.Log,
		mux:  chi.NewMux(),
		port: opts.Port,
	}
}

// Start the Server, connecting to dependencies and listening for HTTP.
func (s *Server) Start() error {
	s.setupRoutes()

	address := net.JoinHostPort(s.host, strconv.Itoa(s.port))
	s.server = &http.Server{
		Addr:              address,
		Handler:           s.mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		ErrorLog:          s.log,
	}

	s.log.Printf("Starting on http://%v\n", address)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
