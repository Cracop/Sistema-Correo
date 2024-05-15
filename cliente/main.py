import grpc
import sys
import os

# Añade el directorio './grpc' al PATH del sistema para importar módulos de ese directorio.
sys.path.append('./grpc')

# Importa los módulos generados por Protocol Buffers.
import turbomessage_pb2
import turbomessage_pb2_grpc

class Cliente:
    max_correos = 5  # Define el número máximo de correos permitidos en bandejas de entrada y salida.

    def __init__(self, puerto):
        # Inicializa la clase Cliente con la configuración de gRPC y las variables necesarias.
        self.logeado = False  # Estado de inicio de sesión del usuario.
        self.sesion = {"usuario":""}  # Información de la sesión del usuario.
        self.puerto = puerto  # Puerto al que se conectará el cliente.
        self.canal = grpc.insecure_channel("localhost:" + puerto)  # Canal de comunicación gRPC.
        self.stub = turbomessage_pb2_grpc.TurboMessageStub(self.canal)  # Stub para llamar a los métodos del servidor.
        self.bandejaEntrada = []  # Lista para almacenar correos en la bandeja de entrada.
        self.bandejaSalida = []  # Lista para almacenar correos en la bandeja de salida.

    def limpiar_pantalla(self):
        # Limpia la pantalla de la consola para una mejor presentación.
        if os.name == 'nt':
            os.system('cls')  # Comando para limpiar la pantalla en Windows.
        else:
            os.system('clear')  # Comando para limpiar la pantalla en otros sistemas operativos.

    def menu_principal(self):
        # Muestra el menú principal donde el usuario puede elegir entre iniciar sesión, registrarse o salir.
        self.limpiar_pantalla()
        print(("="*10) + "Bienvenido a 'TurboMessage'" + ("="*10))

        while True:
            print("1: Login")
            print("2: Registro")
            print("q: Salir")
            accion = input("#Ingresa que acción deseas realizar: ")

            if accion == "1":
                self.login()  # Llama al método para iniciar sesión.
            elif accion == "2":
                self.registrarse()  # Llama al método para registrarse.
            elif accion == "q":
                exit()  # Sale del programa.
            else:
                print("No hay tal acción")  # Mensaje de error si la acción ingresada no es válida.


            
    def registrarse(self):
        # Permite al usuario registrarse en el sistema.
        self.limpiar_pantalla()
        correcta = False
        print("Gracias por elegir TurboMessage, por favor completa la siguiente información")
        usuario = input("\n #Usuario: ")

        while not correcta:
            # Solicita y confirma la contraseña del usuario.
            passwd1 = input(" #Contraseña: ")
            passwd2 = input(" #Confirma tu contraseña ")

            if passwd1 == passwd2:
                correcta = True  # Las contraseñas coinciden.
            else:
                print("#Las contraseñas no coinciden, por favor intentalo de nuevo")

        # Envía la solicitud de registro al servidor.
        respuesta = self.stub.nuevoUsuario(turbomessage_pb2.Usuario(usuario=usuario, contrasena=passwd1))
        if respuesta.success:
            print("Registro completado con exito")
            print("Procediendo a logearte directo")
            self.sesion["usuario"] = usuario
            self.bandeja()  # Redirige al menú de la bandeja.
        else:
            print("ERROR: Ya existe un usuario con ese nombre")

    def login(self):
        # Permite al usuario iniciar sesión en el sistema.
        self.limpiar_pantalla()
        usuario = input("\n #Usuario: ")
        passwd = input(" #Contraseña: ")

        # Envía la solicitud de autenticación al servidor.
        respuesta = self.stub.revisarUsuario(turbomessage_pb2.Usuario(usuario=usuario, contrasena=passwd))
        if respuesta.success:
            print("Login completado con exito")
            self.sesion["usuario"] = usuario
            self.bandeja()  # Redirige al menú de la bandeja.
        else:
            print("ERROR: Credenciales Incorrectas")

    def bandeja(self):
        # Muestra el menú de la bandeja de usuario autenticado.
        self.limpiar_pantalla()
        print("Bienvenid@: " + self.sesion["usuario"])

        while True:
            print("1: Ver bandeja de entrada")
            print("2: Ver bandeja de salida")
            print("3: Escribir correo")
            print("4: Ver todos los usuarios")
            print("q: Regresar")
            accion = input("#Ingresa que acción deseas realizar: ")

            if accion == "1":
                self.verBandejaEntrada()  # Ver la bandeja de entrada.
            elif accion == "2":
                self.verBandejaSalida()  # Ver la bandeja de salida.
            elif accion == "3":
                self.escribirCorreo()  # Escribir un nuevo correo.
            elif accion == "4":
                self.verDirectorio()  # Ver el directorio de usuarios.
            elif accion == "q":
                return  # Regresar al menú principal.
            else:
                print("No hay tal acción")

    def verBandejaEntrada(self):
        # Muestra los correos en la bandeja de entrada del usuario.
        self.limpiar_pantalla()

        while True:
            self.bandejaEntrada.clear()  # Limpia la lista de correos.
            print("BANDEJA DE ENTRADA DE " + self.sesion["usuario"])

            # Solicita los correos de la bandeja de entrada del usuario al servidor.
            correos = self.stub.correosEntrada(turbomessage_pb2.Usuario(usuario=self.sesion["usuario"], contrasena=""))
            for correo in correos:
                self.bandejaEntrada.append(correo)
                print(("-"*30))
                print("ID:" + str(len(self.bandejaEntrada)) + "| Asunto: " + correo.tema + "| De:" + correo.emisor + " | Leído: " + str(correo.leido))
                print(("-"*30))

            action = input("Escribe el ID del correo que quieres leer o 'q' si quieres regresar: ")
            if action == "q":
                return  # Regresa al menú de la bandeja.

            try:
                ac = int(action)
                ac -= 1
                if ac < len(self.bandejaEntrada):
                    self.verCorreo(ac, True)  # Ver el correo seleccionado.
                else:
                    print("No hay tal correo")
            except:
                print("No hay tal accion")  # Maneja la excepción si el ID ingresado no es válido.
     

    def verBandejaSalida(self):
        # Muestra los correos en la bandeja de salida del usuario.
        self.limpiar_pantalla()

        while True:
            self.bandejaSalida.clear()  # Limpia la lista de correos.
            print("BANDEJA DE SALIDA DE " + self.sesion["usuario"])

            # Solicita los correos de la bandeja de salida del usuario al servidor.
            correos = self.stub.correosSalida(turbomessage_pb2.Usuario(usuario=self.sesion["usuario"], contrasena=""))
            for correo in correos:
                self.bandejaSalida.append(correo)
                print(("-"*30))
                print("ID:" + str(len(self.bandejaSalida)) + "| Asunto: " + correo.tema + "| A: " + correo.destinatario)
            print(("-"*30))

            action = input("Escribe el ID del correo que quieres leer o 'q' si quieres regresar: ")
            if action == "q":
                return  # Regresa al menú de la bandeja.

            try:
                ac = int(action)
                ac -= 1
                if ac < len(self.bandejaSalida):
                    self.verCorreo(ac, False)  # Ver el correo seleccionado.
                else:
                    print("No hay tal correo")
            except:
                print("No hay tal accion")  # Maneja la excepción si el ID ingresado no es válido.

    def escribirCorreo(self):
        # Permite al usuario escribir y enviar un correo.
        self.limpiar_pantalla()
        print("Eres: " + self.sesion["usuario"])
        self.verDirectorio()
        print(("="*10) + "Escritura de Correo'" + ("="*10))
        dest = input("Destinatario: ")
        tema = input("Asunto: ")
        contenido = input("Contenido: \n")
        
        # Envía el correo al servidor.
        respuesta = self.stub.enviarCorreo(turbomessage_pb2.Correo(
            identificador=0,
            tema=tema,
            emisor=self.sesion["usuario"],
            destinatario=dest,
            contenido=contenido,
            leido=False
        ))
        
        print(respuesta.mensaje)

    def verDirectorio(self):
        # Muestra el directorio de usuarios.
        print(("="*10) + "Directorio de Usuarios'" + ("="*10))
        for elemento in self.stub.directorioUsuario(turbomessage_pb2.Empty()):
            print("Nombre: " + str(elemento.usuario))
        print(("="*12) + "Fin del Directorio'" + ("="*12))

    def verCorreo(self, correoID_local, entrada):
        # Muestra el contenido de un correo específico.
        if entrada:
            id = self.bandejaEntrada[correoID_local].identificador
            # Marca el correo como leído en el servidor.
            self.stub.correoLeido(turbomessage_pb2.Correo(
                identificador=id,
                tema="",
                emisor="",
                destinatario="",
                contenido="",
                leido=True
            ))
        self.imprimeCorreo(correoID_local, entrada)

    def imprimeCorreo(self, id, entrada):
        # Imprime el contenido de un correo en la consola.
        if entrada:
            correo = self.bandejaEntrada[id]
        else:
            correo = self.bandejaSalida[id]
        
        print(("="*30))
        print("ASUNTO: " + correo.tema)
        print("De: " + correo.emisor)
        print("A: " + correo.destinatario)
        print(("-"*30))
        print("Contenido: " + correo.contenido)
        print(("="*30))

        while True:
            ac = input("Borrar? (Y/N): ")
            if ac == "Y":
                if entrada:
                    self.eliminarCorreoEntrada(id)
                else:
                    self.eliminarCorreoSalida(id)
                return
            elif ac == "N":
                return
            else:
                print("No hay tal acción")

    def eliminarCorreoEntrada(self, correoID):
        # Elimina un correo de la bandeja de entrada.
        id = self.bandejaEntrada[correoID].identificador
        dest = self.bandejaEntrada[correoID].destinatario
        resp = self.stub.eliminarCorreosEntrada(turbomessage_pb2.Correo(
            identificador=id,
            tema="",
            emisor="",
            destinatario=dest,
            contenido="",
            leido=True
        ))
        print(resp.mensaje)

    def eliminarCorreoSalida(self, correoID):
        # Elimina un correo de la bandeja de salida.
        id = self.bandejaSalida[correoID].identificador
        emi = self.bandejaSalida[correoID].emisor
        resp = self.stub.eliminarCorreosSalida(turbomessage_pb2.Correo(
            identificador=id,
            tema="",
            emisor=emi,
            destinatario="",
            contenido="",
            leido=True
        ))
        print(resp.mensaje)

# def run():
#     with grpc.insecure_channel("localhost:50051") as channel:
#         stub = turbomessage_pb2_grpc.TurboMessageStub(channel)
#         respuesta = stub.nuevoUsuario(turbomessage_pb2.Usuario(usuario="lala", contrasena="lala"))
#         print(respuesta)

if __name__ == '__main__':
    # Punto de entrada del programa. Inicia el cliente y muestra el menú principal.
    puerto = "50051"
    cliente = Cliente(puerto)
    cliente.menu_principal()
