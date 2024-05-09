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
	"sync"

	"google.golang.org/grpc"
)

var (
	port            = flag.Int("port", 50051, "The server port")
	filenameUsers   = "db_users.json"
	filenameCorreos = "db_correos.json"
	usersMap        = make(map[string]Usuario)
	correosMap      = make(map[int]Correo)
	usersLock       sync.Mutex
	correosLock     sync.Mutex
	numMax          = 5
)

type Usuario struct {
	User            string `json:"usuario"`
	Passwd          string `json:"passwd"`
	BandejaEntradas []int  `json:"bandEntrada"`
	BandejaSalidas  []int  `json:"bandSalida"`
}

type Correo struct {
	// User         string `json:"usuario"`
	Tema         string `json:"tema"`
	Destinatario string `json:"destinatario"`
	Emisor       string `json:"emisor"`
	Contenido    string `json:"contenido"`
	Leido        bool   `json:"leido"`
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
	defer reloadUserDBs()

	if _, exists := usersMap[*in.Usuario]; !exists {
		usersMap[*in.Usuario] = Usuario{Passwd: *in.Contrasena}
		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"usuario creado con éxito"}[0]}, nil
	} else {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Usuario ya existe"}[0]}, nil
	}

}

func (s *Server) RevisarUsuario(ctx context.Context, in *pb.Usuario) (*pb.Status, error) {
	if _, exists := usersMap[*in.Usuario]; exists {
		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Usuario existe"}[0]}, nil
	} else {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Usuario no existe"}[0]}, nil
	}

}

func (s *Server) DirectorioUsuario(em *pb.Empty, stream pb.TurboMessage_DirectorioUsuarioServer) error {
	for id := range usersMap {
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

func (s *Server) EnviarCorreo(ctx context.Context, in *pb.Correo) (*pb.Status, error) {

	correosLock.Lock()
	usersLock.Lock()
	//LIFO
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	defer reloadCorreoDBs()
	defer reloadUserDBs()

	if _, exists := usersMap[*in.Destinatario]; !exists {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"No existe tal usuario"}[0]}, nil
	}

	//Probar si la bandeja de salida del emisor está llena
	if len(usersMap[*in.Emisor].BandejaSalidas) >= numMax {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Bandeja de Salida llena"}[0]}, nil
	}
	//Probar si la bandeja de entrada del destinatario está llenaz
	if len(usersMap[*in.Destinatario].BandejaEntradas) >= numMax {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Bandeja de Entrada llena"}[0]}, nil
	}
	// var found = false
	id := rand.Intn(101) + 1

	if _, exists := correosMap[id]; exists {

		for {
			id = rand.Intn(101) + 1
			if _, exists := correosMap[id]; !exists {
				// If the id does not exist in the map, break out of the loop
				break
			}

		}

	}

	correosMap[id] = Correo{Tema: *in.Tema,
		Destinatario: *in.Destinatario,
		Emisor:       *in.Emisor,
		Contenido:    *in.Contenido,
		Leido:        *in.Leido}

	usuario := usersMap[*in.Emisor]
	usuario.BandejaSalidas = append(usuario.BandejaSalidas, id)
	usersMap[*in.Emisor] = usuario

	usuario = usersMap[*in.Destinatario]
	usuario.BandejaEntradas = append(usuario.BandejaEntradas, id)
	usersMap[*in.Destinatario] = usuario

	return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Correo enviado con éxito"}[0]}, nil
}

// func revisarRestriccion(emisor string, destinatario string) (bool, string) {

// }

func (s *Server) CorreosEntrada(in *pb.Usuario, stream pb.TurboMessage_CorreosEntradaServer) error {
	correosLock.Lock()
	usersLock.Lock()
	//LIFO
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	defer reloadCorreoDBs()
	defer reloadUserDBs()

	bandeja := usersMap[*in.Usuario].BandejaEntradas

	for _, id := range bandeja {
		// fmt.Println(id)
		// fmt.Println(bandeja)
		// fmt.Println(correosMap)
		correo := correosMap[id]
		ID := (int32(id))
		// fmt.Println(id)

		tempCorreo := &pb.Correo{Identificador: &ID,
			Tema:         &correo.Tema,
			Destinatario: &correo.Destinatario,
			Emisor:       &correo.Emisor,
			Contenido:    &correo.Contenido,
			Leido:        &correo.Leido}

		if err := stream.Send(tempCorreo); err != nil {
			return err
		}
	}
	return nil

}

