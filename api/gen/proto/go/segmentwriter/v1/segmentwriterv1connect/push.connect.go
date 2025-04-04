// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: segmentwriter/v1/push.proto

package segmentwriterv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/segmentwriter/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SegmentWriterServiceName is the fully-qualified name of the SegmentWriterService service.
	SegmentWriterServiceName = "segmentwriter.v1.SegmentWriterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SegmentWriterServicePushProcedure is the fully-qualified name of the SegmentWriterService's Push
	// RPC.
	SegmentWriterServicePushProcedure = "/segmentwriter.v1.SegmentWriterService/Push"
)

// SegmentWriterServiceClient is a client for the segmentwriter.v1.SegmentWriterService service.
type SegmentWriterServiceClient interface {
	Push(context.Context, *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error)
}

// NewSegmentWriterServiceClient constructs a client for the segmentwriter.v1.SegmentWriterService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSegmentWriterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SegmentWriterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	segmentWriterServiceMethods := v1.File_segmentwriter_v1_push_proto.Services().ByName("SegmentWriterService").Methods()
	return &segmentWriterServiceClient{
		push: connect.NewClient[v1.PushRequest, v1.PushResponse](
			httpClient,
			baseURL+SegmentWriterServicePushProcedure,
			connect.WithSchema(segmentWriterServiceMethods.ByName("Push")),
			connect.WithClientOptions(opts...),
		),
	}
}

// segmentWriterServiceClient implements SegmentWriterServiceClient.
type segmentWriterServiceClient struct {
	push *connect.Client[v1.PushRequest, v1.PushResponse]
}

// Push calls segmentwriter.v1.SegmentWriterService.Push.
func (c *segmentWriterServiceClient) Push(ctx context.Context, req *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// SegmentWriterServiceHandler is an implementation of the segmentwriter.v1.SegmentWriterService
// service.
type SegmentWriterServiceHandler interface {
	Push(context.Context, *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error)
}

// NewSegmentWriterServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSegmentWriterServiceHandler(svc SegmentWriterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	segmentWriterServiceMethods := v1.File_segmentwriter_v1_push_proto.Services().ByName("SegmentWriterService").Methods()
	segmentWriterServicePushHandler := connect.NewUnaryHandler(
		SegmentWriterServicePushProcedure,
		svc.Push,
		connect.WithSchema(segmentWriterServiceMethods.ByName("Push")),
		connect.WithHandlerOptions(opts...),
	)
	return "/segmentwriter.v1.SegmentWriterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SegmentWriterServicePushProcedure:
			segmentWriterServicePushHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSegmentWriterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSegmentWriterServiceHandler struct{}

func (UnimplementedSegmentWriterServiceHandler) Push(context.Context, *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("segmentwriter.v1.SegmentWriterService.Push is not implemented"))
}
