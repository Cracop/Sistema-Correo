package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "servidor/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.TurboMessageServer
}

// El nombre de la función lo busco en turbomessage_grpc.pb
func (s *server) NuevoUsuario(ctx context.Context, in *pb.Usuario) (*pb.Status, error) {
	fmt.Print(in.Usuario)
	return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"John"}[0]}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTurboMessageServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
