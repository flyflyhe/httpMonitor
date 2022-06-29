// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.2
// source: rpc/url.proto

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

// UrlServiceClient is the client API for UrlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlServiceClient interface {
	SetUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error)
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UrlListResponse, error)
	GetAllDomainAndInterval(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UrlIntervalResponse, error)
	DeleteUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error)
	SetProxy(ctx context.Context, in *ProxyRequest, opts ...grpc.CallOption) (*ProxyResponse, error)
	GetAllProxy(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProxyListResponse, error)
	DeleteProxy(ctx context.Context, in *ProxyRequest, opts ...grpc.CallOption) (*ProxyResponse, error)
}

type urlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlServiceClient(cc grpc.ClientConnInterface) UrlServiceClient {
	return &urlServiceClient{cc}
}

func (c *urlServiceClient) SetUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, "/rpc.UrlService/SetUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UrlListResponse, error) {
	out := new(UrlListResponse)
	err := c.cc.Invoke(ctx, "/rpc.UrlService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) GetAllDomainAndInterval(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UrlIntervalResponse, error) {
	out := new(UrlIntervalResponse)
	err := c.cc.Invoke(ctx, "/rpc.UrlService/GetAllDomainAndInterval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) DeleteUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, "/rpc.UrlService/DeleteUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) SetProxy(ctx context.Context, in *ProxyRequest, opts ...grpc.CallOption) (*ProxyResponse, error) {
	out := new(ProxyResponse)
	err := c.cc.Invoke(ctx, "/rpc.UrlService/SetProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) GetAllProxy(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProxyListResponse, error) {
	out := new(ProxyListResponse)
	err := c.cc.Invoke(ctx, "/rpc.UrlService/GetAllProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) DeleteProxy(ctx context.Context, in *ProxyRequest, opts ...grpc.CallOption) (*ProxyResponse, error) {
	out := new(ProxyResponse)
	err := c.cc.Invoke(ctx, "/rpc.UrlService/DeleteProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlServiceServer is the server API for UrlService service.
// All implementations must embed UnimplementedUrlServiceServer
// for forward compatibility
type UrlServiceServer interface {
	SetUrl(context.Context, *UrlRequest) (*UrlResponse, error)
	GetAll(context.Context, *empty.Empty) (*UrlListResponse, error)
	GetAllDomainAndInterval(context.Context, *empty.Empty) (*UrlIntervalResponse, error)
	DeleteUrl(context.Context, *UrlRequest) (*UrlResponse, error)
	SetProxy(context.Context, *ProxyRequest) (*ProxyResponse, error)
	GetAllProxy(context.Context, *empty.Empty) (*ProxyListResponse, error)
	DeleteProxy(context.Context, *ProxyRequest) (*ProxyResponse, error)
	mustEmbedUnimplementedUrlServiceServer()
}

// UnimplementedUrlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUrlServiceServer struct {
}

func (UnimplementedUrlServiceServer) SetUrl(context.Context, *UrlRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUrl not implemented")
}
func (UnimplementedUrlServiceServer) GetAll(context.Context, *empty.Empty) (*UrlListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUrlServiceServer) GetAllDomainAndInterval(context.Context, *empty.Empty) (*UrlIntervalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDomainAndInterval not implemented")
}
func (UnimplementedUrlServiceServer) DeleteUrl(context.Context, *UrlRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUrl not implemented")
}
func (UnimplementedUrlServiceServer) SetProxy(context.Context, *ProxyRequest) (*ProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProxy not implemented")
}
func (UnimplementedUrlServiceServer) GetAllProxy(context.Context, *empty.Empty) (*ProxyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProxy not implemented")
}
func (UnimplementedUrlServiceServer) DeleteProxy(context.Context, *ProxyRequest) (*ProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProxy not implemented")
}
func (UnimplementedUrlServiceServer) mustEmbedUnimplementedUrlServiceServer() {}

// UnsafeUrlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlServiceServer will
// result in compilation errors.
type UnsafeUrlServiceServer interface {
	mustEmbedUnimplementedUrlServiceServer()
}

func RegisterUrlServiceServer(s grpc.ServiceRegistrar, srv UrlServiceServer) {
	s.RegisterService(&UrlService_ServiceDesc, srv)
}

func _UrlService_SetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).SetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UrlService/SetUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).SetUrl(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UrlService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).GetAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_GetAllDomainAndInterval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).GetAllDomainAndInterval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UrlService/GetAllDomainAndInterval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).GetAllDomainAndInterval(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_DeleteUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).DeleteUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UrlService/DeleteUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).DeleteUrl(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_SetProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).SetProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UrlService/SetProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).SetProxy(ctx, req.(*ProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_GetAllProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).GetAllProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UrlService/GetAllProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).GetAllProxy(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_DeleteProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).DeleteProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UrlService/DeleteProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).DeleteProxy(ctx, req.(*ProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlService_ServiceDesc is the grpc.ServiceDesc for UrlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.UrlService",
	HandlerType: (*UrlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUrl",
			Handler:    _UrlService_SetUrl_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UrlService_GetAll_Handler,
		},
		{
			MethodName: "GetAllDomainAndInterval",
			Handler:    _UrlService_GetAllDomainAndInterval_Handler,
		},
		{
			MethodName: "DeleteUrl",
			Handler:    _UrlService_DeleteUrl_Handler,
		},
		{
			MethodName: "SetProxy",
			Handler:    _UrlService_SetProxy_Handler,
		},
		{
			MethodName: "GetAllProxy",
			Handler:    _UrlService_GetAllProxy_Handler,
		},
		{
			MethodName: "DeleteProxy",
			Handler:    _UrlService_DeleteProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/url.proto",
}
