package main

import (
	context "context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	pb "servidor/proto"
	"sync"

	"google.golang.org/grpc"
)

var (
	port            = flag.Int("port", 50051, "The server port")
	filenameUsers   = "db_users.json"
	filenameCorreos = "db_correos.json"
	usersMap        = make(map[string]Usuario)
	usersLock       sync.Mutex
)

type Usuario struct {
	User   string `json:"usuario"`
	Passwd string `json:"passwd"`
}

type Server struct {
	pb.TurboMessageServer
}

// El nombre de la función lo busco en turbomessage_grpc.pb
func (s *Server) NuevoUsuario(ctx context.Context, in *pb.Usuario) (*pb.Status, error) {
	fmt.Print(in.Usuario)
	usersLock.Lock()
	//LIFO
	defer usersLock.Unlock()
	defer reloadDBs()

	if _, exists := usersMap[*in.Usuario]; !exists {
		usersMap[*in.Usuario] = Usuario{Passwd: *in.Contrasena}
		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"usuario creado con éxito"}[0]}, nil
	} else {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Usuario ya existe"}[0]}, nil
	}

}

func (s *Server) RevisarUsuario(ctx context.Context, in *pb.Usuario) (*pb.Status, error) {
	if _, exists := usersMap[*in.Usuario]; exists {
		usersMap[*in.Usuario] = Usuario{Passwd: *in.Contrasena}
		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Usuario existe"}[0]}, nil
	} else {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Usuario no existe"}[0]}, nil
	}

}

func (s *Server) DirectorioUsuario(em *pb.Empty, stream pb.TurboMessage_DirectorioUsuarioServer) error {
	for id, _ := range usersMap {
		// tempUser := Usuario{person.User, strconv.Itoa(id)}
		idP := ""
		tempUser := &pb.Usuario{Usuario: &id, Contrasena: &idP}

		if err := stream.Send(tempUser); err != nil {
			return err
		}
	}
	return nil
}

// func (UnimplementedTurboMessageServer) DirectorioUsuario(*Empty, TurboMessage_DirectorioUsuarioServer) error {
// 	return status.Errorf(codes.Unimplemented, "method DirectorioUsuario not implemented")
// }

func reloadDBs() {

	jsonData, err := json.Marshal(usersMap)
	if err != nil {
		return
	}
	os.WriteFile(filenameUsers, jsonData, 0644)

	data, err := os.ReadFile(filenameUsers)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &usersMap); err != nil {
		return
	}
}

func init() {
	_, err := os.Stat(filenameUsers)
	if os.IsNotExist(err) {
		// File doesn't exist, create it
		file1, err := os.Create(filenameUsers)

		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file1.Close()
		fmt.Println("File created:", filenameUsers)
	}

	_, err = os.Stat(filenameCorreos)
	if os.IsNotExist(err) {
		// File doesn't exist, create it
		file1, err := os.Create(filenameCorreos)

		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file1.Close()
		fmt.Println("File created:", filenameCorreos)
	}

	data, err := os.ReadFile(filenameUsers)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &usersMap); err != nil {
		return
	}
}

// func startServer() {

// }

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTurboMessageServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
