// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: turbomessage.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Usuario struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usuario    *string `protobuf:"bytes,1,opt,name=usuario,proto3,oneof" json:"usuario,omitempty"`
	Contrasena *string `protobuf:"bytes,2,opt,name=contrasena,proto3,oneof" json:"contrasena,omitempty"`
}

func (x *Usuario) Reset() {
	*x = Usuario{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turbomessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usuario) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usuario) ProtoMessage() {}

func (x *Usuario) ProtoReflect() protoreflect.Message {
	mi := &file_turbomessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usuario.ProtoReflect.Descriptor instead.
func (*Usuario) Descriptor() ([]byte, []int) {
	return file_turbomessage_proto_rawDescGZIP(), []int{0}
}

func (x *Usuario) GetUsuario() string {
	if x != nil && x.Usuario != nil {
		return *x.Usuario
	}
	return ""
}

func (x *Usuario) GetContrasena() string {
	if x != nil && x.Contrasena != nil {
		return *x.Contrasena
	}
	return ""
}

type Correo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identificador *int32  `protobuf:"varint,1,opt,name=identificador,proto3,oneof" json:"identificador,omitempty"`
	Tema          *string `protobuf:"bytes,2,opt,name=tema,proto3,oneof" json:"tema,omitempty"`
	Emisor        *string `protobuf:"bytes,3,opt,name=emisor,proto3,oneof" json:"emisor,omitempty"`
	Destinatario  *string `protobuf:"bytes,4,opt,name=destinatario,proto3,oneof" json:"destinatario,omitempty"`
	Contenido     *string `protobuf:"bytes,5,opt,name=contenido,proto3,oneof" json:"contenido,omitempty"`
	Leido         *bool   `protobuf:"varint,6,opt,name=leido,proto3,oneof" json:"leido,omitempty"`
}

func (x *Correo) Reset() {
	*x = Correo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turbomessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Correo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Correo) ProtoMessage() {}

