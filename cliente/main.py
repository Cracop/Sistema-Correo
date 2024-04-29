import grpc
import turbomessage_pb2
import turbomessage_pb2_grpc

def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = turbomessage_pb2_grpc.TurboMessageStub(channel)
        respuesta = stub.nuevoUsuario(turbomessage_pb2.Usuario(usuario="lala", contrasena="lala"))
        print(respuesta)


if __name__ == '__main__':
    run()