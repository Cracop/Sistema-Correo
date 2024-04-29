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
        self.sesion = {"usuario":"","id":0}
        self.puerto = puerto
        self.canal = grpc.insecure_channel("localhost:"+puerto)
        self.stub = turbomessage_pb2_grpc.TurboMessageStub(self.canal)

    def limpiar_pantalla(self):
        if os.name == 'nt':
            os.system('cls')
        else:
            os.system('clear')

    def menu_principal(self):
        self.limpiar_pantalla()
        print(("="*10) + "Bienvenido a 'TurboMessage'"+("="*10))
        print("1: Login")
        print("2: Registro")
        print("q: Salir")

        while True:
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
            self.sesion["id"]=int(respuesta.mensaje)
            self.sesion["usuario"]=usuario
            self.bandeja()
        else:
            print("ERROR: Ya existe un usuario con ese nombre")

    def login(self):
        usuario = input("\n #Usuario: ")
        passwd = input(" #Contraseña: ")
        respuesta = self.stub.revisarUsuario(turbomessage_pb2.Usuario(usuario=usuario, contrasena=passwd))
        if respuesta.success:
            print("Login completado con exito")
            self.sesion["id"]=int(respuesta.mensaje)
            self.sesion["usuario"]=usuario
            self.bandeja()
        else:
            print("ERROR: Credenciales Incorrectas")    

    def bandeja(self):
        print("Bienvenid@: "+self.sesion["usuario"]+str(self.sesion["id"]))
        



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