// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/proto/goload.proto

package goload

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccountService_CreateAccount_FullMethodName = "/goload.AccountService/CreateAccount"
	AccountService_CreateSession_FullMethodName = "/goload.AccountService/CreateSession"
)

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility.
type AccountServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountServiceServer struct{}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}
func (UnimplementedAccountServiceServer) testEmbeddedByValue()                        {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccountServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goload.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _AccountService_CreateSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/goload.proto",
}

const (
	DownloadService_CreateDownloadTask_FullMethodName  = "/goload.DownloadService/CreateDownloadTask"
	DownloadService_GetDownloadTaskList_FullMethodName = "/goload.DownloadService/GetDownloadTaskList"
	DownloadService_GetDownloadTask_FullMethodName     = "/goload.DownloadService/GetDownloadTask"
	DownloadService_UpdateDownloadTask_FullMethodName  = "/goload.DownloadService/UpdateDownloadTask"
	DownloadService_DeleteDownloadTask_FullMethodName  = "/goload.DownloadService/DeleteDownloadTask"
	DownloadService_GetDownloadFile_FullMethodName     = "/goload.DownloadService/GetDownloadFile"
)

// DownloadServiceClient is the client API for DownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloadServiceClient interface {
	CreateDownloadTask(ctx context.Context, in *CreateDownloadTaskRequest, opts ...grpc.CallOption) (*CreateDownloadTaskResponse, error)
	GetDownloadTaskList(ctx context.Context, in *GetDownloadTaskListRequest, opts ...grpc.CallOption) (*GetDownloadTaskListResponse, error)
	GetDownloadTask(ctx context.Context, in *GetDownloadTaskRequest, opts ...grpc.CallOption) (*GetDownloadTaskResponse, error)
	UpdateDownloadTask(ctx context.Context, in *UpdateDownloadTaskRequest, opts ...grpc.CallOption) (*UpdateDownloadTaskResponse, error)
	DeleteDownloadTask(ctx context.Context, in *DeleteDownloadTaskRequest, opts ...grpc.CallOption) (*DeleteDownloadTaskResponse, error)
	GetDownloadFile(ctx context.Context, in *GetDownloadFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDownloadFileResponse], error)
}

type downloadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadServiceClient(cc grpc.ClientConnInterface) DownloadServiceClient {
	return &downloadServiceClient{cc}
}

