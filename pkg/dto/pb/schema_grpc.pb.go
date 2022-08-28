// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: schema.proto

package pb

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

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldClient interface {
	GetGreeting(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Answer, error)
	FileForDownload(ctx context.Context, in *FileReq, opts ...grpc.CallOption) (*FileRpl, error)
}

type helloWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldClient(cc grpc.ClientConnInterface) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) GetGreeting(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.HelloWorld/GetGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) FileForDownload(ctx context.Context, in *FileReq, opts ...grpc.CallOption) (*FileRpl, error) {
	out := new(FileRpl)
	err := c.cc.Invoke(ctx, "/pb.HelloWorld/FileForDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
// All implementations should embed UnimplementedHelloWorldServer
// for forward compatibility
type HelloWorldServer interface {
	GetGreeting(context.Context, *Input) (*Answer, error)
	FileForDownload(context.Context, *FileReq) (*FileRpl, error)
}

// UnimplementedHelloWorldServer should be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (UnimplementedHelloWorldServer) GetGreeting(context.Context, *Input) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeting not implemented")
}
func (UnimplementedHelloWorldServer) FileForDownload(context.Context, *FileReq) (*FileRpl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileForDownload not implemented")
}

// UnsafeHelloWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldServer will
// result in compilation errors.
type UnsafeHelloWorldServer interface {
	mustEmbedUnimplementedHelloWorldServer()
}

func RegisterHelloWorldServer(s grpc.ServiceRegistrar, srv HelloWorldServer) {
	s.RegisterService(&HelloWorld_ServiceDesc, srv)
}

func _HelloWorld_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloWorld/GetGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GetGreeting(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_FileForDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).FileForDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloWorld/FileForDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).FileForDownload(ctx, req.(*FileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloWorld_ServiceDesc is the grpc.ServiceDesc for HelloWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGreeting",
			Handler:    _HelloWorld_GetGreeting_Handler,
		},
		{
			MethodName: "FileForDownload",
			Handler:    _HelloWorld_FileForDownload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema.proto",
}