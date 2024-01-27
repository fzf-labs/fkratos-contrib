// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: conf/v1/data.proto

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

// 数据
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gorm  *Data_Gorm  `protobuf:"bytes,1,opt,name=gorm,proto3" json:"gorm,omitempty"`   // 数据库 gorm
	Redis *Data_Redis `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"` // Redis
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetGorm() *Data_Gorm {
	if x != nil {
		return x.Gorm
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

// 数据库 gorm
type Data_Gorm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSourceName  string               `protobuf:"bytes,1,opt,name=dataSourceName,proto3" json:"dataSourceName,omitempty"`   // DSN
	MaxIdleConn     int32                `protobuf:"varint,2,opt,name=maxIdleConn,proto3" json:"maxIdleConn,omitempty"`        // 闲置连接数
	MaxOpenConn     int32                `protobuf:"varint,3,opt,name=maxOpenConn,proto3" json:"maxOpenConn,omitempty"`        // 最大打开的连接数
	ConnMaxLifeTime *durationpb.Duration `protobuf:"bytes,4,opt,name=connMaxLifeTime,proto3" json:"connMaxLifeTime,omitempty"` // 连接可以重复使用的最长时间
	ShowLog         bool                 `protobuf:"varint,5,opt,name=showLog,proto3" json:"showLog,omitempty"`                // 慢日志开关
	Tracing         bool                 `protobuf:"varint,6,opt,name=tracing,proto3" json:"tracing,omitempty"`                // 链路追踪开关
}

func (x *Data_Gorm) Reset() {
	*x = Data_Gorm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Gorm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Gorm) ProtoMessage() {}

func (x *Data_Gorm) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Gorm.ProtoReflect.Descriptor instead.
func (*Data_Gorm) Descriptor() ([]byte, []int) {
	return file_conf_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Data_Gorm) GetDataSourceName() string {
	if x != nil {
		return x.DataSourceName
	}
	return ""
}

func (x *Data_Gorm) GetMaxIdleConn() int32 {
	if x != nil {
		return x.MaxIdleConn
	}
	return 0
}

func (x *Data_Gorm) GetMaxOpenConn() int32 {
	if x != nil {
		return x.MaxOpenConn
	}
	return 0
}

func (x *Data_Gorm) GetConnMaxLifeTime() *durationpb.Duration {
	if x != nil {
		return x.ConnMaxLifeTime
	}
	return nil
}

func (x *Data_Gorm) GetShowLog() bool {
	if x != nil {
		return x.ShowLog
	}
	return false
}

func (x *Data_Gorm) GetTracing() bool {
	if x != nil {
		return x.Tracing
	}
	return false
}

// redis
type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`           // 网络
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`                 // 服务端地址
	Username     string               `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`         // 账号
	Password     string               `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`         // 密码
	Db           int32                `protobuf:"varint,5,opt,name=db,proto3" json:"db,omitempty"`                    // 数据库索引
	DialTimeout  *durationpb.Duration `protobuf:"bytes,6,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`   // 连接超时时间
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,7,opt,name=readTimeout,proto3" json:"readTimeout,omitempty"`   // 读取超时时间
	WriteTimeout *durationpb.Duration `protobuf:"bytes,8,opt,name=writeTimeout,proto3" json:"writeTimeout,omitempty"` // 写入超时时间
	Tracing      bool                 `protobuf:"varint,9,opt,name=tracing,proto3" json:"tracing,omitempty"`          // 链路追踪开关
	Metrics      bool                 `protobuf:"varint,10,opt,name=metrics,proto3" json:"metrics,omitempty"`         // 指标开关
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_v1_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Data_Redis) GetTracing() bool {
	if x != nil {
		return x.Tracing
	}
	return false
}

func (x *Data_Redis) GetMetrics() bool {
	if x != nil {
		return x.Metrics
	}
	return false
}

var File_conf_v1_data_proto protoreflect.FileDescriptor

var file_conf_v1_data_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x05, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x04, 0x67, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x6f,
	0x72, 0x6d, 0x52, 0x04, 0x67, 0x6f, 0x72, 0x6d, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x1a, 0xeb, 0x01, 0x0a, 0x04, 0x47, 0x6f, 0x72, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f,
	0x6e, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4f, 0x70, 0x65,
	0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x43, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78,
	0x4c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x4d,
	0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68,
	0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x68, 0x6f,
	0x77, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0xea,
	0x02, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x64, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x64, 0x62, 0x12, 0x3b,
	0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x72,
	0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x52, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x09, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0xca, 0x02,
	0x04, 0x43, 0x6f, 0x6e, 0x66, 0xe2, 0x02, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_data_proto_rawDescOnce sync.Once
	file_conf_v1_data_proto_rawDescData = file_conf_v1_data_proto_rawDesc
)

func file_conf_v1_data_proto_rawDescGZIP() []byte {
	file_conf_v1_data_proto_rawDescOnce.Do(func() {
		file_conf_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_data_proto_rawDescData)
	})
	return file_conf_v1_data_proto_rawDescData
}

var file_conf_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_conf_v1_data_proto_goTypes = []interface{}{
	(*Data)(nil),                // 0: conf.Data
	(*Data_Gorm)(nil),           // 1: conf.Data.Gorm
	(*Data_Redis)(nil),          // 2: conf.Data.Redis
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_conf_v1_data_proto_depIdxs = []int32{
	1, // 0: conf.Data.gorm:type_name -> conf.Data.Gorm
	2, // 1: conf.Data.redis:type_name -> conf.Data.Redis
	3, // 2: conf.Data.Gorm.connMaxLifeTime:type_name -> google.protobuf.Duration
	3, // 3: conf.Data.Redis.dialTimeout:type_name -> google.protobuf.Duration
	3, // 4: conf.Data.Redis.readTimeout:type_name -> google.protobuf.Duration
	3, // 5: conf.Data.Redis.writeTimeout:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_conf_v1_data_proto_init() }
func file_conf_v1_data_proto_init() {
	if File_conf_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_conf_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Gorm); i {
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
		file_conf_v1_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
			RawDescriptor: file_conf_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_data_proto_goTypes,
		DependencyIndexes: file_conf_v1_data_proto_depIdxs,
		MessageInfos:      file_conf_v1_data_proto_msgTypes,
	}.Build()
	File_conf_v1_data_proto = out.File
	file_conf_v1_data_proto_rawDesc = nil
	file_conf_v1_data_proto_goTypes = nil
	file_conf_v1_data_proto_depIdxs = nil
}