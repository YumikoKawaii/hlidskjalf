// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: pawn/service.proto

/*
Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package api

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_GreetService_Greet_0(ctx context.Context, marshaler runtime.Marshaler, client GreetServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GreetRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Greet(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_GreetService_Greet_0(ctx context.Context, marshaler runtime.Marshaler, server GreetServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GreetRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Greet(ctx, &protoReq)
	return msg, metadata, err

}

func request_PerformanceService_GetStudents_0(ctx context.Context, marshaler runtime.Marshaler, client PerformanceServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetStudentsRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetStudents(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PerformanceService_GetStudents_0(ctx context.Context, marshaler runtime.Marshaler, server PerformanceServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetStudentsRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetStudents(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterGreetServiceHandlerServer registers the http handlers for service GreetService to "mux".
// UnaryRPC     :call GreetServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features (such as grpc.SendHeader, etc) to stop working. Consider using RegisterGreetServiceHandlerFromEndpoint instead.
func RegisterGreetServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GreetServiceServer) error {

	mux.Handle("GET", pattern_GreetService_Greet_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GreetService_Greet_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GreetService_Greet_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterPerformanceServiceHandlerServer registers the http handlers for service PerformanceService to "mux".
// UnaryRPC     :call PerformanceServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features (such as grpc.SendHeader, etc) to stop working. Consider using RegisterPerformanceServiceHandlerFromEndpoint instead.
func RegisterPerformanceServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server PerformanceServiceServer) error {

	mux.Handle("GET", pattern_PerformanceService_GetStudents_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PerformanceService_GetStudents_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PerformanceService_GetStudents_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterGreetServiceHandlerFromEndpoint is same as RegisterGreetServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGreetServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterGreetServiceHandler(ctx, mux, conn)
}

// RegisterGreetServiceHandler registers the http handlers for service GreetService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGreetServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGreetServiceHandlerClient(ctx, mux, NewGreetServiceClient(conn))
}

// RegisterGreetServiceHandlerClient registers the http handlers for service GreetService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GreetServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GreetServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GreetServiceClient" to call the correct interceptors.
func RegisterGreetServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GreetServiceClient) error {

	mux.Handle("GET", pattern_GreetService_Greet_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GreetService_Greet_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GreetService_Greet_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_GreetService_Greet_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "greet"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_GreetService_Greet_0 = runtime.ForwardResponseMessage
)

// RegisterPerformanceServiceHandlerFromEndpoint is same as RegisterPerformanceServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPerformanceServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterPerformanceServiceHandler(ctx, mux, conn)
}

// RegisterPerformanceServiceHandler registers the http handlers for service PerformanceService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPerformanceServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPerformanceServiceHandlerClient(ctx, mux, NewPerformanceServiceClient(conn))
}

// RegisterPerformanceServiceHandlerClient registers the http handlers for service PerformanceService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PerformanceServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PerformanceServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PerformanceServiceClient" to call the correct interceptors.
func RegisterPerformanceServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PerformanceServiceClient) error {

	mux.Handle("GET", pattern_PerformanceService_GetStudents_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PerformanceService_GetStudents_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PerformanceService_GetStudents_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PerformanceService_GetStudents_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "students"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_PerformanceService_GetStudents_0 = runtime.ForwardResponseMessage
)
