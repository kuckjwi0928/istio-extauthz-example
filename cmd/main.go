package main

import (
	"fmt"
	"istio-extauthz-example/configs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	authorization = "Authorization"
	bearer        = "Bearer "
)

type ExtAuthzServer struct {
	httpServer *http.Server
}

func (s *ExtAuthzServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	authHeader := request.Header.Get(authorization)
	if authHeader == "" {
		log.Printf("No authorization header found")
		response.WriteHeader(http.StatusForbidden)
		return
	}
	token := s.extractToken(authHeader)
	log.Printf("Token: %s", token)
	// TODO: implement your own token validation logic here
}

func (s *ExtAuthzServer) extractToken(authHeader string) string {
	return authHeader[len(bearer):]
}

func (s *ExtAuthzServer) startHTTP(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to create HTTP server: %v", err)
	}

	s.httpServer = &http.Server{Handler: s}

	log.Printf("Starting HTTP server at %s", listener.Addr())
	if err := s.httpServer.Serve(listener); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func (s *ExtAuthzServer) stopHTTP() {
	log.Printf("Stopping HTTP server")
	if err := s.httpServer.Close(); err != nil {
		log.Fatalf("Failed to stop HTTP server: %v", err)
	}
}

func NewExtAuthzServer() *ExtAuthzServer {
	return &ExtAuthzServer{}
}

func main() {
	server := NewExtAuthzServer()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go server.startHTTP(fmt.Sprintf(":%d", configs.Cfg.Port))

	<-done
	server.stopHTTP()
}
