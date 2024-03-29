// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: protos/frontend/iam/v1/service.proto

package v1

import (
	context "context"
	data "github.com/xefino/quantum-api-go/data"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IamService_RegisterClient_FullMethodName = "/protos.frontend.iam.v1.IamService/RegisterClient"
	IamService_GetClients_FullMethodName     = "/protos.frontend.iam.v1.IamService/GetClients"
	IamService_GetClient_FullMethodName      = "/protos.frontend.iam.v1.IamService/GetClient"
	IamService_DeleteClient_FullMethodName   = "/protos.frontend.iam.v1.IamService/DeleteClient"
	IamService_CreateUser_FullMethodName     = "/protos.frontend.iam.v1.IamService/CreateUser"
	IamService_GetUser_FullMethodName        = "/protos.frontend.iam.v1.IamService/GetUser"
	IamService_UpdateEmail_FullMethodName    = "/protos.frontend.iam.v1.IamService/UpdateEmail"
	IamService_UpdatePassword_FullMethodName = "/protos.frontend.iam.v1.IamService/UpdatePassword"
)

// IamServiceClient is the client API for IamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamServiceClient interface {
	// Creates a new client associated with the current user, returning the client's information along with its credentials. This request is
	// allowed only to only create clients which are either for a non-Xefino domain or else for a non-browser client. This restriction is
	// intended as a security measure to prevent user clients from impersonating our official web client. Client credentials will be
	// available from this endpoint when the client is created, and only at that time.
	RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error)
	// Retrieves information on all the clients that have been registered for the current user. Note that secrets will not be returned from
	// this endpoint.
	GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (IamService_GetClientsClient, error)
	// Retrieves the client associated with the current user and the client ID. Note that the secret associated with this client will not
	// be returned from this endpoint. If there is no client associated with the user ID and client ID provided, then this endpoint will
	// return a not-found response
	GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*data.Client, error)
	// Removes the client associated with the current user and client ID, returning the client's information. Note that the secret
	// associated with this client will not be returned from this endpoint. If there is not client associated with the user ID and client ID
	// provided, then this endpoint will return a not-found response
	DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*data.Client, error)
	// Creates a new user with the desired username and password. This request will fail if a user is created with the same email address as
	// another user, or if the password is empty. Only our official client is allowed to access this endpoint.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*data.User, error)
	// Retrieves the account information for the current user. This function will not return any sensitive information associated with the user.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*data.User, error)
	// Updates the email address (username) associated with the current user. This function will return the updated user information, but this will
	// not include any sensitive information.
	UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*data.User, error)
	// Updates the password associated with the current user. This function will not return any sensitive information associated with the user.
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*data.User, error)
}

type iamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIamServiceClient(cc grpc.ClientConnInterface) IamServiceClient {
	return &iamServiceClient{cc}
}

