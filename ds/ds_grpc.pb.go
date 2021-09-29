// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ds

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DSClient is the client API for DS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DSClient interface {
	Lookup(ctx context.Context, in *DSQuery, opts ...grpc.CallOption) (*DSResponse, error)
	Register(ctx context.Context, opts ...grpc.CallOption) (DS_RegisterClient, error)
}

type dSClient struct {
	cc grpc.ClientConnInterface
}

func NewDSClient(cc grpc.ClientConnInterface) DSClient {
	return &dSClient{cc}
}

func (c *dSClient) Lookup(ctx context.Context, in *DSQuery, opts ...grpc.CallOption) (*DSResponse, error) {
	out := new(DSResponse)
	err := c.cc.Invoke(ctx, "/ds.DS/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSClient) Register(ctx context.Context, opts ...grpc.CallOption) (DS_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &DS_ServiceDesc.Streams[0], "/ds.DS/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &dSRegisterClient{stream}
	return x, nil
}

type DS_RegisterClient interface {
	Send(*Registration) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type dSRegisterClient struct {
	grpc.ClientStream
}

func (x *dSRegisterClient) Send(m *Registration) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dSRegisterClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DSServer is the server API for DS service.
// All implementations must embed UnimplementedDSServer
// for forward compatibility
type DSServer interface {
	Lookup(context.Context, *DSQuery) (*DSResponse, error)
	Register(DS_RegisterServer) error
	mustEmbedUnimplementedDSServer()
}

// UnimplementedDSServer must be embedded to have forward compatible implementations.
type UnimplementedDSServer struct {
}

func (UnimplementedDSServer) Lookup(context.Context, *DSQuery) (*DSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (UnimplementedDSServer) Register(DS_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDSServer) mustEmbedUnimplementedDSServer() {}

// UnsafeDSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DSServer will
// result in compilation errors.
type UnsafeDSServer interface {
	mustEmbedUnimplementedDSServer()
}

func RegisterDSServer(s grpc.ServiceRegistrar, srv DSServer) {
	s.RegisterService(&DS_ServiceDesc, srv)
}

func _DS_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DSQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ds.DS/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSServer).Lookup(ctx, req.(*DSQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _DS_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DSServer).Register(&dSRegisterServer{stream})
}

type DS_RegisterServer interface {
	SendAndClose(*Empty) error
	Recv() (*Registration, error)
	grpc.ServerStream
}

type dSRegisterServer struct {
	grpc.ServerStream
}

func (x *dSRegisterServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dSRegisterServer) Recv() (*Registration, error) {
	m := new(Registration)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DS_ServiceDesc is the grpc.ServiceDesc for DS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ds.DS",
	HandlerType: (*DSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _DS_Lookup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _DS_Register_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "ds/ds.proto",
}
