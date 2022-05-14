// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: internal/grpc/domain/repo.proto

package domain

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

type RepositoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId    int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	IsPrivate bool   `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
}

func (x *RepositoryRequest) Reset() {
	*x = RepositoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_domain_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryRequest) ProtoMessage() {}

func (x *RepositoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_domain_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryRequest.ProtoReflect.Descriptor instead.
func (*RepositoryRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_domain_repo_proto_rawDescGZIP(), []int{0}
}

func (x *RepositoryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RepositoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoryRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RepositoryRequest) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type AddRepositoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddedRepository *RepositoryRequest `protobuf:"bytes,1,opt,name=addedRepository,proto3" json:"addedRepository,omitempty"`
	Error           *Error             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddRepositoryResponse) Reset() {
	*x = AddRepositoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_domain_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRepositoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRepositoryResponse) ProtoMessage() {}

func (x *AddRepositoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_domain_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRepositoryResponse.ProtoReflect.Descriptor instead.
func (*AddRepositoryResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_domain_repo_proto_rawDescGZIP(), []int{1}
}

func (x *AddRepositoryResponse) GetAddedRepository() *RepositoryRequest {
	if x != nil {
		return x.AddedRepository
	}
	return nil
}

func (x *AddRepositoryResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_domain_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_domain_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_internal_grpc_domain_repo_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_grpc_domain_repo_proto protoreflect.FileDescriptor

var file_internal_grpc_domain_repo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x6d, 0x0a, 0x11, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x54, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x19, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_grpc_domain_repo_proto_rawDescOnce sync.Once
	file_internal_grpc_domain_repo_proto_rawDescData = file_internal_grpc_domain_repo_proto_rawDesc
)

func file_internal_grpc_domain_repo_proto_rawDescGZIP() []byte {
	file_internal_grpc_domain_repo_proto_rawDescOnce.Do(func() {
		file_internal_grpc_domain_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_domain_repo_proto_rawDescData)
	})
	return file_internal_grpc_domain_repo_proto_rawDescData
}

var file_internal_grpc_domain_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_grpc_domain_repo_proto_goTypes = []interface{}{
	(*RepositoryRequest)(nil),     // 0: domain.RepositoryRequest
	(*AddRepositoryResponse)(nil), // 1: domain.AddRepositoryResponse
	(*Error)(nil),                 // 2: domain.Error
}
var file_internal_grpc_domain_repo_proto_depIdxs = []int32{
	0, // 0: domain.AddRepositoryResponse.addedRepository:type_name -> domain.RepositoryRequest
	2, // 1: domain.AddRepositoryResponse.error:type_name -> domain.Error
	0, // 2: domain.ResositoryService.Add:input_type -> domain.RepositoryRequest
	1, // 3: domain.ResositoryService.Add:output_type -> domain.AddRepositoryResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_grpc_domain_repo_proto_init() }
func file_internal_grpc_domain_repo_proto_init() {
	if File_internal_grpc_domain_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_domain_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryRequest); i {
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
		file_internal_grpc_domain_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRepositoryResponse); i {
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
		file_internal_grpc_domain_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_internal_grpc_domain_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_domain_repo_proto_goTypes,
		DependencyIndexes: file_internal_grpc_domain_repo_proto_depIdxs,
		MessageInfos:      file_internal_grpc_domain_repo_proto_msgTypes,
	}.Build()
	File_internal_grpc_domain_repo_proto = out.File
	file_internal_grpc_domain_repo_proto_rawDesc = nil
	file_internal_grpc_domain_repo_proto_goTypes = nil
	file_internal_grpc_domain_repo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResositoryServiceClient is the client API for ResositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResositoryServiceClient interface {
	Add(ctx context.Context, in *RepositoryRequest, opts ...grpc.CallOption) (*AddRepositoryResponse, error)
}

type resositoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResositoryServiceClient(cc grpc.ClientConnInterface) ResositoryServiceClient {
	return &resositoryServiceClient{cc}
}

func (c *resositoryServiceClient) Add(ctx context.Context, in *RepositoryRequest, opts ...grpc.CallOption) (*AddRepositoryResponse, error) {
	out := new(AddRepositoryResponse)
	err := c.cc.Invoke(ctx, "/domain.ResositoryService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResositoryServiceServer is the server API for ResositoryService service.
type ResositoryServiceServer interface {
	Add(context.Context, *RepositoryRequest) (*AddRepositoryResponse, error)
}

// UnimplementedResositoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResositoryServiceServer struct {
}

func (*UnimplementedResositoryServiceServer) Add(context.Context, *RepositoryRequest) (*AddRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterResositoryServiceServer(s *grpc.Server, srv ResositoryServiceServer) {
	s.RegisterService(&_ResositoryService_serviceDesc, srv)
}

func _ResositoryService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResositoryServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.ResositoryService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResositoryServiceServer).Add(ctx, req.(*RepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "domain.ResositoryService",
	HandlerType: (*ResositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ResositoryService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/grpc/domain/repo.proto",
}
