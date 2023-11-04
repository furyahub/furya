// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfer/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("transfer/v1/query.proto", fileDescriptor_3f47107eac5604ce) }

var fileDescriptor_3f47107eac5604ce = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x57, 0x61, 0x03, 0xe3, 0x45, 0x72, 0x11, 0xc6, 0xc8, 0x61, 0xec, 0xa0, 0x53, 0x13,
	0xbb, 0x4d, 0x1f, 0x40, 0x3c, 0xe8, 0x4d, 0xc5, 0x93, 0x17, 0x49, 0x6b, 0x6c, 0x0b, 0x5b, 0x92,
	0x25, 0xe9, 0x70, 0xc8, 0x2e, 0x3e, 0x81, 0xb0, 0x57, 0xf0, 0x20, 0xe2, 0x83, 0x78, 0x1c, 0x78,
	0xf1, 0x28, 0x9b, 0x0f, 0x22, 0x4d, 0x5b, 0xa7, 0xe0, 0x9f, 0xf5, 0x56, 0xe8, 0xf7, 0x93, 0xef,
	0xe7, 0xf7, 0x4b, 0x0b, 0xd6, 0x8c, 0xa2, 0x5c, 0x5f, 0x31, 0x45, 0x06, 0x2e, 0xe9, 0xc7, 0x4c,
	0x0d, 0xb1, 0x54, 0xc2, 0x08, 0xb8, 0xca, 0x59, 0x6c, 0x94, 0xe0, 0x38, 0x0f, 0x54, 0xd7, 0x23,
	0xcf, 0x27, 0x54, 0xca, 0x6e, 0xe4, 0x53, 0x13, 0x09, 0xae, 0xc9, 0x2f, 0x6c, 0xb5, 0x16, 0x08,
	0x11, 0x74, 0x19, 0xa1, 0x32, 0x22, 0x94, 0x73, 0x61, 0xd2, 0x7c, 0xfa, 0xb6, 0x75, 0x5f, 0x06,
	0xe5, 0x93, 0x24, 0x0d, 0x9f, 0x1c, 0x00, 0x0e, 0x18, 0x17, 0xbd, 0x33, 0x45, 0x7d, 0x06, 0x3b,
	0x38, 0xf2, 0x7c, 0xfc, 0xb5, 0xe1, 0xb3, 0x1c, 0x0f, 0x5c, 0x6c, 0x99, 0x79, 0xfc, 0x94, 0xf5,
	0x63, 0xa6, 0x4d, 0x75, 0xb7, 0x20, 0xa5, 0xa5, 0xe0, 0x9a, 0xd5, 0xdd, 0xdb, 0x97, 0xf7, 0xf1,
	0xd2, 0x26, 0xdc, 0x20, 0xd9, 0x58, 0xdf, 0xc7, 0xb9, 0x4c, 0x88, 0x0b, 0x93, 0x20, 0x9a, 0xdc,
	0x84, 0x54, 0x87, 0x23, 0xf8, 0xe0, 0x80, 0x95, 0xf9, 0x49, 0x1a, 0x16, 0x6b, 0xd6, 0xb9, 0xf0,
	0x5e, 0x51, 0x2c, 0x33, 0x6e, 0x5a, 0xe3, 0x06, 0xac, 0xff, 0x6f, 0x0c, 0xc7, 0x0e, 0xa8, 0x1c,
	0x53, 0x45, 0x7b, 0x1a, 0xee, 0x2c, 0x50, 0x97, 0x46, 0x73, 0x41, 0xb7, 0x00, 0x91, 0xb9, 0x35,
	0xac, 0x1b, 0x82, 0xb5, 0x9f, 0xdd, 0x64, 0xaa, 0xf2, 0xe8, 0x80, 0x65, 0x3b, 0xd9, 0x21, 0xd5,
	0x21, 0x6c, 0x2f, 0xba, 0x87, 0x24, 0x9d, 0xbb, 0x75, 0x8a, 0x41, 0x99, 0x5e, 0xcb, 0xea, 0x6d,
	0xc1, 0xe6, 0x5f, 0xab, 0x4b, 0x2e, 0x39, 0xb9, 0x6c, 0xbb, 0xc2, 0xd1, 0xfe, 0xd1, 0xf3, 0x14,
	0x39, 0x93, 0x29, 0x72, 0xde, 0xa6, 0xc8, 0xb9, 0x9b, 0xa1, 0xd2, 0x64, 0x86, 0x4a, 0xaf, 0x33,
	0x54, 0x3a, 0x27, 0x41, 0x64, 0xc2, 0xd8, 0xc3, 0xbe, 0xe8, 0x91, 0xec, 0x2f, 0xd9, 0x16, 0x2a,
	0xc8, 0x9f, 0xc9, 0xf5, 0xfc, 0x70, 0x33, 0x94, 0x4c, 0x7b, 0x15, 0xfb, 0xe1, 0xb7, 0x3f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0x5f, 0xcb, 0x48, 0x6d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// DenomTrace queries a denomination trace information.
	DenomTrace(ctx context.Context, in *types.QueryDenomTraceRequest, opts ...grpc.CallOption) (*types.QueryDenomTraceResponse, error)
	// DenomTraces queries all denomination traces.
	DenomTraces(ctx context.Context, in *types.QueryDenomTracesRequest, opts ...grpc.CallOption) (*types.QueryDenomTracesResponse, error)
	// Params queries all parameters of the ibc-transfer module.
	Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error)
	// DenomHash queries a denomination hash information.
	DenomHash(ctx context.Context, in *types.QueryDenomHashRequest, opts ...grpc.CallOption) (*types.QueryDenomHashResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DenomTrace(ctx context.Context, in *types.QueryDenomTraceRequest, opts ...grpc.CallOption) (*types.QueryDenomTraceResponse, error) {
	out := new(types.QueryDenomTraceResponse)
	err := c.cc.Invoke(ctx, "/furya.transfer.Query/DenomTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomTraces(ctx context.Context, in *types.QueryDenomTracesRequest, opts ...grpc.CallOption) (*types.QueryDenomTracesResponse, error) {
	out := new(types.QueryDenomTracesResponse)
	err := c.cc.Invoke(ctx, "/furya.transfer.Query/DenomTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	out := new(types.QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/furya.transfer.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomHash(ctx context.Context, in *types.QueryDenomHashRequest, opts ...grpc.CallOption) (*types.QueryDenomHashResponse, error) {
	out := new(types.QueryDenomHashResponse)
	err := c.cc.Invoke(ctx, "/furya.transfer.Query/DenomHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DenomTrace queries a denomination trace information.
	DenomTrace(context.Context, *types.QueryDenomTraceRequest) (*types.QueryDenomTraceResponse, error)
	// DenomTraces queries all denomination traces.
	DenomTraces(context.Context, *types.QueryDenomTracesRequest) (*types.QueryDenomTracesResponse, error)
	// Params queries all parameters of the ibc-transfer module.
	Params(context.Context, *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
	// DenomHash queries a denomination hash information.
	DenomHash(context.Context, *types.QueryDenomHashRequest) (*types.QueryDenomHashResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DenomTrace(ctx context.Context, req *types.QueryDenomTraceRequest) (*types.QueryDenomTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTrace not implemented")
}
func (*UnimplementedQueryServer) DenomTraces(ctx context.Context, req *types.QueryDenomTracesRequest) (*types.QueryDenomTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTraces not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DenomHash(ctx context.Context, req *types.QueryDenomHashRequest) (*types.QueryDenomHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomHash not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DenomTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furya.transfer.Query/DenomTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTrace(ctx, req.(*types.QueryDenomTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furya.transfer.Query/DenomTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTraces(ctx, req.(*types.QueryDenomTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furya.transfer.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*types.QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furya.transfer.Query/DenomHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomHash(ctx, req.(*types.QueryDenomHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "furya.transfer.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DenomTrace",
			Handler:    _Query_DenomTrace_Handler,
		},
		{
			MethodName: "DenomTraces",
			Handler:    _Query_DenomTraces_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DenomHash",
			Handler:    _Query_DenomHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer/v1/query.proto",
}
