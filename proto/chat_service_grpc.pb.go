// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatService_StartThread_FullMethodName             = "/chatbot.chat.v4.ChatService/StartThread"
	ChatService_PostMessage_FullMethodName             = "/chatbot.chat.v4.ChatService/PostMessage"
	ChatService_GetThread_FullMethodName               = "/chatbot.chat.v4.ChatService/GetThread"
	ChatService_ListThreadIDs_FullMethodName           = "/chatbot.chat.v4.ChatService/ListThreadIDs"
	ChatService_DeleteThread_FullMethodName            = "/chatbot.chat.v4.ChatService/DeleteThread"
	ChatService_DeleteMessageFromThread_FullMethodName = "/chatbot.chat.v4.ChatService/DeleteMessageFromThread"
	ChatService_BatchChat_FullMethodName               = "/chatbot.chat.v4.ChatService/BatchChat"
	ChatService_Completion_FullMethodName              = "/chatbot.chat.v4.ChatService/Completion"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	StartThread(ctx context.Context, in *ThreadPrompt, opts ...grpc.CallOption) (*Thread, error)
	PostMessage(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Message, error)
	GetThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*Thread, error)
	ListThreadIDs(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*ThreadIDs, error)
	DeleteThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteMessageFromThread(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BatchChat(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error)
	Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) StartThread(ctx context.Context, in *ThreadPrompt, opts ...grpc.CallOption) (*Thread, error) {
	out := new(Thread)
	err := c.cc.Invoke(ctx, ChatService_StartThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) PostMessage(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, ChatService_PostMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*Thread, error) {
	out := new(Thread)
	err := c.cc.Invoke(ctx, ChatService_GetThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ListThreadIDs(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*ThreadIDs, error) {
	out := new(ThreadIDs)
	err := c.cc.Invoke(ctx, ChatService_ListThreadIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteThread(ctx context.Context, in *ThreadID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_DeleteThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteMessageFromThread(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_DeleteMessageFromThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) BatchChat(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.cc.Invoke(ctx, ChatService_BatchChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error) {
	out := new(CompletionResponse)
	err := c.cc.Invoke(ctx, ChatService_Completion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	StartThread(context.Context, *ThreadPrompt) (*Thread, error)
	PostMessage(context.Context, *Prompt) (*Message, error)
	GetThread(context.Context, *ThreadID) (*Thread, error)
	ListThreadIDs(context.Context, *Collection) (*ThreadIDs, error)
	DeleteThread(context.Context, *ThreadID) (*emptypb.Empty, error)
	DeleteMessageFromThread(context.Context, *MessageID) (*emptypb.Empty, error)
	BatchChat(context.Context, *BatchRequest) (*BatchResponse, error)
	Completion(context.Context, *CompletionRequest) (*CompletionResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) StartThread(context.Context, *ThreadPrompt) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartThread not implemented")
}
func (UnimplementedChatServiceServer) PostMessage(context.Context, *Prompt) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMessage not implemented")
}
func (UnimplementedChatServiceServer) GetThread(context.Context, *ThreadID) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThread not implemented")
}
func (UnimplementedChatServiceServer) ListThreadIDs(context.Context, *Collection) (*ThreadIDs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListThreadIDs not implemented")
}
func (UnimplementedChatServiceServer) DeleteThread(context.Context, *ThreadID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThread not implemented")
}
func (UnimplementedChatServiceServer) DeleteMessageFromThread(context.Context, *MessageID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessageFromThread not implemented")
}
func (UnimplementedChatServiceServer) BatchChat(context.Context, *BatchRequest) (*BatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchChat not implemented")
}
func (UnimplementedChatServiceServer) Completion(context.Context, *CompletionRequest) (*CompletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Completion not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_StartThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadPrompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).StartThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_StartThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).StartThread(ctx, req.(*ThreadPrompt))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_PostMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prompt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).PostMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_PostMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).PostMessage(ctx, req.(*Prompt))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetThread(ctx, req.(*ThreadID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ListThreadIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ListThreadIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_ListThreadIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ListThreadIDs(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteThread(ctx, req.(*ThreadID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteMessageFromThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteMessageFromThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteMessageFromThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteMessageFromThread(ctx, req.(*MessageID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_BatchChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).BatchChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_BatchChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).BatchChat(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Completion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Completion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Completion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Completion(ctx, req.(*CompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatbot.chat.v4.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartThread",
			Handler:    _ChatService_StartThread_Handler,
		},
		{
			MethodName: "PostMessage",
			Handler:    _ChatService_PostMessage_Handler,
		},
		{
			MethodName: "GetThread",
			Handler:    _ChatService_GetThread_Handler,
		},
		{
			MethodName: "ListThreadIDs",
			Handler:    _ChatService_ListThreadIDs_Handler,
		},
		{
			MethodName: "DeleteThread",
			Handler:    _ChatService_DeleteThread_Handler,
		},
		{
			MethodName: "DeleteMessageFromThread",
			Handler:    _ChatService_DeleteMessageFromThread_Handler,
		},
		{
			MethodName: "BatchChat",
			Handler:    _ChatService_BatchChat_Handler,
		},
		{
			MethodName: "Completion",
			Handler:    _ChatService_Completion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat_service.proto",
}
