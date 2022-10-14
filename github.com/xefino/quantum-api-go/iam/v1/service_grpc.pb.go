// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: protos/iam/v1/service.proto

package v1

import (
	context "context"
	data "github.com/xefino/quantum-api-go/data"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IAMClient is the client API for IAM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IAMClient interface {
	// Creates a new client associated with the current user, returning
	// the client's information along with its credentials
	RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*data.Client, error)
	// Retrieves all clients registered with the current user
	GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (IAM_GetClientsClient, error)
	// Retrieves the client associated with the current user and the client ID
	GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*data.Client, error)
	// Removes the client associated with the current user and client ID
	DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Creates a new user with the desired username and password
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*data.User, error)
	// Retrieves the account information for the current user
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*data.User, error)
	// Updates the email address (username) associated with the current user
	UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*data.User, error)
	// Updates the password associated with the current user
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*data.User, error)
}

type iAMClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMClient(cc grpc.ClientConnInterface) IAMClient {
	return &iAMClient{cc}
}

func (c *iAMClient) RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*data.Client, error) {
	out := new(data.Client)
	err := c.cc.Invoke(ctx, "/quantumapi.iam.v1.IAM/RegisterClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (IAM_GetClientsClient, error) {
	stream, err := c.cc.NewStream(ctx, &IAM_ServiceDesc.Streams[0], "/quantumapi.iam.v1.IAM/GetClients", opts...)
	if err != nil {
		return nil, err
	}
	x := &iAMGetClientsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IAM_GetClientsClient interface {
	Recv() (*data.Client, error)
	grpc.ClientStream
}

type iAMGetClientsClient struct {
	grpc.ClientStream
}

func (x *iAMGetClientsClient) Recv() (*data.Client, error) {
	m := new(data.Client)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iAMClient) GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*data.Client, error) {
	out := new(data.Client)
	err := c.cc.Invoke(ctx, "/quantumapi.iam.v1.IAM/GetClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/quantumapi.iam.v1.IAM/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, "/quantumapi.iam.v1.IAM/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, "/quantumapi.iam.v1.IAM/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, "/quantumapi.iam.v1.IAM/UpdateEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, "/quantumapi.iam.v1.IAM/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMServer is the server API for IAM service.
// All implementations must embed UnimplementedIAMServer
// for forward compatibility
type IAMServer interface {
	// Creates a new client associated with the current user, returning
	// the client's information along with its credentials
	RegisterClient(context.Context, *RegisterClientRequest) (*data.Client, error)
	// Retrieves all clients registered with the current user
	GetClients(*GetClientsRequest, IAM_GetClientsServer) error
	// Retrieves the client associated with the current user and the client ID
	GetClient(context.Context, *GetClientRequest) (*data.Client, error)
	// Removes the client associated with the current user and client ID
	DeleteClient(context.Context, *DeleteClientRequest) (*emptypb.Empty, error)
	// Creates a new user with the desired username and password
	CreateUser(context.Context, *CreateUserRequest) (*data.User, error)
	// Retrieves the account information for the current user
	GetUser(context.Context, *GetUserRequest) (*data.User, error)
	// Updates the email address (username) associated with the current user
	UpdateEmail(context.Context, *UpdateEmailRequest) (*data.User, error)
	// Updates the password associated with the current user
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*data.User, error)
	mustEmbedUnimplementedIAMServer()
}

// UnimplementedIAMServer must be embedded to have forward compatible implementations.
type UnimplementedIAMServer struct {
}

func (UnimplementedIAMServer) RegisterClient(context.Context, *RegisterClientRequest) (*data.Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClient not implemented")
}
func (UnimplementedIAMServer) GetClients(*GetClientsRequest, IAM_GetClientsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetClients not implemented")
}
func (UnimplementedIAMServer) GetClient(context.Context, *GetClientRequest) (*data.Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedIAMServer) DeleteClient(context.Context, *DeleteClientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedIAMServer) CreateUser(context.Context, *CreateUserRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedIAMServer) GetUser(context.Context, *GetUserRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIAMServer) UpdateEmail(context.Context, *UpdateEmailRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmail not implemented")
}
func (UnimplementedIAMServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedIAMServer) mustEmbedUnimplementedIAMServer() {}

// UnsafeIAMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IAMServer will
// result in compilation errors.
type UnsafeIAMServer interface {
	mustEmbedUnimplementedIAMServer()
}

func RegisterIAMServer(s grpc.ServiceRegistrar, srv IAMServer) {
	s.RegisterService(&IAM_ServiceDesc, srv)
}

func _IAM_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quantumapi.iam.v1.IAM/RegisterClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).RegisterClient(ctx, req.(*RegisterClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_GetClients_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetClientsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IAMServer).GetClients(m, &iAMGetClientsServer{stream})
}

type IAM_GetClientsServer interface {
	Send(*data.Client) error
	grpc.ServerStream
}

type iAMGetClientsServer struct {
	grpc.ServerStream
}

func (x *iAMGetClientsServer) Send(m *data.Client) error {
	return x.ServerStream.SendMsg(m)
}

func _IAM_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quantumapi.iam.v1.IAM/GetClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).GetClient(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quantumapi.iam.v1.IAM/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).DeleteClient(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quantumapi.iam.v1.IAM/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quantumapi.iam.v1.IAM/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_UpdateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).UpdateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quantumapi.iam.v1.IAM/UpdateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).UpdateEmail(ctx, req.(*UpdateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAM_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quantumapi.iam.v1.IAM/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IAM_ServiceDesc is the grpc.ServiceDesc for IAM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IAM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quantumapi.iam.v1.IAM",
	HandlerType: (*IAMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterClient",
			Handler:    _IAM_RegisterClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _IAM_GetClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _IAM_DeleteClient_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _IAM_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _IAM_GetUser_Handler,
		},
		{
			MethodName: "UpdateEmail",
			Handler:    _IAM_UpdateEmail_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _IAM_UpdatePassword_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetClients",
			Handler:       _IAM_GetClients_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/iam/v1/service.proto",
}
