// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: protos/frontend/ohlc/v1/service.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"github.com/xefino/quantum-api-go/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_OhlcService_Aggregates_0(ctx context.Context, marshaler runtime.Marshaler, client OhlcServiceClient, req *http.Request, pathParams map[string]string) (OhlcService_AggregatesClient, runtime.ServerMetadata, error) {
	var protoReq GetAggregatesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["symbol"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "symbol")
	}

	protoReq.Symbol, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "symbol", err)
	}

	val, ok = pathParams["multiplier"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "multiplier")
	}

	protoReq.Multiplier, err = runtime.Uint32(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "multiplier", err)
	}

	val, ok = pathParams["frequency"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "frequency")
	}

	e, err = runtime.Enum(val, data.Frequency_value)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "frequency", err)
	}

	protoReq.Frequency = data.Frequency(e)

	stream, err := client.Aggregates(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterOhlcServiceHandlerServer registers the http handlers for service OhlcService to "mux".
// UnaryRPC     :call OhlcServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterOhlcServiceHandlerFromEndpoint instead.
func RegisterOhlcServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server OhlcServiceServer) error {

	mux.Handle("POST", pattern_OhlcService_Aggregates_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterOhlcServiceHandlerFromEndpoint is same as RegisterOhlcServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterOhlcServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterOhlcServiceHandler(ctx, mux, conn)
}

// RegisterOhlcServiceHandler registers the http handlers for service OhlcService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterOhlcServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterOhlcServiceHandlerClient(ctx, mux, NewOhlcServiceClient(conn))
}

// RegisterOhlcServiceHandlerClient registers the http handlers for service OhlcService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "OhlcServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "OhlcServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "OhlcServiceClient" to call the correct interceptors.
func RegisterOhlcServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client OhlcServiceClient) error {

	mux.Handle("POST", pattern_OhlcService_Aggregates_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/protos.frontend.ohlc.v1.OhlcService/Aggregates", runtime.WithHTTPPathPattern("/ohlc/aggregate/{symbol}/range/{multiplier}/{frequency}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_OhlcService_Aggregates_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_OhlcService_Aggregates_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_OhlcService_Aggregates_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"ohlc", "aggregate", "symbol", "range", "multiplier", "frequency"}, ""))
)

var (
	forward_OhlcService_Aggregates_0 = runtime.ForwardResponseStream
)
