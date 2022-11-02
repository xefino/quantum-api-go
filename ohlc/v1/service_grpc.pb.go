// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/frontend/ohlc/v1/service.proto

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

// OhlcServiceClient is the client API for OhlcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OhlcServiceClient interface {
	// Retrieves aggregated trading data over the time period, for the symbol and time frequency provided. This endpoint will not return
	// a not-found response if the symbol wasn't associated with any known value. instead, an empty list will be streamed to the client.
	// A bad-request response will be returned if the multiplier is negative, the frequency is invalid or the start time (from) comes after
	// the end time (to). If the time range contains part of a bar then request can contain a flag to either pull in that extra bar or
	// ignore it. This endpoint will adjust for splits by default. If this functionality is not necessary thn the associated flag should be set.
	Aggregates(ctx context.Context, in *GetAggregatesRequest, opts ...grpc.CallOption) (OhlcService_AggregatesClient, error)
}

type ohlcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOhlcServiceClient(cc grpc.ClientConnInterface) OhlcServiceClient {
	return &ohlcServiceClient{cc}
}

func (c *ohlcServiceClient) Aggregates(ctx context.Context, in *GetAggregatesRequest, opts ...grpc.CallOption) (OhlcService_AggregatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &OhlcService_ServiceDesc.Streams[0], "/protos.frontend.ohlc.v1.OhlcService/Aggregates", opts...)
	if err != nil {
		return nil, err
	}
	x := &ohlcServiceAggregatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OhlcService_AggregatesClient interface {
	Recv() (*data.Bar, error)
	grpc.ClientStream
}

type ohlcServiceAggregatesClient struct {
	grpc.ClientStream
}

func (x *ohlcServiceAggregatesClient) Recv() (*data.Bar, error) {
	m := new(data.Bar)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OhlcServiceServer is the server API for OhlcService service.
// All implementations must embed UnimplementedOhlcServiceServer
// for forward compatibility
type OhlcServiceServer interface {
	// Retrieves aggregated trading data over the time period, for the symbol and time frequency provided. This endpoint will not return
	// a not-found response if the symbol wasn't associated with any known value. instead, an empty list will be streamed to the client.
	// A bad-request response will be returned if the multiplier is negative, the frequency is invalid or the start time (from) comes after
	// the end time (to). If the time range contains part of a bar then request can contain a flag to either pull in that extra bar or
	// ignore it. This endpoint will adjust for splits by default. If this functionality is not necessary thn the associated flag should be set.
	Aggregates(*GetAggregatesRequest, OhlcService_AggregatesServer) error
	mustEmbedUnimplementedOhlcServiceServer()
}

// UnimplementedOhlcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOhlcServiceServer struct {
}

func (UnimplementedOhlcServiceServer) Aggregates(*GetAggregatesRequest, OhlcService_AggregatesServer) error {
	return status.Errorf(codes.Unimplemented, "method Aggregates not implemented")
}
func (UnimplementedOhlcServiceServer) mustEmbedUnimplementedOhlcServiceServer() {}

// UnsafeOhlcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OhlcServiceServer will
// result in compilation errors.
type UnsafeOhlcServiceServer interface {
	mustEmbedUnimplementedOhlcServiceServer()
}

func RegisterOhlcServiceServer(s grpc.ServiceRegistrar, srv OhlcServiceServer) {
	s.RegisterService(&OhlcService_ServiceDesc, srv)
}

func _OhlcService_Aggregates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAggregatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OhlcServiceServer).Aggregates(m, &ohlcServiceAggregatesServer{stream})
}

type OhlcService_AggregatesServer interface {
	Send(*data.Bar) error
	grpc.ServerStream
}

type ohlcServiceAggregatesServer struct {
	grpc.ServerStream
}

func (x *ohlcServiceAggregatesServer) Send(m *data.Bar) error {
	return x.ServerStream.SendMsg(m)
}

// OhlcService_ServiceDesc is the grpc.ServiceDesc for OhlcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OhlcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.frontend.ohlc.v1.OhlcService",
	HandlerType: (*OhlcServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Aggregates",
			Handler:       _OhlcService_Aggregates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/frontend/ohlc/v1/service.proto",
}
