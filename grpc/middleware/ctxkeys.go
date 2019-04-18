package middleware

type (
	// private type used to define context keys
	ctxKey int
)

const (
	// RequestMethodKey is the context key used to store the full method name
	// of the gRPC request.
	RequestMethodKey ctxKey = iota + 1

	// RequestPeerAddrKey is the context key used to store the peer address from
	// which the request originated.
	RequestPeerAddrKey

	// RequestUserAgentKey is the context key used to store the "user-agent"
	// metadata value.
	RequestUserAgentKey

	// RequestXRequestIDKey is the context key used to store the "x-request-id"
	// metadata value.
	RequestXRequestIDKey
)