func (s *Server) CorreosSalida(in *pb.Usuario, stream pb.TurboMessage_CorreosSalidaServer) error {
	correosLock.Lock()
	usersLock.Lock()
	//LIFO
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	defer reloadCorreoDBs()
	defer reloadUserDBs()

	bandeja := usersMap[*in.Usuario].BandejaSalidas

	for _, id := range bandeja {
		// fmt.Println(id)
		// fmt.Println(bandeja)
		// fmt.Println(correosMap)
		correo := correosMap[id]
		ID := (int32(id))
		// fmt.Println(id)

		tempCorreo := &pb.Correo{Identificador: &ID,
			Tema:         &correo.Tema,
			Destinatario: &correo.Destinatario,
			Emisor:       &correo.Emisor,
			Contenido:    &correo.Contenido,
			Leido:        &correo.Leido}

		if err := stream.Send(tempCorreo); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) CorreoLeido(ctx context.Context, in *pb.Correo) (*pb.Status, error) {

	correosLock.Lock()
	//LIFO
	defer correosLock.Unlock()
	defer reloadCorreoDBs()

	if _, exists := correosMap[int(*in.Identificador)]; exists {

		correo := correosMap[int(*in.Identificador)]
		correo.Leido = true
		correosMap[int(*in.Identificador)] = correo

		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Correo leído con éxito"}[0]}, nil
	} else {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Algo falló al leer el correo"}[0]}, nil
	}

}

func (s *Server) EliminarCorreosEntrada(ctx context.Context, in *pb.Correo) (*pb.Status, error) {
	correosLock.Lock()
	usersLock.Lock()
	//LIFO
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	defer reloadCorreoDBs()
	defer reloadUserDBs()

	target := int(*in.Identificador)
	r := -1
	for i, v := range usersMap[*in.Destinatario].BandejaEntradas {
		if v == target {
			r = i
		}
	}
	if r == -1 {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Algo falló al borrar el correo"}[0]}, nil
	}

	fmt.Println(usersMap[*in.Destinatario].BandejaEntradas)
	newEntrada := append(usersMap[*in.Destinatario].BandejaEntradas[:r], usersMap[*in.Destinatario].BandejaEntradas[r+1:]...)
	fmt.Println(newEntrada)

	tempUser := usersMap[*in.Destinatario]
	fmt.Println(tempUser)
	tempUser.BandejaEntradas = newEntrada
	usersMap[*in.Destinatario] = tempUser
	fmt.Println(newEntrada)
	fmt.Println(usersMap[*in.Destinatario])

	// delete(correosMap, int(*in.Identificador))

	return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Exito"}[0]}, nil

}
func (s *Server) EliminarCorreosSalida(ctx context.Context, in *pb.Correo) (*pb.Status, error) {
	correosLock.Lock()
	usersLock.Lock()
	//LIFO
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	// defer reloadCorreoDBs()
	defer reloadUserDBs()

	target := int(*in.Identificador)
	r := -1
	for i, v := range usersMap[*in.Emisor].BandejaSalidas {
		if v == target {
			r = i
		}
	}
	if r == -1 {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Algo falló al borrar el correo"}[0]}, nil
	}

	fmt.Println(usersMap[*in.Emisor].BandejaSalidas)
	newSalida := append(usersMap[*in.Emisor].BandejaSalidas[:r], usersMap[*in.Emisor].BandejaSalidas[r+1:]...)
	fmt.Println(newSalida)

	tempUser := usersMap[*in.Emisor]
	fmt.Println(tempUser)
	tempUser.BandejaSalidas = newSalida
	usersMap[*in.Emisor] = tempUser
	fmt.Println(newSalida)
	fmt.Println(usersMap[*in.Emisor])

	// delete(correosMap, int(*in.Identificador))

	return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Exito"}[0]}, nil
}

func reloadUserDBs() {

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

func reloadCorreoDBs() {

	jsonData, err := json.Marshal(correosMap)
	if err != nil {
		return
	}
	os.WriteFile(filenameCorreos, jsonData, 0644)

	data, err := os.ReadFile(filenameCorreos)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &correosMap); err != nil {
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

	data, err = os.ReadFile(filenameCorreos)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &correosMap); err != nil {
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
