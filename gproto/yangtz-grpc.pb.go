// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.3
// source: yangtz-grpc.proto

// 定义包名

package gproto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 定义 Req 消息结构
type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 类型 字段 = 标识号
	JsonStr string `protobuf:"bytes,1,opt,name=JsonStr,proto3" json:"JsonStr,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yangtz_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_yangtz_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_yangtz_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetJsonStr() string {
	if x != nil {
		return x.JsonStr
	}
	return ""
}

// 定义 Res 消息结构
type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status_Txt  string `protobuf:"bytes,1,opt,name=Status_Txt,json=StatusTxt,proto3" json:"Status_Txt,omitempty"`
	Status_Code int32  `protobuf:"varint,2,opt,name=Status_Code,json=StatusCode,proto3" json:"Status_Code,omitempty"`
	Data        string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yangtz_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_yangtz_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_yangtz_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetStatus_Txt() string {
	if x != nil {
		return x.Status_Txt
	}
	return ""
}

func (x *Res) GetStatus_Code() int32 {
	if x != nil {
		return x.Status_Code
	}
	return 0
}

func (x *Res) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_yangtz_grpc_proto protoreflect.FileDescriptor

var file_yangtz_grpc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x79, 0x61, 0x6e, 0x67, 0x74, 0x7a, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x03, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x22, 0x59, 0x0a, 0x03,
	0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x54, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54,
	0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0xaa, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x0b, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0b, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x6f, 0x70, 0x12,
	0x0b, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x09, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yangtz_grpc_proto_rawDescOnce sync.Once
	file_yangtz_grpc_proto_rawDescData = file_yangtz_grpc_proto_rawDesc
)

func file_yangtz_grpc_proto_rawDescGZIP() []byte {
	file_yangtz_grpc_proto_rawDescOnce.Do(func() {
		file_yangtz_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_yangtz_grpc_proto_rawDescData)
	})
	return file_yangtz_grpc_proto_rawDescData
}

var file_yangtz_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yangtz_grpc_proto_goTypes = []interface{}{
	(*Req)(nil), // 0: gproto.Req
	(*Res)(nil), // 1: gproto.Res
}
var file_yangtz_grpc_proto_depIdxs = []int32{
	0, // 0: gproto.WorkerServer.Ping:input_type -> gproto.Req
	0, // 1: gproto.WorkerServer.JobStart:input_type -> gproto.Req
	0, // 2: gproto.WorkerServer.JobStop:input_type -> gproto.Req
	0, // 3: gproto.WorkerServer.JobStatus:input_type -> gproto.Req
	1, // 4: gproto.WorkerServer.Ping:output_type -> gproto.Res
	1, // 5: gproto.WorkerServer.JobStart:output_type -> gproto.Res
	1, // 6: gproto.WorkerServer.JobStop:output_type -> gproto.Res
	1, // 7: gproto.WorkerServer.JobStatus:output_type -> gproto.Res
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yangtz_grpc_proto_init() }
func file_yangtz_grpc_proto_init() {
	if File_yangtz_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yangtz_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_yangtz_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_yangtz_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yangtz_grpc_proto_goTypes,
		DependencyIndexes: file_yangtz_grpc_proto_depIdxs,
		MessageInfos:      file_yangtz_grpc_proto_msgTypes,
	}.Build()
	File_yangtz_grpc_proto = out.File
	file_yangtz_grpc_proto_rawDesc = nil
	file_yangtz_grpc_proto_goTypes = nil
	file_yangtz_grpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkerServerClient is the client API for WorkerServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerServerClient interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	Ping(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	JobStart(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	JobStop(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	JobStatus(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type workerServerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServerClient(cc grpc.ClientConnInterface) WorkerServerClient {
	return &workerServerClient{cc}
}

func (c *workerServerClient) Ping(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gproto.WorkerServer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServerClient) JobStart(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gproto.WorkerServer/JobStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServerClient) JobStop(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gproto.WorkerServer/JobStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServerClient) JobStatus(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gproto.WorkerServer/JobStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServerServer is the server API for WorkerServer service.
type WorkerServerServer interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	Ping(context.Context, *Req) (*Res, error)
	JobStart(context.Context, *Req) (*Res, error)
	JobStop(context.Context, *Req) (*Res, error)
	JobStatus(context.Context, *Req) (*Res, error)
}

// UnimplementedWorkerServerServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServerServer struct {
}

func (*UnimplementedWorkerServerServer) Ping(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedWorkerServerServer) JobStart(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobStart not implemented")
}
func (*UnimplementedWorkerServerServer) JobStop(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobStop not implemented")
}
func (*UnimplementedWorkerServerServer) JobStatus(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobStatus not implemented")
}

func RegisterWorkerServerServer(s *grpc.Server, srv WorkerServerServer) {
	s.RegisterService(&_WorkerServer_serviceDesc, srv)
}

func _WorkerServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.WorkerServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServerServer).Ping(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerServer_JobStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServerServer).JobStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.WorkerServer/JobStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServerServer).JobStart(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerServer_JobStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServerServer).JobStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.WorkerServer/JobStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServerServer).JobStop(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerServer_JobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServerServer).JobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.WorkerServer/JobStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServerServer).JobStatus(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gproto.WorkerServer",
	HandlerType: (*WorkerServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _WorkerServer_Ping_Handler,
		},
		{
			MethodName: "JobStart",
			Handler:    _WorkerServer_JobStart_Handler,
		},
		{
			MethodName: "JobStop",
			Handler:    _WorkerServer_JobStop_Handler,
		},
		{
			MethodName: "JobStatus",
			Handler:    _WorkerServer_JobStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yangtz-grpc.proto",
}