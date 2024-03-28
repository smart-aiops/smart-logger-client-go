// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.11.2
// source: grpc.proto

package proto

import (
	context "context"
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

type RawData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName  string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`     //应用名称
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`           //主机
	Port     int64  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`          //端口号
	ClientIp string `protobuf:"bytes,4,opt,name=clientIp,proto3" json:"clientIp,omitempty"`   //请求客户端的Ip
	ReqUrl   string `protobuf:"bytes,5,opt,name=reqUrl,proto3" json:"reqUrl,omitempty"`       //请求地址
	Headers  string `protobuf:"bytes,6,opt,name=headers,proto3" json:"headers,omitempty"`     //请求头部信息(可选择获取) 默认获取user-agent,content-type
	Tag      string `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag,omitempty"`             //操作标签,默认值undefined
	Content  string `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`     //方法步骤内容,默认是空,可使用LogData.step进行内容步骤记录, 当前只支持java sdk
	Method   string `protobuf:"bytes,9,opt,name=method,proto3" json:"method,omitempty"`       //请求的本地方法
	Args     string `protobuf:"bytes,10,opt,name=args,proto3" json:"args,omitempty"`          //方法请求参数
	RespBody string `protobuf:"bytes,11,opt,name=respBody,proto3" json:"respBody,omitempty"`  //方法响应参数
	CostTime int64  `protobuf:"varint,12,opt,name=costTime,proto3" json:"costTime,omitempty"` //整个方法内部耗时
	LogDate  int64  `protobuf:"varint,13,opt,name=logDate,proto3" json:"logDate,omitempty"`   //Log产生时间,LogData对象初始化的时间
	Success  bool   `protobuf:"varint,14,opt,name=success,proto3" json:"success,omitempty"`   //执行状态,成功(true)/异常(false)
	Account  string `protobuf:"bytes,15,opt,name=account,proto3" json:"account,omitempty"`    //操作用户名
	Level    int64  `protobuf:"varint,16,opt,name=level,proto3" json:"level,omitempty"`       //日志级别 0 一般 1 低危 2 中危 3 高危
}

func (x *RawData) Reset() {
	*x = RawData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawData) ProtoMessage() {}

func (x *RawData) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawData.ProtoReflect.Descriptor instead.
func (*RawData) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *RawData) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RawData) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RawData) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RawData) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *RawData) GetReqUrl() string {
	if x != nil {
		return x.ReqUrl
	}
	return ""
}

func (x *RawData) GetHeaders() string {
	if x != nil {
		return x.Headers
	}
	return ""
}

func (x *RawData) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *RawData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *RawData) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RawData) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *RawData) GetRespBody() string {
	if x != nil {
		return x.RespBody
	}
	return ""
}

func (x *RawData) GetCostTime() int64 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *RawData) GetLogDate() int64 {
	if x != nil {
		return x.LogDate
	}
	return 0
}

func (x *RawData) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RawData) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *RawData) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x07, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x71, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x71, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x42,
	0x6f, 0x64, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x40, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2d, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_goTypes = []interface{}{
	(*RawData)(nil),  // 0: proto.RawData
	(*Response)(nil), // 1: proto.Response
}
var file_grpc_proto_depIdxs = []int32{
	0, // 0: proto.TransferService.Transfer:input_type -> proto.RawData
	1, // 1: proto.TransferService.Transfer:output_type -> proto.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawData); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferServiceClient interface {
	Transfer(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Response, error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) Transfer(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.TransferService/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServiceServer is the server API for TransferService service.
type TransferServiceServer interface {
	Transfer(context.Context, *RawData) (*Response, error)
}

// UnimplementedTransferServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransferServiceServer struct {
}

func (*UnimplementedTransferServiceServer) Transfer(context.Context, *RawData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}

func RegisterTransferServiceServer(s *grpc.Server, srv TransferServiceServer) {
	s.RegisterService(&_TransferService_serviceDesc, srv)
}

func _TransferService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TransferService/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Transfer(ctx, req.(*RawData))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransferService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _TransferService_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
