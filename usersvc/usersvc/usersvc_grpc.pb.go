// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.17.3
// source: usersvc.proto

package usersvc

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
	Usersvc_Ping_FullMethodName = "/usersvc.Usersvc/Ping"
)

// UsersvcClient is the client API for Usersvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersvcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type usersvcClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersvcClient(cc grpc.ClientConnInterface) UsersvcClient {
	return &usersvcClient{cc}
}

func (c *usersvcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Usersvc_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersvcServer is the server API for Usersvc service.
// All implementations must embed UnimplementedUsersvcServer
// for forward compatibility.
type UsersvcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedUsersvcServer()
}

// UnimplementedUsersvcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersvcServer struct{}

func (UnimplementedUsersvcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUsersvcServer) mustEmbedUnimplementedUsersvcServer() {}
func (UnimplementedUsersvcServer) testEmbeddedByValue()                 {}

// UnsafeUsersvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersvcServer will
// result in compilation errors.
type UnsafeUsersvcServer interface {
	mustEmbedUnimplementedUsersvcServer()
}

func RegisterUsersvcServer(s grpc.ServiceRegistrar, srv UsersvcServer) {
	// If the following call pancis, it indicates UnimplementedUsersvcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Usersvc_ServiceDesc, srv)
}

func _Usersvc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersvcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usersvc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersvcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Usersvc_ServiceDesc is the grpc.ServiceDesc for Usersvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usersvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usersvc.Usersvc",
	HandlerType: (*UsersvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Usersvc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usersvc.proto",
}