func (c *downloadServiceClient) CreateDownloadTask(ctx context.Context, in *CreateDownloadTaskRequest, opts ...grpc.CallOption) (*CreateDownloadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDownloadTaskResponse)
	err := c.cc.Invoke(ctx, DownloadService_CreateDownloadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadServiceClient) GetDownloadTaskList(ctx context.Context, in *GetDownloadTaskListRequest, opts ...grpc.CallOption) (*GetDownloadTaskListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDownloadTaskListResponse)
	err := c.cc.Invoke(ctx, DownloadService_GetDownloadTaskList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadServiceClient) GetDownloadTask(ctx context.Context, in *GetDownloadTaskRequest, opts ...grpc.CallOption) (*GetDownloadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDownloadTaskResponse)
	err := c.cc.Invoke(ctx, DownloadService_GetDownloadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadServiceClient) UpdateDownloadTask(ctx context.Context, in *UpdateDownloadTaskRequest, opts ...grpc.CallOption) (*UpdateDownloadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDownloadTaskResponse)
	err := c.cc.Invoke(ctx, DownloadService_UpdateDownloadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadServiceClient) DeleteDownloadTask(ctx context.Context, in *DeleteDownloadTaskRequest, opts ...grpc.CallOption) (*DeleteDownloadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDownloadTaskResponse)
	err := c.cc.Invoke(ctx, DownloadService_DeleteDownloadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadServiceClient) GetDownloadFile(ctx context.Context, in *GetDownloadFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDownloadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DownloadService_ServiceDesc.Streams[0], DownloadService_GetDownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetDownloadFileRequest, GetDownloadFileResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DownloadService_GetDownloadFileClient = grpc.ServerStreamingClient[GetDownloadFileResponse]

// DownloadServiceServer is the server API for DownloadService service.
// All implementations must embed UnimplementedDownloadServiceServer
// for forward compatibility.
type DownloadServiceServer interface {
	CreateDownloadTask(context.Context, *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error)
	GetDownloadTaskList(context.Context, *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error)
	GetDownloadTask(context.Context, *GetDownloadTaskRequest) (*GetDownloadTaskResponse, error)
	UpdateDownloadTask(context.Context, *UpdateDownloadTaskRequest) (*UpdateDownloadTaskResponse, error)
	DeleteDownloadTask(context.Context, *DeleteDownloadTaskRequest) (*DeleteDownloadTaskResponse, error)
	GetDownloadFile(*GetDownloadFileRequest, grpc.ServerStreamingServer[GetDownloadFileResponse]) error
	mustEmbedUnimplementedDownloadServiceServer()
}

// UnimplementedDownloadServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDownloadServiceServer struct{}

func (UnimplementedDownloadServiceServer) CreateDownloadTask(context.Context, *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDownloadTask not implemented")
}
func (UnimplementedDownloadServiceServer) GetDownloadTaskList(context.Context, *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadTaskList not implemented")
}
func (UnimplementedDownloadServiceServer) GetDownloadTask(context.Context, *GetDownloadTaskRequest) (*GetDownloadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadTask not implemented")
}
func (UnimplementedDownloadServiceServer) UpdateDownloadTask(context.Context, *UpdateDownloadTaskRequest) (*UpdateDownloadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDownloadTask not implemented")
}
func (UnimplementedDownloadServiceServer) DeleteDownloadTask(context.Context, *DeleteDownloadTaskRequest) (*DeleteDownloadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDownloadTask not implemented")
}
func (UnimplementedDownloadServiceServer) GetDownloadFile(*GetDownloadFileRequest, grpc.ServerStreamingServer[GetDownloadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetDownloadFile not implemented")
}
func (UnimplementedDownloadServiceServer) mustEmbedUnimplementedDownloadServiceServer() {}
func (UnimplementedDownloadServiceServer) testEmbeddedByValue()                         {}

// UnsafeDownloadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloadServiceServer will
// result in compilation errors.
type UnsafeDownloadServiceServer interface {
	mustEmbedUnimplementedDownloadServiceServer()
}

func RegisterDownloadServiceServer(s grpc.ServiceRegistrar, srv DownloadServiceServer) {
	// If the following call pancis, it indicates UnimplementedDownloadServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DownloadService_ServiceDesc, srv)
}

func _DownloadService_CreateDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).CreateDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DownloadService_CreateDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).CreateDownloadTask(ctx, req.(*CreateDownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownloadService_GetDownloadTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadTaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).GetDownloadTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DownloadService_GetDownloadTaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).GetDownloadTaskList(ctx, req.(*GetDownloadTaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownloadService_GetDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).GetDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DownloadService_GetDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).GetDownloadTask(ctx, req.(*GetDownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownloadService_UpdateDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).UpdateDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DownloadService_UpdateDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).UpdateDownloadTask(ctx, req.(*UpdateDownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownloadService_DeleteDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).DeleteDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DownloadService_DeleteDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).DeleteDownloadTask(ctx, req.(*DeleteDownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownloadService_GetDownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownloadServiceServer).GetDownloadFile(m, &grpc.GenericServerStream[GetDownloadFileRequest, GetDownloadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DownloadService_GetDownloadFileServer = grpc.ServerStreamingServer[GetDownloadFileResponse]

// DownloadService_ServiceDesc is the grpc.ServiceDesc for DownloadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DownloadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goload.DownloadService",
	HandlerType: (*DownloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDownloadTask",
			Handler:    _DownloadService_CreateDownloadTask_Handler,
		},
		{
			MethodName: "GetDownloadTaskList",
			Handler:    _DownloadService_GetDownloadTaskList_Handler,
		},
		{
			MethodName: "GetDownloadTask",
			Handler:    _DownloadService_GetDownloadTask_Handler,
		},
		{
			MethodName: "UpdateDownloadTask",
			Handler:    _DownloadService_UpdateDownloadTask_Handler,
		},
		{
			MethodName: "DeleteDownloadTask",
			Handler:    _DownloadService_DeleteDownloadTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDownloadFile",
			Handler:       _DownloadService_GetDownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/goload.proto",
}
