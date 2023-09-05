// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: service_pc.proto

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

// PCServiceClient is the client API for PCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PCServiceClient interface {
	CreatePC(ctx context.Context, in *CreatePCRequest, opts ...grpc.CallOption) (*CreatePCResponse, error)
	SearchPC(ctx context.Context, in *SearchPCRequest, opts ...grpc.CallOption) (*SearchPCResponse, error)
}

type pCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPCServiceClient(cc grpc.ClientConnInterface) PCServiceClient {
	return &pCServiceClient{cc}
}

func (c *pCServiceClient) CreatePC(ctx context.Context, in *CreatePCRequest, opts ...grpc.CallOption) (*CreatePCResponse, error) {
	out := new(CreatePCResponse)
	err := c.cc.Invoke(ctx, "/pcbook.PCService/CreatePC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCServiceClient) SearchPC(ctx context.Context, in *SearchPCRequest, opts ...grpc.CallOption) (*SearchPCResponse, error) {
	out := new(SearchPCResponse)
	err := c.cc.Invoke(ctx, "/pcbook.PCService/SearchPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PCServiceServer is the server API for PCService service.
// All implementations must embed UnimplementedPCServiceServer
// for forward compatibility
type PCServiceServer interface {
	CreatePC(context.Context, *CreatePCRequest) (*CreatePCResponse, error)
	SearchPC(context.Context, *SearchPCRequest) (*SearchPCResponse, error)
	mustEmbedUnimplementedPCServiceServer()
}

// UnimplementedPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPCServiceServer struct {
}

func (UnimplementedPCServiceServer) CreatePC(context.Context, *CreatePCRequest) (*CreatePCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePC not implemented")
}
func (UnimplementedPCServiceServer) SearchPC(context.Context, *SearchPCRequest) (*SearchPCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPC not implemented")
}
func (UnimplementedPCServiceServer) mustEmbedUnimplementedPCServiceServer() {}

// UnsafePCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PCServiceServer will
// result in compilation errors.
type UnsafePCServiceServer interface {
	mustEmbedUnimplementedPCServiceServer()
}

func RegisterPCServiceServer(s grpc.ServiceRegistrar, srv PCServiceServer) {
	s.RegisterService(&PCService_ServiceDesc, srv)
}

func _PCService_CreatePC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCServiceServer).CreatePC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcbook.PCService/CreatePC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCServiceServer).CreatePC(ctx, req.(*CreatePCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCService_SearchPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCServiceServer).SearchPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcbook.PCService/SearchPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCServiceServer).SearchPC(ctx, req.(*SearchPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PCService_ServiceDesc is the grpc.ServiceDesc for PCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pcbook.PCService",
	HandlerType: (*PCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePC",
			Handler:    _PCService_CreatePC_Handler,
		},
		{
			MethodName: "SearchPC",
			Handler:    _PCService_SearchPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_pc.proto",
}
