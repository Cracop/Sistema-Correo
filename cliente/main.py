import grpc
import sys
import os
sys.path.append('./grpc')
import turbomessage_pb2
import turbomessage_pb2_grpc

class Cliente:
    max_correos = 5
    def __init__(self, puerto):
        self.logeado = False
        self.sesion = {"usuario":""}
        self.puerto = puerto
        self.canal = grpc.insecure_channel("localhost:"+puerto)
        self.stub = turbomessage_pb2_grpc.TurboMessageStub(self.canal)
        self.bandejaEntrada = []
        self.bandejaSalida = []

    def limpiar_pantalla(self):
        if os.name == 'nt':
            os.system('cls')
        else:
            os.system('clear')

    def menu_principal(self):
        self.limpiar_pantalla()
        print(("="*10) + "Bienvenido a 'TurboMessage'"+("="*10))
        

        while True:
            print("1: Login")
            print("2: Registro")
            print("q: Salir")
            accion = input("#Ingresa que acción deseas realizar: ")

            if accion == "1":
                self.login()
            elif accion == "2":
                self.registrarse()
            elif accion == "q":
                exit()
            else:
                print("No hay tal acción")

            
    def registrarse(self):
        self.limpiar_pantalla()
        correcta = False
        print("Gracias por elegir TurboMessage, por favor completa la siguiente información")
        usuario = input("\n #Usuario: ")
        while  not correcta:
            # usuario = input("\n Usuario: ")
            passwd1 = input(" #Contraseña: ")
            passwd2 = input(" #Confirma tu contraseña ")

            if passwd1 == passwd2:
                correcta = True
            else:
                print("#Las contraseñas no coinciden, por favor intentalo de nuevo")

        respuesta = self.stub.nuevoUsuario(turbomessage_pb2.Usuario(usuario=usuario, contrasena=passwd1))
        if respuesta.success:
            print("Registro completado con exito")
            print("Procediendo a logearte directo")
            self.sesion["usuario"]=usuario
            self.bandeja()
        else:
            print("ERROR: Ya existe un usuario con ese nombre")

    def login(self):
        self.limpiar_pantalla()
        usuario = input("\n #Usuario: ")
        passwd = input(" #Contraseña: ")
        respuesta = self.stub.revisarUsuario(turbomessage_pb2.Usuario(usuario=usuario, contrasena=passwd))
        if respuesta.success:
            print("Login completado con exito")
            self.sesion["usuario"]=usuario
            self.bandeja()
        else:
            print("ERROR: Credenciales Incorrectas")    

    def bandeja(self):
        self.limpiar_pantalla()
        print("Bienvenid@: "+self.sesion["usuario"])
        while True:
            print("1: Ver bandeja de entrada")
            print("2: Ver bandeja de salida")
            print("3: Escribir correo")
            print("4: Ver todos los usuarios")
            print("q: Regresar")
            accion = input("#Ingresa que acción deseas realizar: ")

            if accion == "1":
                self.verBandejaEntrada()
            elif accion == "2":
                self.verBandejaSalida()
            elif accion == "3":
                self.escribirCorreo()
            elif accion == "4":
                self.verDirectorio()
            elif accion == "q":
                return
            else:
                print("No hay tal acción")
        

    def verBandejaEntrada(self):
        self.limpiar_pantalla()  

        while True:
            self.bandejaEntrada.clear()
            print("BANDEJA DE ENTRADA DE "+self.sesion["usuario"])
            
            correos = self.stub.correosEntrada(turbomessage_pb2.Usuario(usuario=self.sesion["usuario"], contrasena=""))
            for correo in correos:
                self.bandejaEntrada.append(correo)
                print(("-"*30))
                print("ID:"+str(len(self.bandejaEntrada))+"| Asunto: "+ correo.tema +"| De:" + correo.emisor+ " | Leído: "+str(correo.leido))
                print(("-"*30)) 

            action = input("Escribe el ID del correo que quieres leer o 'q' si quieres regresar: ")
            if action == "q":
                return
            
            try:
                ac = int(action)
                # print(ac)
                # print(type(ac))
                # print(len(self.bandejaEntrada))
                ac -=1
                if ac < len(self.bandejaEntrada):
                    self.verCorreo(ac, True) 
                else:
                    print("No hay tal correo")

            except:
                print("No hay tal accion")     

    def verBandejaSalida(self):
        self.limpiar_pantalla()

        while True:

            self.bandejaSalida.clear()
            print("BANDEJA DE SALIDA DE "+self.sesion["usuario"])
            correos = self.stub.correosSalida(turbomessage_pb2.Usuario(usuario=self.sesion["usuario"], contrasena=""))
            for correo in correos:
                self.bandejaSalida.append(correo)
                print(("-"*30))
                print("ID:"+str(len(self.bandejaSalida))+"| Asunto: "+ correo.tema +"| A: " + correo.destinatario )
            print(("-"*30))

            action = input("Escribe el ID del correo que quieres leer o 'q' si quieres regresar: ")
            if action == "q":
                return
            
            try:
                ac = int(action)
                # print(ac)
                # print(type(ac))
                # print(len(self.bandejaSalida))
                ac -=1
                if ac < len(self.bandejaSalida):
                    self.verCorreo(ac, False) 
                else:
                    print("No hay tal correo")

            except:
                    print("No hay tal accion") 
        

    def escribirCorreo(self):
        self.limpiar_pantalla()
        print("Eres: "+self.sesion["usuario"])
        self.verDirectorio()
        print(("="*10) + "Escritura de Correo'"+("="*10))
        dest = input("Destinatario: ")
        tema = input("Asunto: ")
        contenido = input("Contenido: \n")
        respuesta = self.stub.enviarCorreo(turbomessage_pb2.Correo(identificador=0, 
                                                                   tema=tema,
                                                                   emisor=self.sesion["usuario"],
                                                                   destinatario=dest,
                                                                   contenido=contenido,
                                                                   leido=False))
        
        print(respuesta.mensaje)


    def verDirectorio(self):
        print(("="*10) + "Directorio de Usuarios'"+("="*10))
        for elemento in self.stub.directorioUsuario(turbomessage_pb2.Empty()):
            print("Nombre: "+str(elemento.usuario))
        print(("="*12) + "Fin del Directorio'"+("="*12))
        
    def verCorreo(self,correoID_local, entrada):
        # print(correoID_local)
        if entrada:
            id = self.bandejaEntrada[correoID_local].identificador
            self.stub.correoLeido(turbomessage_pb2.Correo(identificador=id, 
                                                                   tema="",
                                                                   emisor="",
                                                                   destinatario="",
                                                                   contenido="",
                                                                   leido=True))
            # print(self.bandejaEntrada[correoID_local])
        self.imprimeCorreo(correoID_local,entrada)
        # else:
        #     print(self.bandejaSalida[correoID_local])


    def imprimeCorreo(self,id, entrada):
        if entrada:
            correo = self.bandejaEntrada[id]
        else:
            correo = self.bandejaSalida[id]
        print(("="*30))
        print("ASUNTO: "+correo.tema)
        print("De: "+correo.emisor)
        print("A: "+correo.destinatario)
        print(("-"*30))
        print("A: "+correo.contenido)
        print(("="*30))

        while True:
            ac = input("Borrar? (Y/N)")
            if ac == "Y":
                
                if entrada:
                    self.eliminarCorreoEntrada(id)
                else:
                    self.eliminarCorreoSalida(id)

                return
            elif ac=="N":
                return
            else:
                print("No hay tal acción")


    def eliminarCorreoEntrada(self, correoID):
        id = self.bandejaEntrada[correoID].identificador
        dest = self.bandejaEntrada[correoID].destinatario
        resp = self.stub.eliminarCorreosEntrada(turbomessage_pb2.Correo(identificador=id, 
                                                                   tema="",
                                                                   emisor="",
                                                                   destinatario=dest,
                                                                   contenido="",
                                                                   leido=True))
        print(resp.mensaje)

    def eliminarCorreoSalida(self, correoID):
        id = self.bandejaSalida[correoID].identificador
        emi = self.bandejaSalida[correoID].emisor
        resp = self.stub.eliminarCorreosSalida(turbomessage_pb2.Correo(identificador=id, 
                                                                   tema="",
                                                                   emisor=emi,
                                                                   destinatario="",
                                                                   contenido="",
                                                                   leido=True))
        print(resp.mensaje)

    

# def run():
#     with grpc.insecure_channel("localhost:50051") as channel:
#         stub = turbomessage_pb2_grpc.TurboMessageStub(channel)
#         respuesta = stub.nuevoUsuario(turbomessage_pb2.Usuario(usuario="lala", contrasena="lala"))
#         print(respuesta)


if __name__ == '__main__':
    # puerto = input("¿A que puerto desea conectarse?\n Puerto: ")
    puerto = "50051"
    cliente = Cliente(puerto)
    cliente.menu_principal()