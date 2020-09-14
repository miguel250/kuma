package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	Addr     string
	config   *Config
	running  bool
	srv      *http.Server
	stop     chan os.Signal
	start    chan bool
	listener net.Listener
}

func (s *Server) Start() error {
	err := s.listen()
	if err != nil {
		return fmt.Errorf("failed to listen with %w", err)
	}

	go func() {
		s.running = true
		s.start <- true

		log.Printf("Listening on %s\n", s.Addr)

		if err := s.srv.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start server with %v\n", err)
			signal.Stop(s.stop)
			close(s.stop)
			s.running = false
		}
	}()
	return nil
}

func (s *Server) listen() error {
	l, err := net.Listen("tcp", s.srv.Addr)

	if err != nil {
		return err
	}

	s.Addr = "http://" + l.Addr().String()
	s.listener = l
	return nil
}

func (s *Server) StartAndWait() error {
	if !s.running {
		err := s.Start()
		if err != nil {
			return err
		}
	}

	<-s.stop
	return s.Stop()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to stop server with %v", err)
	}
	return nil
}

func New(config *Config, handler http.Handler) *Server {
	if config == nil {
		config = defaultConfig
	}

	setDefault(config)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Addr, config.Port),
		Handler: handler,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	start := make(chan bool, 1)

	return &Server{
		config: config,
		srv:    srv,
		stop:   stop,
		start:  start}
}