func (c *iamServiceClient) RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error) {
	out := new(RegisterClientResponse)
	err := c.cc.Invoke(ctx, IamService_RegisterClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (IamService_GetClientsClient, error) {
	stream, err := c.cc.NewStream(ctx, &IamService_ServiceDesc.Streams[0], IamService_GetClients_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &iamServiceGetClientsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IamService_GetClientsClient interface {
	Recv() (*data.Client, error)
	grpc.ClientStream
}

type iamServiceGetClientsClient struct {
	grpc.ClientStream
}

func (x *iamServiceGetClientsClient) Recv() (*data.Client, error) {
	m := new(data.Client)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iamServiceClient) GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*data.Client, error) {
	out := new(data.Client)
	err := c.cc.Invoke(ctx, IamService_GetClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*data.Client, error) {
	out := new(data.Client)
	err := c.cc.Invoke(ctx, IamService_DeleteClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, IamService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, IamService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, IamService_UpdateEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*data.User, error) {
	out := new(data.User)
	err := c.cc.Invoke(ctx, IamService_UpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamServiceServer is the server API for IamService service.
// All implementations must embed UnimplementedIamServiceServer
// for forward compatibility
type IamServiceServer interface {
	// Creates a new client associated with the current user, returning the client's information along with its credentials. This request is
	// allowed only to only create clients which are either for a non-Xefino domain or else for a non-browser client. This restriction is
	// intended as a security measure to prevent user clients from impersonating our official web client. Client credentials will be
	// available from this endpoint when the client is created, and only at that time.
	RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error)
	// Retrieves information on all the clients that have been registered for the current user. Note that secrets will not be returned from
	// this endpoint.
	GetClients(*GetClientsRequest, IamService_GetClientsServer) error
	// Retrieves the client associated with the current user and the client ID. Note that the secret associated with this client will not
	// be returned from this endpoint. If there is no client associated with the user ID and client ID provided, then this endpoint will
	// return a not-found response
	GetClient(context.Context, *GetClientRequest) (*data.Client, error)
	// Removes the client associated with the current user and client ID, returning the client's information. Note that the secret
	// associated with this client will not be returned from this endpoint. If there is not client associated with the user ID and client ID
	// provided, then this endpoint will return a not-found response
	DeleteClient(context.Context, *DeleteClientRequest) (*data.Client, error)
	// Creates a new user with the desired username and password. This request will fail if a user is created with the same email address as
	// another user, or if the password is empty. Only our official client is allowed to access this endpoint.
	CreateUser(context.Context, *CreateUserRequest) (*data.User, error)
	// Retrieves the account information for the current user. This function will not return any sensitive information associated with the user.
	GetUser(context.Context, *GetUserRequest) (*data.User, error)
	// Updates the email address (username) associated with the current user. This function will return the updated user information, but this will
	// not include any sensitive information.
	UpdateEmail(context.Context, *UpdateEmailRequest) (*data.User, error)
	// Updates the password associated with the current user. This function will not return any sensitive information associated with the user.
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*data.User, error)
	mustEmbedUnimplementedIamServiceServer()
}

// UnimplementedIamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIamServiceServer struct {
}

func (UnimplementedIamServiceServer) RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClient not implemented")
}
func (UnimplementedIamServiceServer) GetClients(*GetClientsRequest, IamService_GetClientsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetClients not implemented")
}
func (UnimplementedIamServiceServer) GetClient(context.Context, *GetClientRequest) (*data.Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedIamServiceServer) DeleteClient(context.Context, *DeleteClientRequest) (*data.Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedIamServiceServer) CreateUser(context.Context, *CreateUserRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedIamServiceServer) GetUser(context.Context, *GetUserRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIamServiceServer) UpdateEmail(context.Context, *UpdateEmailRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmail not implemented")
}
func (UnimplementedIamServiceServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*data.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedIamServiceServer) mustEmbedUnimplementedIamServiceServer() {}

// UnsafeIamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamServiceServer will
// result in compilation errors.
type UnsafeIamServiceServer interface {
	mustEmbedUnimplementedIamServiceServer()
}

func RegisterIamServiceServer(s grpc.ServiceRegistrar, srv IamServiceServer) {
	s.RegisterService(&IamService_ServiceDesc, srv)
}

func _IamService_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamService_RegisterClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).RegisterClient(ctx, req.(*RegisterClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_GetClients_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetClientsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IamServiceServer).GetClients(m, &iamServiceGetClientsServer{stream})
}

type IamService_GetClientsServer interface {
	Send(*data.Client) error
	grpc.ServerStream
}

type iamServiceGetClientsServer struct {
	grpc.ServerStream
}

func (x *iamServiceGetClientsServer) Send(m *data.Client) error {
	return x.ServerStream.SendMsg(m)
}

func _IamService_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamService_GetClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).GetClient(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamService_DeleteClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).DeleteClient(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_UpdateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).UpdateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamService_UpdateEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).UpdateEmail(ctx, req.(*UpdateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamService_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IamService_ServiceDesc is the grpc.ServiceDesc for IamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.frontend.iam.v1.IamService",
	HandlerType: (*IamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterClient",
			Handler:    _IamService_RegisterClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _IamService_GetClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _IamService_DeleteClient_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _IamService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _IamService_GetUser_Handler,
		},
		{
			MethodName: "UpdateEmail",
			Handler:    _IamService_UpdateEmail_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _IamService_UpdatePassword_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetClients",
			Handler:       _IamService_GetClients_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/frontend/iam/v1/service.proto",
}
