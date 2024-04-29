from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Usuario(_message.Message):
    __slots__ = ("usuario", "contrasena")
    USUARIO_FIELD_NUMBER: _ClassVar[int]
    CONTRASENA_FIELD_NUMBER: _ClassVar[int]
    usuario: str
    contrasena: str
    def __init__(self, usuario: _Optional[str] = ..., contrasena: _Optional[str] = ...) -> None: ...

class Correo(_message.Message):
    __slots__ = ("identificador", "tema", "emisor", "destinatario", "contenido", "leido")
    IDENTIFICADOR_FIELD_NUMBER: _ClassVar[int]
    TEMA_FIELD_NUMBER: _ClassVar[int]
    EMISOR_FIELD_NUMBER: _ClassVar[int]
    DESTINATARIO_FIELD_NUMBER: _ClassVar[int]
    CONTENIDO_FIELD_NUMBER: _ClassVar[int]
    LEIDO_FIELD_NUMBER: _ClassVar[int]
    identificador: int
    tema: str
    emisor: str
    destinatario: str
    contenido: str
    leido: bool
    def __init__(self, identificador: _Optional[int] = ..., tema: _Optional[str] = ..., emisor: _Optional[str] = ..., destinatario: _Optional[str] = ..., contenido: _Optional[str] = ..., leido: bool = ...) -> None: ...

class Status(_message.Message):
    __slots__ = ("success", "mensaje")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    MENSAJE_FIELD_NUMBER: _ClassVar[int]
    success: bool
    mensaje: str
    def __init__(self, success: bool = ..., mensaje: _Optional[str] = ...) -> None: ...