func (x *Correo) ProtoReflect() protoreflect.Message {
	mi := &file_turbomessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Correo.ProtoReflect.Descriptor instead.
func (*Correo) Descriptor() ([]byte, []int) {
	return file_turbomessage_proto_rawDescGZIP(), []int{1}
}

func (x *Correo) GetIdentificador() int32 {
	if x != nil && x.Identificador != nil {
		return *x.Identificador
	}
	return 0
}

func (x *Correo) GetTema() string {
	if x != nil && x.Tema != nil {
		return *x.Tema
	}
	return ""
}

func (x *Correo) GetEmisor() string {
	if x != nil && x.Emisor != nil {
		return *x.Emisor
	}
	return ""
}

func (x *Correo) GetDestinatario() string {
	if x != nil && x.Destinatario != nil {
		return *x.Destinatario
	}
	return ""
}

func (x *Correo) GetContenido() string {
	if x != nil && x.Contenido != nil {
		return *x.Contenido
	}
	return ""
}

func (x *Correo) GetLeido() bool {
	if x != nil && x.Leido != nil {
		return *x.Leido
	}
	return false
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success *bool   `protobuf:"varint,1,opt,name=success,proto3,oneof" json:"success,omitempty"`
	Mensaje *string `protobuf:"bytes,2,opt,name=mensaje,proto3,oneof" json:"mensaje,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turbomessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_turbomessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_turbomessage_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *Status) GetMensaje() string {
	if x != nil && x.Mensaje != nil {
		return *x.Mensaje
	}
	return ""
}

type ListadoUsuarios struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objetos []*Usuario `protobuf:"bytes,1,rep,name=objetos,proto3" json:"objetos,omitempty"`
}

func (x *ListadoUsuarios) Reset() {
	*x = ListadoUsuarios{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turbomessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListadoUsuarios) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListadoUsuarios) ProtoMessage() {}

func (x *ListadoUsuarios) ProtoReflect() protoreflect.Message {
	mi := &file_turbomessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListadoUsuarios.ProtoReflect.Descriptor instead.
func (*ListadoUsuarios) Descriptor() ([]byte, []int) {
	return file_turbomessage_proto_rawDescGZIP(), []int{3}
}

func (x *ListadoUsuarios) GetObjetos() []*Usuario {
	if x != nil {
		return x.Objetos
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turbomessage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_turbomessage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_turbomessage_proto_rawDescGZIP(), []int{4}
}

var File_turbomessage_proto protoreflect.FileDescriptor

var file_turbomessage_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x22, 0x68, 0x0a,
	0x07, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x73, 0x75, 0x61,
	0x72, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x75, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x73, 0x65, 0x6e, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x73, 0x65, 0x6e, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x73, 0x65, 0x6e, 0x61, 0x22, 0x9f, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x72, 0x72,
	0x65, 0x6f, 0x12, 0x29, 0x0a, 0x0d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x74, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x74,
	0x65, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x65, 0x6d, 0x69, 0x73, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x65, 0x6d, 0x69, 0x73, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x61,
	0x72, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x69, 0x64, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x69, 0x64, 0x6f, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x6c, 0x65, 0x69, 0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05,
	0x52, 0x05, 0x6c, 0x65, 0x69, 0x64, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x74, 0x65, 0x6d, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x6d, 0x69, 0x73, 0x6f, 0x72,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x61, 0x72, 0x69,
	0x6f, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x69, 0x64, 0x6f, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x69, 0x64, 0x6f, 0x22, 0x5e, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x12, 0x2a, 0x0a, 0x07,
	0x6f, 0x62, 0x6a, 0x65, 0x74, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x52,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x74, 0x6f, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0x88, 0x04, 0x0a, 0x0c, 0x54, 0x75, 0x72, 0x62, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x55, 0x73, 0x75, 0x61, 0x72,
	0x69, 0x6f, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x61, 0x72, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x11, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x55, 0x73, 0x75, 0x61,
	0x72, 0x69, 0x6f, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x55, 0x73,
	0x75, 0x61, 0x72, 0x69, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x0c, 0x65, 0x6e, 0x76,
	0x69, 0x61, 0x72, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x64, 0x61, 0x12,
	0x10, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69,
	0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x72,
	0x65, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f,
	0x73, 0x53, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f,
	0x73, 0x2e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c,
	0x0a, 0x16, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x64, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x15,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x53,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73, 0x2e,
	0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6f, 0x4c, 0x65, 0x69, 0x64, 0x6f, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6f, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_turbomessage_proto_rawDescOnce sync.Once
	file_turbomessage_proto_rawDescData = file_turbomessage_proto_rawDesc
)

func file_turbomessage_proto_rawDescGZIP() []byte {
	file_turbomessage_proto_rawDescOnce.Do(func() {
		file_turbomessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_turbomessage_proto_rawDescData)
	})
	return file_turbomessage_proto_rawDescData
}

var file_turbomessage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_turbomessage_proto_goTypes = []interface{}{
	(*Usuario)(nil),         // 0: correos.Usuario
	(*Correo)(nil),          // 1: correos.Correo
	(*Status)(nil),          // 2: correos.Status
	(*ListadoUsuarios)(nil), // 3: correos.ListadoUsuarios
	(*Empty)(nil),           // 4: correos.Empty
}
var file_turbomessage_proto_depIdxs = []int32{
	0,  // 0: correos.ListadoUsuarios.objetos:type_name -> correos.Usuario
	0,  // 1: correos.TurboMessage.nuevoUsuario:input_type -> correos.Usuario
	0,  // 2: correos.TurboMessage.revisarUsuario:input_type -> correos.Usuario
	4,  // 3: correos.TurboMessage.directorioUsuario:input_type -> correos.Empty
	1,  // 4: correos.TurboMessage.enviarCorreo:input_type -> correos.Correo
	0,  // 5: correos.TurboMessage.correosEntrada:input_type -> correos.Usuario
	0,  // 6: correos.TurboMessage.correosSalida:input_type -> correos.Usuario
	1,  // 7: correos.TurboMessage.eliminarCorreosEntrada:input_type -> correos.Correo
	1,  // 8: correos.TurboMessage.eliminarCorreosSalida:input_type -> correos.Correo
	1,  // 9: correos.TurboMessage.correoLeido:input_type -> correos.Correo
	2,  // 10: correos.TurboMessage.nuevoUsuario:output_type -> correos.Status
	2,  // 11: correos.TurboMessage.revisarUsuario:output_type -> correos.Status
	0,  // 12: correos.TurboMessage.directorioUsuario:output_type -> correos.Usuario
	2,  // 13: correos.TurboMessage.enviarCorreo:output_type -> correos.Status
	1,  // 14: correos.TurboMessage.correosEntrada:output_type -> correos.Correo
	1,  // 15: correos.TurboMessage.correosSalida:output_type -> correos.Correo
	2,  // 16: correos.TurboMessage.eliminarCorreosEntrada:output_type -> correos.Status
	2,  // 17: correos.TurboMessage.eliminarCorreosSalida:output_type -> correos.Status
	2,  // 18: correos.TurboMessage.correoLeido:output_type -> correos.Status
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_turbomessage_proto_init() }
func file_turbomessage_proto_init() {
	if File_turbomessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_turbomessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usuario); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_turbomessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Correo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_turbomessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_turbomessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListadoUsuarios); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_turbomessage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_turbomessage_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_turbomessage_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_turbomessage_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_turbomessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_turbomessage_proto_goTypes,
		DependencyIndexes: file_turbomessage_proto_depIdxs,
		MessageInfos:      file_turbomessage_proto_msgTypes,
	}.Build()
	File_turbomessage_proto = out.File
	file_turbomessage_proto_rawDesc = nil
	file_turbomessage_proto_goTypes = nil
	file_turbomessage_proto_depIdxs = nil
}
