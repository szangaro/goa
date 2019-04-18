package middleware

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// PopulateUnaryRequestContext returns a middleware which populates the context
// with a number of standard gRPC unary request values. Those values may be
// extracted using the corresponding ContextKey type in this package.
func PopulateUnaryRequestContext() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = populateContext(ctx, info.FullMethod)
		return handler(ctx, req)
	})
}

// PopulateStreamRequestContext returns a middleware which populates the context
// with a number of standard gRPC stream request values. Those values may be
// extracted using the corresponding ContextKey type in this package.
func PopulateStreamRequestContext() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := populateContext(ss.Context(), info.FullMethod)
		wss := NewWrappedServerStream(ctx, ss)
		return handler(srv, wss)
	})
}

func populateContext(ctx context.Context, method string) context.Context {
	ctx = context.WithValue(ctx, RequestMethodKey, method)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = context.WithValue(ctx, RequestUserAgentKey, MetadataValue(md, "user-agent"))
		ctx = context.WithValue(ctx, RequestXRequestIDKey, MetadataValue(md, RequestIDMetadataKey))
	}
	if p, ok := peer.FromContext(ctx); ok {
		ip, _, _ := net.SplitHostPort(p.Addr.String())
		ctx = context.WithValue(ctx, RequestPeerAddrKey, ip)
	}
	return ctx
}
