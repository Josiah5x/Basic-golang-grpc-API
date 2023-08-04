package main

import (
	"context"
	"grpc/example/interface/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Defining a struct for grpc server
type Server struct {
	proto.UnimplementedAddServiceServer
}

// implementing Add function generated by grpc compiler located in proto folder dir
func (s *Server) Add(ctx context.Context, r *proto.Request) (*proto.Response, error) {
	a, b := r.GetA(), r.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

// implementing Add function generated by grpc compiler located in proto folder dir
func (s *Server) Multiply(ctx context.Context, r *proto.Request) (*proto.Response, error) {
	a, b := r.GetA(), r.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}

func main() {
	// connection
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// new grpc server
	srv := grpc.NewServer()
	// we need to import Register service struct from proto generated code
	proto.RegisterAddServiceServer(srv, &Server{})
	reflection.Register(srv)
	// Serving the tcp connection with the server
	if err := srv.Serve(conn); err != nil {
		log.Fatalln(err)
	}
}
