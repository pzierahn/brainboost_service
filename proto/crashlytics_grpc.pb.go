// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: crashlytics.proto

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
	CrashlyticsService_RecordError_FullMethodName = "/crashlytics.v1.CrashlyticsService/RecordError"
)

// CrashlyticsServiceClient is the client API for CrashlyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrashlyticsServiceClient interface {
	RecordError(ctx context.Context, in *Error, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type crashlyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrashlyticsServiceClient(cc grpc.ClientConnInterface) CrashlyticsServiceClient {
	return &crashlyticsServiceClient{cc}
}

func (c *crashlyticsServiceClient) RecordError(ctx context.Context, in *Error, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CrashlyticsService_RecordError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrashlyticsServiceServer is the server API for CrashlyticsService service.
// All implementations must embed UnimplementedCrashlyticsServiceServer
// for forward compatibility
type CrashlyticsServiceServer interface {
	RecordError(context.Context, *Error) (*emptypb.Empty, error)
	mustEmbedUnimplementedCrashlyticsServiceServer()
}

// UnimplementedCrashlyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrashlyticsServiceServer struct {
}

func (UnimplementedCrashlyticsServiceServer) RecordError(context.Context, *Error) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordError not implemented")
}
func (UnimplementedCrashlyticsServiceServer) mustEmbedUnimplementedCrashlyticsServiceServer() {}

// UnsafeCrashlyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrashlyticsServiceServer will
// result in compilation errors.
type UnsafeCrashlyticsServiceServer interface {
	mustEmbedUnimplementedCrashlyticsServiceServer()
}

func RegisterCrashlyticsServiceServer(s grpc.ServiceRegistrar, srv CrashlyticsServiceServer) {
	s.RegisterService(&CrashlyticsService_ServiceDesc, srv)
}

func _CrashlyticsService_RecordError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Error)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrashlyticsServiceServer).RecordError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrashlyticsService_RecordError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrashlyticsServiceServer).RecordError(ctx, req.(*Error))
	}
	return interceptor(ctx, in, info, handler)
}

// CrashlyticsService_ServiceDesc is the grpc.ServiceDesc for CrashlyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrashlyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crashlytics.v1.CrashlyticsService",
	HandlerType: (*CrashlyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordError",
			Handler:    _CrashlyticsService_RecordError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crashlytics.proto",
}
