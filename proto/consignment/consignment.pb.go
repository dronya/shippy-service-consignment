// shippy-service-consignment/proto/consignment/consignment.proto
//protoc -I. --go_out=plugins=grpc:. consignment.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.7
// source: consignment.proto

package consignment

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

type Consignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
}

func (x *Consignment) Reset() {
	*x = Consignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consignment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consignment) ProtoMessage() {}

func (x *Consignment) ProtoReflect() protoreflect.Message {
	mi := &file_consignment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consignment.ProtoReflect.Descriptor instead.
func (*Consignment) Descriptor() ([]byte, []int) {
	return file_consignment_proto_rawDescGZIP(), []int{0}
}

func (x *Consignment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Consignment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Consignment) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Consignment) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *Consignment) GetVesselId() string {
	if x != nil {
		return x.VesselId
	}
	return ""
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consignment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_consignment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_consignment_proto_rawDescGZIP(), []int{1}
}

func (x *Container) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Container) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Container) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Container) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Created a blank get request
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consignment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consignment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_consignment_proto_rawDescGZIP(), []int{2}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created     bool         `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment *Consignment `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	// Added a pluralised consignment to our generic response message
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consignment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_consignment_proto_msgTypes[3]
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
	return file_consignment_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *Response) GetConsignment() *Consignment {
	if x != nil {
		return x.Consignment
	}
	return nil
}

func (x *Response) GetConsignments() []*Consignment {
	if x != nil {
		return x.Consignments
	}
	return nil
}

var File_consignment_proto protoreflect.FileDescriptor

var file_consignment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0xac, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x49, 0x64, 0x22,
	0x6d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0c,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9e, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x3c, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x9e, 0x01,
	0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consignment_proto_rawDescOnce sync.Once
	file_consignment_proto_rawDescData = file_consignment_proto_rawDesc
)

func file_consignment_proto_rawDescGZIP() []byte {
	file_consignment_proto_rawDescOnce.Do(func() {
		file_consignment_proto_rawDescData = protoimpl.X.CompressGZIP(file_consignment_proto_rawDescData)
	})
	return file_consignment_proto_rawDescData
}

var file_consignment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_consignment_proto_goTypes = []interface{}{
	(*Consignment)(nil), // 0: consignment.Consignment
	(*Container)(nil),   // 1: consignment.Container
	(*GetRequest)(nil),  // 2: consignment.GetRequest
	(*Response)(nil),    // 3: consignment.Response
}
var file_consignment_proto_depIdxs = []int32{
	1, // 0: consignment.Consignment.containers:type_name -> consignment.Container
	0, // 1: consignment.Response.consignment:type_name -> consignment.Consignment
	0, // 2: consignment.Response.consignments:type_name -> consignment.Consignment
	0, // 3: consignment.ShippingService.CreateConsignment:input_type -> consignment.Consignment
	2, // 4: consignment.ShippingService.GetConsignments:input_type -> consignment.GetRequest
	3, // 5: consignment.ShippingService.CreateConsignment:output_type -> consignment.Response
	3, // 6: consignment.ShippingService.GetConsignments:output_type -> consignment.Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_consignment_proto_init() }
func file_consignment_proto_init() {
	if File_consignment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consignment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consignment); i {
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
		file_consignment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_consignment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_consignment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_consignment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consignment_proto_goTypes,
		DependencyIndexes: file_consignment_proto_depIdxs,
		MessageInfos:      file_consignment_proto_msgTypes,
	}.Build()
	File_consignment_proto = out.File
	file_consignment_proto_rawDesc = nil
	file_consignment_proto_goTypes = nil
	file_consignment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
	// Created a new method
	GetConsignments(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.ShippingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.ShippingService/GetConsignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
	// Created a new method
	GetConsignments(context.Context, *GetRequest) (*Response, error)
}

// UnimplementedShippingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShippingServiceServer struct {
}

func (*UnimplementedShippingServiceServer) CreateConsignment(context.Context, *Consignment) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsignment not implemented")
}
func (*UnimplementedShippingServiceServer) GetConsignments(context.Context, *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsignments not implemented")
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetConsignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetConsignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.ShippingService/GetConsignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetConsignments(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
		{
			MethodName: "GetConsignments",
			Handler:    _ShippingService_GetConsignments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consignment.proto",
}
