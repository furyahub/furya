// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchaintxs/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/furyahub/furya/x/feerefunder/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// MsgRegisterInterchainAccount is used to register an account on a remote zone.
type MsgRegisterInterchainAccount struct {
	FromAddress         string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ConnectionId        string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
	InterchainAccountId string `protobuf:"bytes,3,opt,name=interchain_account_id,json=interchainAccountId,proto3" json:"interchain_account_id,omitempty" yaml:"interchain_account_id"`
}

func (m *MsgRegisterInterchainAccount) Reset()         { *m = MsgRegisterInterchainAccount{} }
func (m *MsgRegisterInterchainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterInterchainAccount) ProtoMessage()    {}
func (*MsgRegisterInterchainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd987b66c8800e1, []int{0}
}
func (m *MsgRegisterInterchainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterInterchainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterInterchainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterInterchainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterInterchainAccount.Merge(m, src)
}
func (m *MsgRegisterInterchainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterInterchainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterInterchainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterInterchainAccount proto.InternalMessageInfo

// MsgRegisterInterchainAccountResponse is the response type for
// MsgRegisterInterchainAccount.
type MsgRegisterInterchainAccountResponse struct {
}

func (m *MsgRegisterInterchainAccountResponse) Reset()         { *m = MsgRegisterInterchainAccountResponse{} }
func (m *MsgRegisterInterchainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterInterchainAccountResponse) ProtoMessage()    {}
func (*MsgRegisterInterchainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd987b66c8800e1, []int{1}
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterInterchainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterInterchainAccountResponse.Merge(m, src)
}
func (m *MsgRegisterInterchainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterInterchainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterInterchainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterInterchainAccountResponse proto.InternalMessageInfo

// MsgSubmitTx defines the payload for Msg/SubmitTx
type MsgSubmitTx struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// interchain_account_id is supposed to be the unique identifier, e.g.,
	// lido/kava. This allows contracts to have more than one interchain accounts
	// on remote zone This identifier will be a part of the portID that we'll
	// claim our capability for.
	InterchainAccountId string       `protobuf:"bytes,2,opt,name=interchain_account_id,json=interchainAccountId,proto3" json:"interchain_account_id,omitempty"`
	ConnectionId        string       `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Msgs                []*types.Any `protobuf:"bytes,4,rep,name=msgs,proto3" json:"msgs,omitempty"`
	Memo                string       `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	// timeout in seconds after which the packet times out
	Timeout uint64     `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Fee     types1.Fee `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee"`
}

func (m *MsgSubmitTx) Reset()         { *m = MsgSubmitTx{} }
func (m *MsgSubmitTx) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitTx) ProtoMessage()    {}
func (*MsgSubmitTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd987b66c8800e1, []int{2}
}
func (m *MsgSubmitTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitTx.Merge(m, src)
}
func (m *MsgSubmitTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitTx proto.InternalMessageInfo

// MsgSubmitTxResponse defines the response for Msg/SubmitTx
type MsgSubmitTxResponse struct {
	// channel's sequence_id for outgoing ibc packet. Unique per a channel.
	SequenceId uint64 `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// channel src channel on furya side trasaction was submitted from
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (m *MsgSubmitTxResponse) Reset()         { *m = MsgSubmitTxResponse{} }
func (m *MsgSubmitTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitTxResponse) ProtoMessage()    {}
func (*MsgSubmitTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecd987b66c8800e1, []int{3}
}
func (m *MsgSubmitTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitTxResponse.Merge(m, src)
}
func (m *MsgSubmitTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitTxResponse proto.InternalMessageInfo

func (m *MsgSubmitTxResponse) GetSequenceId() uint64 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *MsgSubmitTxResponse) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRegisterInterchainAccount)(nil), "furya.interchaintxs.v1.MsgRegisterInterchainAccount")
	proto.RegisterType((*MsgRegisterInterchainAccountResponse)(nil), "furya.interchaintxs.v1.MsgRegisterInterchainAccountResponse")
	proto.RegisterType((*MsgSubmitTx)(nil), "furya.interchaintxs.v1.MsgSubmitTx")
	proto.RegisterType((*MsgSubmitTxResponse)(nil), "furya.interchaintxs.v1.MsgSubmitTxResponse")
}

func init() { proto.RegisterFile("interchaintxs/v1/tx.proto", fileDescriptor_ecd987b66c8800e1) }

var fileDescriptor_ecd987b66c8800e1 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xbd, 0x6e, 0xdb, 0x3c,
	0x14, 0x95, 0x62, 0x7f, 0x49, 0x3e, 0x3a, 0x5d, 0x98, 0x04, 0x90, 0x0d, 0x43, 0x72, 0xd5, 0x1f,
	0x78, 0x09, 0xd5, 0xb8, 0x45, 0x87, 0x00, 0x2d, 0x60, 0x0f, 0x05, 0x3c, 0xb8, 0x28, 0xd4, 0x4c,
	0x5d, 0x0c, 0x59, 0xba, 0xa2, 0x05, 0x58, 0xa4, 0x2b, 0x52, 0x81, 0x3d, 0x76, 0xeb, 0xd8, 0xa5,
	0x63, 0x81, 0x3c, 0x4e, 0xc6, 0x8c, 0x9d, 0x8c, 0xc0, 0x5e, 0x3a, 0xe7, 0x09, 0x0a, 0x49, 0x96,
	0xff, 0x10, 0x1b, 0x41, 0x37, 0xde, 0xc3, 0xc3, 0xc3, 0x7b, 0x0f, 0x8f, 0x84, 0xca, 0x01, 0x93,
	0x10, 0xb9, 0x7d, 0x27, 0x60, 0x72, 0x24, 0xac, 0xab, 0x73, 0x4b, 0x8e, 0xc8, 0x30, 0xe2, 0x92,
	0x63, 0x8d, 0x41, 0x2c, 0x23, 0xce, 0xc8, 0x1a, 0x85, 0x5c, 0x9d, 0x57, 0xca, 0x2e, 0x17, 0x21,
	0x17, 0xdd, 0x94, 0x67, 0x65, 0x45, 0x76, 0xa8, 0x72, 0x42, 0x39, 0xe5, 0x19, 0x9e, 0xac, 0xe6,
	0xe8, 0x29, 0xe5, 0x9c, 0x0e, 0xc0, 0x72, 0x86, 0x81, 0xd5, 0x97, 0x72, 0x38, 0x87, 0xab, 0x2b,
	0xb0, 0xc3, 0x18, 0x97, 0x8e, 0x0c, 0x38, 0xcb, 0xa5, 0xca, 0xf3, 0xdd, 0xb4, 0xea, 0xc5, 0xbe,
	0xe5, 0xb0, 0x71, 0xae, 0xe7, 0x03, 0x44, 0xe0, 0xc7, 0xcc, 0x83, 0xc8, 0xf2, 0x01, 0x32, 0xd8,
	0xbc, 0x53, 0x51, 0xb5, 0x23, 0xa8, 0x0d, 0x34, 0x10, 0x12, 0xa2, 0xf6, 0xa2, 0xef, 0xa6, 0xeb,
	0xf2, 0x98, 0x49, 0xfc, 0x14, 0x1d, 0xf9, 0x11, 0x0f, 0xbb, 0x8e, 0xe7, 0x45, 0x20, 0x84, 0xa6,
	0xd6, 0xd4, 0xfa, 0xff, 0x76, 0x29, 0xc1, 0x9a, 0x19, 0x84, 0xdf, 0xa1, 0x27, 0x2e, 0x67, 0x0c,
	0xdc, 0xa4, 0x95, 0x6e, 0xe0, 0x69, 0x7b, 0x09, 0xa7, 0xa5, 0xdd, 0x4f, 0x8c, 0x93, 0xb1, 0x13,
	0x0e, 0x2e, 0xcc, 0xb5, 0x6d, 0xd3, 0x3e, 0x5a, 0xd6, 0x6d, 0x0f, 0x5f, 0xa2, 0xd3, 0xa5, 0x5d,
	0x5d, 0x27, 0xbb, 0x37, 0x91, 0x29, 0xa4, 0x32, 0xb5, 0xfb, 0x89, 0x51, 0xcd, 0x64, 0x1e, 0xa4,
	0x99, 0xf6, 0x71, 0xb0, 0xd9, 0x75, 0xdb, 0xbb, 0x38, 0xfc, 0x7e, 0x6d, 0x28, 0x7f, 0xae, 0x0d,
	0xc5, 0x7c, 0x89, 0x9e, 0xef, 0x9a, 0xd0, 0x06, 0x31, 0xe4, 0x4c, 0x80, 0xf9, 0x6b, 0x0f, 0x95,
	0x3a, 0x82, 0x7e, 0x8e, 0x7b, 0x61, 0x20, 0x2f, 0x47, 0x8f, 0x99, 0xbc, 0xb1, 0xad, 0xf5, 0xd4,
	0x81, 0x07, 0x1b, 0xc3, 0xcf, 0x36, 0xdd, 0x4a, 0xc7, 0xdc, 0xf0, 0xa4, 0x8e, 0x8a, 0xa1, 0xa0,
	0x42, 0x2b, 0xd6, 0x0a, 0xf5, 0x52, 0xe3, 0x84, 0x64, 0xef, 0x4a, 0xf2, 0x77, 0x25, 0x4d, 0x36,
	0xb6, 0x53, 0x06, 0xc6, 0xa8, 0x18, 0x42, 0xc8, 0xb5, 0xff, 0x52, 0x95, 0x74, 0x8d, 0x35, 0x74,
	0x20, 0x83, 0x10, 0x78, 0x2c, 0xb5, 0xfd, 0x9a, 0x5a, 0x2f, 0xda, 0x79, 0x89, 0x5f, 0xa1, 0x82,
	0x0f, 0xa0, 0x1d, 0xd4, 0xd4, 0x7a, 0xa9, 0xa1, 0x91, 0x3c, 0xae, 0x2b, 0xd9, 0x20, 0x1f, 0x00,
	0x5a, 0xc5, 0x9b, 0x89, 0xa1, 0xd8, 0x09, 0x75, 0xc5, 0xc7, 0x4f, 0xe8, 0x78, 0xc5, 0x9e, 0xdc,
	0x36, 0x6c, 0xa0, 0x92, 0x80, 0xaf, 0x31, 0x30, 0x17, 0x92, 0x69, 0xd4, 0xf4, 0x42, 0x94, 0x43,
	0x6d, 0x2f, 0xe9, 0xc6, 0xed, 0x3b, 0x8c, 0xc1, 0x60, 0x6e, 0x4b, 0x5e, 0x36, 0xbe, 0xed, 0xa1,
	0x42, 0x47, 0x50, 0xfc, 0x53, 0x45, 0xe5, 0xed, 0x09, 0x7c, 0x4b, 0xb6, 0x7d, 0x55, 0x64, 0xd7,
	0xbb, 0x56, 0xde, 0xff, 0xdb, 0xb9, 0x45, 0x1e, 0x14, 0xdc, 0x43, 0x87, 0x8b, 0x34, 0xbc, 0xd8,
	0xa9, 0x96, 0xd3, 0x2a, 0x67, 0x8f, 0xa2, 0x2d, 0xef, 0x68, 0x7d, 0xbc, 0x99, 0xea, 0xea, 0xed,
	0x54, 0x57, 0xef, 0xa6, 0xba, 0xfa, 0x63, 0xa6, 0x2b, 0xb7, 0x33, 0x5d, 0xf9, 0x3d, 0xd3, 0x95,
	0x2f, 0x6f, 0x68, 0x20, 0xfb, 0x71, 0x8f, 0xb8, 0x3c, 0xb4, 0xe6, 0xa2, 0x67, 0x3c, 0xa2, 0xf9,
	0xda, 0x1a, 0x59, 0xeb, 0x3f, 0x22, 0x39, 0x1e, 0x82, 0xe8, 0xed, 0xa7, 0x19, 0x79, 0xfd, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0xf6, 0x21, 0x72, 0xa6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RegisterInterchainAccount(ctx context.Context, in *MsgRegisterInterchainAccount, opts ...grpc.CallOption) (*MsgRegisterInterchainAccountResponse, error)
	SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterInterchainAccount(ctx context.Context, in *MsgRegisterInterchainAccount, opts ...grpc.CallOption) (*MsgRegisterInterchainAccountResponse, error) {
	out := new(MsgRegisterInterchainAccountResponse)
	err := c.cc.Invoke(ctx, "/furya.interchaintxs.v1.Msg/RegisterInterchainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error) {
	out := new(MsgSubmitTxResponse)
	err := c.cc.Invoke(ctx, "/furya.interchaintxs.v1.Msg/SubmitTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterInterchainAccount(context.Context, *MsgRegisterInterchainAccount) (*MsgRegisterInterchainAccountResponse, error)
	SubmitTx(context.Context, *MsgSubmitTx) (*MsgSubmitTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterInterchainAccount(ctx context.Context, req *MsgRegisterInterchainAccount) (*MsgRegisterInterchainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInterchainAccount not implemented")
}
func (*UnimplementedMsgServer) SubmitTx(ctx context.Context, req *MsgSubmitTx) (*MsgSubmitTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterInterchainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterInterchainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterInterchainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furya.interchaintxs.v1.Msg/RegisterInterchainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterInterchainAccount(ctx, req.(*MsgRegisterInterchainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furya.interchaintxs.v1.Msg/SubmitTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitTx(ctx, req.(*MsgSubmitTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "furya.interchaintxs.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInterchainAccount",
			Handler:    _Msg_RegisterInterchainAccount_Handler,
		},
		{
			MethodName: "SubmitTx",
			Handler:    _Msg_SubmitTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interchaintxs/v1/tx.proto",
}

func (m *MsgRegisterInterchainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterInterchainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterInterchainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterchainAccountId) > 0 {
		i -= len(m.InterchainAccountId)
		copy(dAtA[i:], m.InterchainAccountId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InterchainAccountId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterInterchainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterInterchainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterInterchainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Timeout != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InterchainAccountId) > 0 {
		i -= len(m.InterchainAccountId)
		copy(dAtA[i:], m.InterchainAccountId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InterchainAccountId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if m.SequenceId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SequenceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterInterchainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InterchainAccountId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterInterchainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InterchainAccountId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovTx(uint64(m.Timeout))
	}
	l = m.Fee.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSubmitTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SequenceId != 0 {
		n += 1 + sovTx(uint64(m.SequenceId))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterInterchainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterInterchainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterInterchainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterchainAccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterchainAccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRegisterInterchainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterInterchainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterInterchainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterchainAccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterchainAccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &types.Any{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceId", wireType)
			}
			m.SequenceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
