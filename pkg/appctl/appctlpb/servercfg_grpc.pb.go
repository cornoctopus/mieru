// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package appctlpb

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

// ServerConfigServiceClient is the client API for ServerConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerConfigServiceClient interface {
	// Fetch the server config.
	GetConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerConfig, error)
	// Update server config.
	SetConfig(ctx context.Context, in *ServerConfig, opts ...grpc.CallOption) (*ServerConfig, error)
}

type serverConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerConfigServiceClient(cc grpc.ClientConnInterface) ServerConfigServiceClient {
	return &serverConfigServiceClient{cc}
}

func (c *serverConfigServiceClient) GetConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerConfig, error) {
	out := new(ServerConfig)
	err := c.cc.Invoke(ctx, "/appctl.ServerConfigService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverConfigServiceClient) SetConfig(ctx context.Context, in *ServerConfig, opts ...grpc.CallOption) (*ServerConfig, error) {
	out := new(ServerConfig)
	err := c.cc.Invoke(ctx, "/appctl.ServerConfigService/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerConfigServiceServer is the server API for ServerConfigService service.
// All implementations must embed UnimplementedServerConfigServiceServer
// for forward compatibility
type ServerConfigServiceServer interface {
	// Fetch the server config.
	GetConfig(context.Context, *Empty) (*ServerConfig, error)
	// Update server config.
	SetConfig(context.Context, *ServerConfig) (*ServerConfig, error)
	mustEmbedUnimplementedServerConfigServiceServer()
}

// UnimplementedServerConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerConfigServiceServer struct {
}

func (UnimplementedServerConfigServiceServer) GetConfig(context.Context, *Empty) (*ServerConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedServerConfigServiceServer) SetConfig(context.Context, *ServerConfig) (*ServerConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedServerConfigServiceServer) mustEmbedUnimplementedServerConfigServiceServer() {}

// UnsafeServerConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerConfigServiceServer will
// result in compilation errors.
type UnsafeServerConfigServiceServer interface {
	mustEmbedUnimplementedServerConfigServiceServer()
}

func RegisterServerConfigServiceServer(s grpc.ServiceRegistrar, srv ServerConfigServiceServer) {
	s.RegisterService(&ServerConfigService_ServiceDesc, srv)
}

func _ServerConfigService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerConfigServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerConfigService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerConfigServiceServer).GetConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerConfigService_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerConfigServiceServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appctl.ServerConfigService/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerConfigServiceServer).SetConfig(ctx, req.(*ServerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerConfigService_ServiceDesc is the grpc.ServiceDesc for ServerConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appctl.ServerConfigService",
	HandlerType: (*ServerConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ServerConfigService_GetConfig_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _ServerConfigService_SetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servercfg.proto",
}