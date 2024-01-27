// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: conf/v1/business.proto

package v1

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

// 业务
type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt *Business_Jwt `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_business_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_business_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_conf_v1_business_proto_rawDescGZIP(), []int{0}
}

func (x *Business) GetJwt() *Business_Jwt {
	if x != nil {
		return x.Jwt
	}
	return nil
}

// jwt
type Business_Jwt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessSecret string `protobuf:"bytes,1,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"`  //秘钥
	AccessExpire int64  `protobuf:"varint,2,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"` //过期时间
	RefreshAfter int64  `protobuf:"varint,3,opt,name=refreshAfter,proto3" json:"refreshAfter,omitempty"` //刷新时间 (大于刷新时间,而小于过期时间,则刷新Token)
	Issuer       string `protobuf:"bytes,4,opt,name=issuer,proto3" json:"issuer,omitempty"`              //签发人
}

func (x *Business_Jwt) Reset() {
	*x = Business_Jwt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_business_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business_Jwt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business_Jwt) ProtoMessage() {}

func (x *Business_Jwt) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_business_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business_Jwt.ProtoReflect.Descriptor instead.
func (*Business_Jwt) Descriptor() ([]byte, []int) {
	return file_conf_v1_business_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Business_Jwt) GetAccessSecret() string {
	if x != nil {
		return x.AccessSecret
	}
	return ""
}

func (x *Business_Jwt) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *Business_Jwt) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

func (x *Business_Jwt) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

var File_conf_v1_business_proto protoreflect.FileDescriptor

var file_conf_v1_business_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x22, 0xbc,
	0x01, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x03, 0x6a,
	0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4a, 0x77, 0x74, 0x52, 0x03, 0x6a, 0x77,
	0x74, 0x1a, 0x89, 0x01, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x42, 0x56, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x04,
	0x43, 0x6f, 0x6e, 0x66, 0xca, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0xe2, 0x02, 0x10, 0x43, 0x6f,
	0x6e, 0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x04, 0x43, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_business_proto_rawDescOnce sync.Once
	file_conf_v1_business_proto_rawDescData = file_conf_v1_business_proto_rawDesc
)

func file_conf_v1_business_proto_rawDescGZIP() []byte {
	file_conf_v1_business_proto_rawDescOnce.Do(func() {
		file_conf_v1_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_business_proto_rawDescData)
	})
	return file_conf_v1_business_proto_rawDescData
}

var file_conf_v1_business_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_conf_v1_business_proto_goTypes = []interface{}{
	(*Business)(nil),     // 0: conf.Business
	(*Business_Jwt)(nil), // 1: conf.Business.Jwt
}
var file_conf_v1_business_proto_depIdxs = []int32{
	1, // 0: conf.Business.jwt:type_name -> conf.Business.Jwt
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_conf_v1_business_proto_init() }
func file_conf_v1_business_proto_init() {
	if File_conf_v1_business_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_v1_business_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business); i {
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
		file_conf_v1_business_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business_Jwt); i {
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
			RawDescriptor: file_conf_v1_business_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_business_proto_goTypes,
		DependencyIndexes: file_conf_v1_business_proto_depIdxs,
		MessageInfos:      file_conf_v1_business_proto_msgTypes,
	}.Build()
	File_conf_v1_business_proto = out.File
	file_conf_v1_business_proto_rawDesc = nil
	file_conf_v1_business_proto_goTypes = nil
	file_conf_v1_business_proto_depIdxs = nil
}