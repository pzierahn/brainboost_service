// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: braingain.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Braingain_Chat_FullMethodName          = "/endpoint.braingain.v1.Braingain/Chat"
	Braingain_ListDocuments_FullMethodName = "/endpoint.braingain.v1.Braingain/ListDocuments"
)

// BraingainClient is the client API for Braingain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BraingainClient interface {
	Chat(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Completion, error)
	ListDocuments(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Documents, error)
}

type braingainClient struct {
	cc grpc.ClientConnInterface
}

func NewBraingainClient(cc grpc.ClientConnInterface) BraingainClient {
	return &braingainClient{cc}
}

func (c *braingainClient) Chat(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Completion, error) {
	out := new(Completion)
	err := c.cc.Invoke(ctx, Braingain_Chat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *braingainClient) ListDocuments(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Documents, error) {
	out := new(Documents)
	err := c.cc.Invoke(ctx, Braingain_ListDocuments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BraingainServer is the server API for Braingain service.
// All implementations must embed UnimplementedBraingainServer
// for forward compatibility
type BraingainServer interface {
	Chat(context.Context, *Prompt) (*Completion, error)
	ListDocuments(context.Context, *emptypb.Empty) (*Documents, error)
	mustEmbedUnimplementedBraingainServer()
}

// UnimplementedBraingainServer must be embedded to have forward compatible implementations.
type UnimplementedBraingainServer struct {
}

func (UnimplementedBraingainServer) Chat(context.Context, *Prompt) (*Completion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedBraingainServer) ListDocuments(context.Context, *emptypb.Empty) (*Documents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}
func (UnimplementedBraingainServer) mustEmbedUnimplementedBraingainServer() {}

// UnsafeBraingainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BraingainServer will
// result in compilation errors.
type UnsafeBraingainServer interface {
	mustEmbedUnimplementedBraingainServer()
}

func RegisterBraingainServer(s grpc.ServiceRegistrar, srv BraingainServer) {
	s.RegisterService(&Braingain_ServiceDesc, srv)
}

func _Braingain_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BraingainServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Braingain_Chat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BraingainServer).Chat(ctx, req.(*Prompt))
	}
	return interceptor(ctx, in, info, handler)
}

func _Braingain_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BraingainServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Braingain_ListDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BraingainServer).ListDocuments(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Braingain_ServiceDesc is the grpc.ServiceDesc for Braingain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Braingain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.braingain.v1.Braingain",
	HandlerType: (*BraingainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chat",
			Handler:    _Braingain_Chat_Handler,
		},
		{
			MethodName: "ListDocuments",
			Handler:    _Braingain_ListDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "braingain.proto",
}
