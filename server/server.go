package server

import (
	"log"
	"net"

	"go.gabby.network/api/gabby"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Server is a Gabby server, which handles the endpoints and authorization
type Server struct{}

// New returns a new Server
func New() *Server {
	return &Server{}
}

// ListenAndServe runs the server loop
func (s Server) ListenAndServe(addr string) {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen failure: %v", err)
	}
	g := grpc.NewServer()
	gabby.RegisterServerServer(g, s)
	if err := g.Serve(conn); err != nil {
		log.Fatalf("serve failure: %v", err)
	}
}

// Auth handles the gRPC request for Auth
func (s Server) Auth(c context.Context, r *gabby.AuthRequest) (*gabby.AuthResponse, error) {
	return nil, nil
}
