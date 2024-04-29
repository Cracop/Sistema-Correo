package main

import (
	context "context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	pb "servidor/proto"
	"strconv"
	"sync"

	"google.golang.org/grpc"
)

var (
	port            = flag.Int("port", 50051, "The server port")
	filenameUsers   = "db_users.json"
	filenameCorreos = "db_correos.json"
	usersMap        = make(map[int]Usuario)
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
	var found = false
	id := rand.Intn(101) + 1
	// Check if the map is empty
	if len(usersMap) == 0 {

		usersMap[id] = Usuario{User: *in.Usuario, Passwd: *in.Contrasena}
		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{strconv.Itoa(id)}[0]}, nil
	} else {
		for _, person := range usersMap {
			if person.User == *in.Usuario {
				found = true
				break
			}

		}
		if found {
			return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"0"}[0]}, nil
		} else {
			for {
				id = rand.Intn(101) + 1
				if _, exists := usersMap[id]; !exists {
					// If the id does not exist in the map, break out of the loop
					break
				}
			}
			usersMap[id] = Usuario{User: *in.Usuario, Passwd: *in.Contrasena}
			return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{strconv.Itoa(id)}[0]}, nil
		}

	}

}

func (s *Server) RevisarUsuario(ctx context.Context, in *pb.Usuario) (*pb.Status, error) {
	var exito bool
	mensaje := "0"
	if len(usersMap) == 0 {
		exito = false
	} else {
		for id, person := range usersMap {
			if person.User == *in.Usuario && person.Passwd == *in.Contrasena {
				exito = true
				mensaje = strconv.Itoa(id)
				break
			}
		}
	}

	return &pb.Status{Success: &[]bool{exito}[0], Mensaje: &[]string{mensaje}[0]}, nil
}

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
