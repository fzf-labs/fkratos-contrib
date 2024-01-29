// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: conf/v1/server.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 服务器
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"` // HTTP服务
	Grpc *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"` // gRPC服务
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_v1_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

// HTTP
type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network     string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`          // 网络
	Addr        string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`                // 服务监听地址
	Timeout     *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`          // 超时时间
	Middleware  *Middleware          `protobuf:"bytes,4,opt,name=middleware,proto3" json:"middleware,omitempty"`    // 中间件
	Cors        *Server_HTTP_CORS    `protobuf:"bytes,5,opt,name=cors,proto3" json:"cors,omitempty"`                // 跨域配置
	EnableCors  bool                 `protobuf:"varint,6,opt,name=enableCors,proto3" json:"enableCors,omitempty"`   // 启用跨域
	EnablePprof bool                 `protobuf:"varint,7,opt,name=enablePprof,proto3" json:"enablePprof,omitempty"` // 启用pprof
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_v1_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_HTTP) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

func (x *Server_HTTP) GetCors() *Server_HTTP_CORS {
	if x != nil {
		return x.Cors
	}
	return nil
}

func (x *Server_HTTP) GetEnableCors() bool {
	if x != nil {
		return x.EnableCors
	}
	return false
}

func (x *Server_HTTP) GetEnablePprof() bool {
	if x != nil {
		return x.EnablePprof
	}
	return false
}

// gPRC
type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network    string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`       // 网络
	Addr       string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`             // 服务监听地址
	Timeout    *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`       // 超时时间
	Middleware *Middleware          `protobuf:"bytes,4,opt,name=middleware,proto3" json:"middleware,omitempty"` // 中间件
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_v1_server_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_GRPC) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

type Server_HTTP_CORS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers []string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"` // 允许的请求头
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"` // 允许的请求方法
	Origins []string `protobuf:"bytes,3,rep,name=origins,proto3" json:"origins,omitempty"` // 允许的请求源
}

func (x *Server_HTTP_CORS) Reset() {
	*x = Server_HTTP_CORS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP_CORS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP_CORS) ProtoMessage() {}

func (x *Server_HTTP_CORS) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP_CORS.ProtoReflect.Descriptor instead.
func (*Server_HTTP_CORS) Descriptor() ([]byte, []int) {
	return file_conf_v1_server_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Server_HTTP_CORS) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Server_HTTP_CORS) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Server_HTTP_CORS) GetOrigins() []string {
	if x != nil {
		return x.Origins
	}
	return nil
}

var File_conf_v1_server_proto protoreflect.FileDescriptor

var file_conf_v1_server_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f,
	0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x04, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x25, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a,
	0xdf, 0x02, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x2a, 0x0a,
	0x04, 0x63, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x2e, 0x43,
	0x4f, 0x52, 0x53, 0x52, 0x04, 0x63, 0x6f, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x1a, 0x54, 0x0a, 0x04, 0x43,
	0x4f, 0x52, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x73, 0x1a, 0x9b, 0x01, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x30, 0x0a,
	0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x42,
	0x54, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x04,
	0x43, 0x6f, 0x6e, 0x66, 0xca, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0xe2, 0x02, 0x10, 0x43, 0x6f,
	0x6e, 0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x04, 0x43, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_server_proto_rawDescOnce sync.Once
	file_conf_v1_server_proto_rawDescData = file_conf_v1_server_proto_rawDesc
)

func file_conf_v1_server_proto_rawDescGZIP() []byte {
	file_conf_v1_server_proto_rawDescOnce.Do(func() {
		file_conf_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_server_proto_rawDescData)
	})
	return file_conf_v1_server_proto_rawDescData
}

var file_conf_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conf_v1_server_proto_goTypes = []interface{}{
	(*Server)(nil),              // 0: conf.Server
	(*Server_HTTP)(nil),         // 1: conf.Server.HTTP
	(*Server_GRPC)(nil),         // 2: conf.Server.GRPC
	(*Server_HTTP_CORS)(nil),    // 3: conf.Server.HTTP.CORS
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
	(*Middleware)(nil),          // 5: conf.Middleware
}
var file_conf_v1_server_proto_depIdxs = []int32{
	1, // 0: conf.Server.http:type_name -> conf.Server.HTTP
	2, // 1: conf.Server.grpc:type_name -> conf.Server.GRPC
	4, // 2: conf.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	5, // 3: conf.Server.HTTP.middleware:type_name -> conf.Middleware
	3, // 4: conf.Server.HTTP.cors:type_name -> conf.Server.HTTP.CORS
	4, // 5: conf.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	5, // 6: conf.Server.GRPC.middleware:type_name -> conf.Middleware
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_conf_v1_server_proto_init() }
func file_conf_v1_server_proto_init() {
	if File_conf_v1_server_proto != nil {
		return
	}
	file_conf_v1_middleware_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conf_v1_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_conf_v1_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_conf_v1_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
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
		file_conf_v1_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP_CORS); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_server_proto_goTypes,
		DependencyIndexes: file_conf_v1_server_proto_depIdxs,
		MessageInfos:      file_conf_v1_server_proto_msgTypes,
	}.Build()
	File_conf_v1_server_proto = out.File
	file_conf_v1_server_proto_rawDesc = nil
	file_conf_v1_server_proto_goTypes = nil
	file_conf_v1_server_proto_depIdxs = nil
}
