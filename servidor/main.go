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
	port            = flag.Int("port", 50051, "The server port") // Puerto del servidor
	filenameUsers   = "db_users.json" // Archivo de usuarios
	filenameCorreos = "db_correos.json" // Archivo de correos
	usersMap        = make(map[string]Usuario) // Mapa para almacenar usuarios
	correosMap      = make(map[int]Correo) // Mapa para almacenar correos
	usersLock       sync.Mutex // Mutex para sincronizar el acceso a usersMap
	correosLock     sync.Mutex // Mutex para sincronizar el acceso a correosMap
	numMax          = 5 // Número máximo de correos permitidos en bandeja
)

type Usuario struct {
	User            string `json:"usuario"`
	Passwd          string `json:"passwd"`
	BandejaEntradas []int  `json:"bandEntrada"`
	BandejaSalidas  []int  `json:"bandSalida"`
}

type Correo struct {
	Tema         string `json:"tema"`
	Destinatario string `json:"destinatario"`
	Emisor       string `json:"emisor"`
	Contenido    string `json:"contenido"`
	Leido        bool   `json:"leido"`
}

type Server struct {
	pb.TurboMessageServer
}

// NuevoUsuario agrega un nuevo usuario al sistema
func (s *Server) NuevoUsuario(ctx context.Context, in *pb.Usuario) (*pb.Status, error) {
	fmt.Print(in.Usuario)
	usersLock.Lock()
	defer usersLock.Unlock() // Desbloquea el mutex al final de la función
	defer reloadUserDBs() // Recarga la base de datos de usuarios

	if _, exists := usersMap[*in.Usuario]; !exists {
		usersMap[*in.Usuario] = Usuario{Passwd: *in.Contrasena}
		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"usuario creado con éxito"}[0]}, nil
	} else {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Usuario ya existe"}[0]}, nil
	}
}

// RevisarUsuario verifica si un usuario está registrado
func (s *Server) RevisarUsuario(ctx context.Context, in *pb.Usuario) (*pb.Status, error) {
	if _, exists := usersMap[*in.Usuario]; exists {
		return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Usuario existe"}[0]}, nil
	} else {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Usuario no existe"}[0]}, nil
	}
}

// DirectorioUsuario envía el directorio de usuarios a través de un stream
func (s *Server) DirectorioUsuario(em *pb.Empty, stream pb.TurboMessage_DirectorioUsuarioServer) error {
	for id := range usersMap {
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

// EnviarCorreo maneja el envío de correos entre usuarios
func (s *Server) EnviarCorreo(ctx context.Context, in *pb.Correo) (*pb.Status, error) {

	correosLock.Lock() // Bloquea el acceso concurrente a correosMap
	usersLock.Lock() // Bloquea el acceso concurrente a usersMap
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	defer reloadCorreoDBs()
	defer reloadUserDBs()

	// Verifica si el destinatario existe
	if _, exists := usersMap[*in.Destinatario]; !exists {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"No existe tal usuario"}[0]}, nil
	}

	// Verifica si la bandeja de salida del emisor está llena
	if len(usersMap[*in.Emisor].BandejaSalidas) >= numMax {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Bandeja de Salida llena"}[0]}, nil
	}

	// Verifica si la bandeja de entrada del destinatario está llena
	if len(usersMap[*in.Destinatario].BandejaEntradas) >= numMax {
		return &pb.Status{Success: &[]bool{false}[0], Mensaje: &[]string{"Bandeja de Entrada llena"}[0]}, nil
	}

	// Genera un ID único para el correo
	id := rand.Intn(101) + 1
	if _, exists := correosMap[id]; exists {
		for {
			id = rand.Intn(101) + 1
			if _, exists := correosMap[id]; !exists {
				// Si el ID no existe en el mapa, salir del bucle
				break
			}
		}
	}

	// Almacena el correo en correosMap
	correosMap[id] = Correo{
		Tema:         *in.Tema,
		Destinatario: *in.Destinatario,
		Emisor:       *in.Emisor,
		Contenido:    *in.Contenido,
		Leido:        *in.Leido,
	}

	// Actualiza la bandeja de salida del emisor
	usuario := usersMap[*in.Emisor]
	usuario.BandejaSalidas = append(usuario.BandejaSalidas, id)
	usersMap[*in.Emisor] = usuario

	// Actualiza la bandeja de entrada del destinatario
	usuario = usersMap[*in.Destinatario]
	usuario.BandejaEntradas = append(usuario.BandejaEntradas, id)
	usersMap[*in.Destinatario] = usuario

	return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Correo enviado con éxito"}[0]}, nil
}

