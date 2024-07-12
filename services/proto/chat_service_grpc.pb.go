// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: chat_service.proto

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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Chat_PostMessage_FullMethodName             = "/chatbot.chat.v1.Chat/PostMessage"
	Chat_GetThread_FullMethodName               = "/chatbot.chat.v1.Chat/GetThread"
	Chat_ListThreadIDs_FullMethodName           = "/chatbot.chat.v1.Chat/ListThreadIDs"
	Chat_DeleteThread_FullMethodName            = "/chatbot.chat.v1.Chat/DeleteThread"
	Chat_DeleteMessageFromThread_FullMethodName = "/chatbot.chat.v1.Chat/DeleteMessageFromThread"
	Chat_Completion_FullMethodName              = "/chatbot.chat.v1.Chat/Completion"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	PostMessage(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Message, error)
	GetThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*Thread, error)
	ListThreadIDs(ctx context.Context, in *CollectionId, opts ...grpc.CallOption) (*ThreadIDs, error)
	DeleteThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteMessageFromThread(ctx context.Context, in *MessageIndex, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) PostMessage(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, Chat_PostMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*Thread, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Thread)
	err := c.cc.Invoke(ctx, Chat_GetThread_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) ListThreadIDs(ctx context.Context, in *CollectionId, opts ...grpc.CallOption) (*ThreadIDs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ThreadIDs)
	err := c.cc.Invoke(ctx, Chat_ListThreadIDs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) DeleteThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Chat_DeleteThread_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) DeleteMessageFromThread(ctx context.Context, in *MessageIndex, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Chat_DeleteMessageFromThread_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompletionResponse)
	err := c.cc.Invoke(ctx, Chat_Completion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	PostMessage(context.Context, *Prompt) (*Message, error)
	GetThread(context.Context, *ThreadID) (*Thread, error)
	ListThreadIDs(context.Context, *CollectionId) (*ThreadIDs, error)
	DeleteThread(context.Context, *ThreadID) (*emptypb.Empty, error)
	DeleteMessageFromThread(context.Context, *MessageIndex) (*emptypb.Empty, error)
	Completion(context.Context, *CompletionRequest) (*CompletionResponse, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) PostMessage(context.Context, *Prompt) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMessage not implemented")
}
func (UnimplementedChatServer) GetThread(context.Context, *ThreadID) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThread not implemented")
}
func (UnimplementedChatServer) ListThreadIDs(context.Context, *CollectionId) (*ThreadIDs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListThreadIDs not implemented")
}
func (UnimplementedChatServer) DeleteThread(context.Context, *ThreadID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThread not implemented")
}
func (UnimplementedChatServer) DeleteMessageFromThread(context.Context, *MessageIndex) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessageFromThread not implemented")
}
func (UnimplementedChatServer) Completion(context.Context, *CompletionRequest) (*CompletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Completion not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_PostMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).PostMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_PostMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).PostMessage(ctx, req.(*Prompt))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetThread(ctx, req.(*ThreadID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_ListThreadIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ListThreadIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_ListThreadIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ListThreadIDs(ctx, req.(*CollectionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_DeleteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).DeleteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_DeleteThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).DeleteThread(ctx, req.(*ThreadID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_DeleteMessageFromThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).DeleteMessageFromThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_DeleteMessageFromThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).DeleteMessageFromThread(ctx, req.(*MessageIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Completion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Completion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_Completion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Completion(ctx, req.(*CompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatbot.chat.v1.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostMessage",
			Handler:    _Chat_PostMessage_Handler,
		},
		{
			MethodName: "GetThread",
			Handler:    _Chat_GetThread_Handler,
		},
		{
			MethodName: "ListThreadIDs",
			Handler:    _Chat_ListThreadIDs_Handler,
		},
		{
			MethodName: "DeleteThread",
			Handler:    _Chat_DeleteThread_Handler,
		},
		{
			MethodName: "DeleteMessageFromThread",
			Handler:    _Chat_DeleteMessageFromThread_Handler,
		},
		{
			MethodName: "Completion",
			Handler:    _Chat_Completion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat_service.proto",
}
