# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import turbomessage_pb2 as turbomessage__pb2


class TurboMessageStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.nuevoUsuario = channel.unary_unary(
                '/correos.TurboMessage/nuevoUsuario',
                request_serializer=turbomessage__pb2.Usuario.SerializeToString,
                response_deserializer=turbomessage__pb2.Status.FromString,
                )
        self.revisarUsuario = channel.unary_unary(
                '/correos.TurboMessage/revisarUsuario',
                request_serializer=turbomessage__pb2.Usuario.SerializeToString,
                response_deserializer=turbomessage__pb2.Status.FromString,
                )
        self.directorioUsuario = channel.unary_stream(
                '/correos.TurboMessage/directorioUsuario',
                request_serializer=turbomessage__pb2.Empty.SerializeToString,
                response_deserializer=turbomessage__pb2.Usuario.FromString,
                )
        self.enviarCorreo = channel.unary_unary(
                '/correos.TurboMessage/enviarCorreo',
                request_serializer=turbomessage__pb2.Correo.SerializeToString,
                response_deserializer=turbomessage__pb2.Status.FromString,
                )
        self.correosEntrada = channel.unary_stream(
                '/correos.TurboMessage/correosEntrada',
                request_serializer=turbomessage__pb2.Usuario.SerializeToString,
                response_deserializer=turbomessage__pb2.Correo.FromString,
                )
        self.correosSalida = channel.unary_stream(
                '/correos.TurboMessage/correosSalida',
                request_serializer=turbomessage__pb2.Usuario.SerializeToString,
                response_deserializer=turbomessage__pb2.Correo.FromString,
                )
        self.eliminarCorreosEntrada = channel.unary_unary(
                '/correos.TurboMessage/eliminarCorreosEntrada',
                request_serializer=turbomessage__pb2.Correo.SerializeToString,
                response_deserializer=turbomessage__pb2.Status.FromString,
                )
        self.eliminarCorreosSalida = channel.unary_unary(
                '/correos.TurboMessage/eliminarCorreosSalida',
                request_serializer=turbomessage__pb2.Correo.SerializeToString,
                response_deserializer=turbomessage__pb2.Status.FromString,
                )
        self.correoLeido = channel.unary_unary(
                '/correos.TurboMessage/correoLeido',
                request_serializer=turbomessage__pb2.Correo.SerializeToString,
                response_deserializer=turbomessage__pb2.Status.FromString,
                )


class TurboMessageServicer(object):
    """Missing associated documentation comment in .proto file."""

    def nuevoUsuario(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def revisarUsuario(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def directorioUsuario(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def enviarCorreo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def correosEntrada(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def correosSalida(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def eliminarCorreosEntrada(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def eliminarCorreosSalida(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def correoLeido(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TurboMessageServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'nuevoUsuario': grpc.unary_unary_rpc_method_handler(
                    servicer.nuevoUsuario,
                    request_deserializer=turbomessage__pb2.Usuario.FromString,
                    response_serializer=turbomessage__pb2.Status.SerializeToString,
            ),
            'revisarUsuario': grpc.unary_unary_rpc_method_handler(
                    servicer.revisarUsuario,
                    request_deserializer=turbomessage__pb2.Usuario.FromString,
                    response_serializer=turbomessage__pb2.Status.SerializeToString,
            ),
            'directorioUsuario': grpc.unary_stream_rpc_method_handler(
                    servicer.directorioUsuario,
                    request_deserializer=turbomessage__pb2.Empty.FromString,
                    response_serializer=turbomessage__pb2.Usuario.SerializeToString,
            ),
            'enviarCorreo': grpc.unary_unary_rpc_method_handler(
                    servicer.enviarCorreo,
                    request_deserializer=turbomessage__pb2.Correo.FromString,
                    response_serializer=turbomessage__pb2.Status.SerializeToString,
            ),
            'correosEntrada': grpc.unary_stream_rpc_method_handler(
                    servicer.correosEntrada,
                    request_deserializer=turbomessage__pb2.Usuario.FromString,
                    response_serializer=turbomessage__pb2.Correo.SerializeToString,
            ),
            'correosSalida': grpc.unary_stream_rpc_method_handler(
                    servicer.correosSalida,
                    request_deserializer=turbomessage__pb2.Usuario.FromString,
                    response_serializer=turbomessage__pb2.Correo.SerializeToString,
            ),
            'eliminarCorreosEntrada': grpc.unary_unary_rpc_method_handler(
                    servicer.eliminarCorreosEntrada,
                    request_deserializer=turbomessage__pb2.Correo.FromString,
                    response_serializer=turbomessage__pb2.Status.SerializeToString,
            ),
            'eliminarCorreosSalida': grpc.unary_unary_rpc_method_handler(
                    servicer.eliminarCorreosSalida,
                    request_deserializer=turbomessage__pb2.Correo.FromString,
                    response_serializer=turbomessage__pb2.Status.SerializeToString,
            ),
            'correoLeido': grpc.unary_unary_rpc_method_handler(
                    servicer.correoLeido,
                    request_deserializer=turbomessage__pb2.Correo.FromString,
                    response_serializer=turbomessage__pb2.Status.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'correos.TurboMessage', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TurboMessage(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def nuevoUsuario(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/correos.TurboMessage/nuevoUsuario',
            turbomessage__pb2.Usuario.SerializeToString,
            turbomessage__pb2.Status.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def revisarUsuario(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/correos.TurboMessage/revisarUsuario',
            turbomessage__pb2.Usuario.SerializeToString,
            turbomessage__pb2.Status.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def directorioUsuario(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/correos.TurboMessage/directorioUsuario',
            turbomessage__pb2.Empty.SerializeToString,
            turbomessage__pb2.Usuario.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def enviarCorreo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/correos.TurboMessage/enviarCorreo',
            turbomessage__pb2.Correo.SerializeToString,
            turbomessage__pb2.Status.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def correosEntrada(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/correos.TurboMessage/correosEntrada',
            turbomessage__pb2.Usuario.SerializeToString,
            turbomessage__pb2.Correo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def correosSalida(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/correos.TurboMessage/correosSalida',
            turbomessage__pb2.Usuario.SerializeToString,
            turbomessage__pb2.Correo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def eliminarCorreosEntrada(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/correos.TurboMessage/eliminarCorreosEntrada',
            turbomessage__pb2.Correo.SerializeToString,
            turbomessage__pb2.Status.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def eliminarCorreosSalida(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/correos.TurboMessage/eliminarCorreosSalida',
            turbomessage__pb2.Correo.SerializeToString,
            turbomessage__pb2.Status.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def correoLeido(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/correos.TurboMessage/correoLeido',
            turbomessage__pb2.Correo.SerializeToString,
            turbomessage__pb2.Status.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