// CorreosEntrada envía los correos en la bandeja de entrada del usuario a través de un stream
func (s *Server) CorreosEntrada(in *pb.Usuario, stream pb.TurboMessage_CorreosEntradaServer) error {
	correosLock.Lock() // Bloquea el acceso concurrente a correosMap
	usersLock.Lock() // Bloquea el acceso concurrente a usersMap
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	defer reloadCorreoDBs()
	defer reloadUserDBs()

	bandeja := usersMap[*in.Usuario].BandejaEntradas

	for _, id := range bandeja {
		correo := correosMap[id]
		ID := int32(id)

		tempCorreo := &pb.Correo{
			Identificador: &ID,
			Tema:          &correo.Tema,
			Destinatario:  &correo.Destinatario,
			Emisor:        &correo.Emisor,
			Contenido:     &correo.Contenido,
			Leido:         &correo.Leido,
		}

		if err := stream.Send(tempCorreo); err != nil {
			return err
		}
	}
	return nil
}

// CorreosSalida envía los correos en la bandeja de salida del usuario a través de un stream
func (s *Server) CorreosSalida(in *pb.Usuario, stream pb.TurboMessage_CorreosSalidaServer) error {
	correosLock.Lock() // Bloquea el acceso concurrente a correosMap
	usersLock.Lock() // Bloquea el acceso concurrente a usersMap
	defer correosLock.Unlock()
	defer usersLock.Unlock()
	defer reloadCorreoDBs()
	defer reloadUserDBs()

	bandeja := usersMap[*in.Usuario].BandejaSalidas

	for _, id := range bandeja {
		correo := correosMap[id]
		ID := int32(id)

		tempCorreo := &pb.Correo{
			Identificador: &ID,
			Tema:          &correo.Tema,
			Destinatario:  &correo.Destinatario,
			Emisor:        &correo.Emisor,
			Contenido:     &correo.Contenido,
			Leido:         &correo.Leido,
		}

		if err := stream.Send(tempCorreo); err != nil {
			return err
		}
	}
	return nil
}

// CorreoLeido marca un correo como leído
func (s *Server) CorreoLeido(ctx context.Context, in *pb.Correo) (*pb.Status, error) {
	correosLock.Lock() // Bloquea el acceso concurrente a correosMap
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

// EliminarCorreosEntrada elimina un correo de la bandeja de entrada
func (s *Server) EliminarCorreosEntrada(ctx context.Context, in *pb.Correo) (*pb.Status, error) {
	correosLock.Lock() // Bloquea el acceso concurrente a correosMap
	usersLock.Lock() // Bloquea el acceso concurrente a usersMap
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

	// Actualiza la bandeja de entrada del destinatario
	newEntrada := append(usersMap[*in.Destinatario].BandejaEntradas[:r], usersMap[*in.Destinatario].BandejaEntradas[r+1:]...)
	tempUser := usersMap[*in.Destinatario]
	tempUser.BandejaEntradas = newEntrada
	usersMap[*in.Destinatario] = tempUser

	return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Exito"}[0]}, nil
}

func (s *Server) EliminarCorreosSalida(ctx context.Context, in *pb.Correo) (*pb.Status, error) {
    correosLock.Lock() // Bloquea el acceso concurrente a correosMap
    usersLock.Lock() // Bloquea el acceso concurrente a usersMap
    defer correosLock.Unlock()
    defer usersLock.Unlock()
    defer reloadCorreoDBs()
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

    // Elimina el correo de la bandeja de salida del emisor
    fmt.Println(usersMap[*in.Emisor].BandejaSalidas)
    newSalida := append(usersMap[*in.Emisor].BandejaSalidas[:r], usersMap[*in.Emisor].BandejaSalidas[r+1:]...)
    fmt.Println(newSalida)

    tempUser := usersMap[*in.Emisor]
    fmt.Println(tempUser)
    tempUser.BandejaSalidas = newSalida
    usersMap[*in.Emisor] = tempUser
    fmt.Println(newSalida)
    fmt.Println(usersMap[*in.Emisor])

    // delete(correosMap, int(*in.Identificador)) // Comentar o descomentar según sea necesario

    return &pb.Status{Success: &[]bool{true}[0], Mensaje: &[]string{"Exito"}[0]}, nil
}

func reloadUserDBs() {
    // Recarga la base de datos de usuarios desde el archivo JSON
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
    // Recarga la base de datos de correos desde el archivo JSON
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
    // Inicializa los archivos de usuarios y correos si no existen
    _, err := os.Stat(filenameUsers)
    if os.IsNotExist(err) {
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
        file1, err := os.Create(filenameCorreos)
        if err != nil {
            fmt.Println("Error creating file:", err)
            return
        }
        defer file1.Close()
        fmt.Println("File created:", filenameCorreos)
    }

    // Carga los datos de los archivos JSON a los mapas correspondientes
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