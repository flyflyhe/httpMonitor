// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.2
// source: rpc/monitor.proto

package rpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MonitorServerClient is the client API for MonitorServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorServerClient interface {
	Start(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (MonitorServer_StartClient, error)
	Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type monitorServerClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorServerClient(cc grpc.ClientConnInterface) MonitorServerClient {
	return &monitorServerClient{cc}
}

func (c *monitorServerClient) Start(ctx context.Context, in *MonitorRequest, opts ...grpc.CallOption) (MonitorServer_StartClient, error) {
	stream, err := c.cc.NewStream(ctx, &MonitorServer_ServiceDesc.Streams[0], "/rpc.MonitorServer/Start", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorServerStartClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitorServer_StartClient interface {
	Recv() (*MonitorResponse, error)
	grpc.ClientStream
}

type monitorServerStartClient struct {
	grpc.ClientStream
}

func (x *monitorServerStartClient) Recv() (*MonitorResponse, error) {
	m := new(MonitorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *monitorServerClient) Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rpc.MonitorServer/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServerServer is the server API for MonitorServer service.
// All implementations must embed UnimplementedMonitorServerServer
// for forward compatibility
type MonitorServerServer interface {
	Start(*MonitorRequest, MonitorServer_StartServer) error
	Stop(context.Context, *empty.Empty) (*empty.Empty, error)
	mustEmbedUnimplementedMonitorServerServer()
}

// UnimplementedMonitorServerServer must be embedded to have forward compatible implementations.
type UnimplementedMonitorServerServer struct {
}

func (UnimplementedMonitorServerServer) Start(*MonitorRequest, MonitorServer_StartServer) error {
	return status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedMonitorServerServer) Stop(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedMonitorServerServer) mustEmbedUnimplementedMonitorServerServer() {}

// UnsafeMonitorServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorServerServer will
// result in compilation errors.
type UnsafeMonitorServerServer interface {
	mustEmbedUnimplementedMonitorServerServer()
}

func RegisterMonitorServerServer(s grpc.ServiceRegistrar, srv MonitorServerServer) {
	s.RegisterService(&MonitorServer_ServiceDesc, srv)
}

func _MonitorServer_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorServerServer).Start(m, &monitorServerStartServer{stream})
}

type MonitorServer_StartServer interface {
	Send(*MonitorResponse) error
	grpc.ServerStream
}

type monitorServerStartServer struct {
	grpc.ServerStream
}

func (x *monitorServerStartServer) Send(m *MonitorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MonitorServer_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MonitorServer/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServerServer).Stop(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MonitorServer_ServiceDesc is the grpc.ServiceDesc for MonitorServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonitorServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.MonitorServer",
	HandlerType: (*MonitorServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stop",
			Handler:    _MonitorServer_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Start",
			Handler:       _MonitorServer_Start_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc/monitor.proto",
}
