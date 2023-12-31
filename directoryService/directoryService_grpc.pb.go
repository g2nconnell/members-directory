// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: directoryService.proto

package directoryService

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

const (
	DirectoryService_CreateMember_FullMethodName = "/DirectoryService/CreateMember"
	DirectoryService_DeleteMember_FullMethodName = "/DirectoryService/DeleteMember"
	DirectoryService_UpdateMember_FullMethodName = "/DirectoryService/UpdateMember"
	DirectoryService_GetMember_FullMethodName    = "/DirectoryService/GetMember"
	DirectoryService_ListMembers_FullMethodName  = "/DirectoryService/ListMembers"
)

// DirectoryServiceClient is the client API for DirectoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryServiceClient interface {
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error)
	DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*DeleteMemberResponse, error)
	UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*UpdateMemberResponse, error)
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberResponse, error)
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error)
}

type directoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryServiceClient(cc grpc.ClientConnInterface) DirectoryServiceClient {
	return &directoryServiceClient{cc}
}

func (c *directoryServiceClient) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error) {
	out := new(CreateMemberResponse)
	err := c.cc.Invoke(ctx, DirectoryService_CreateMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*DeleteMemberResponse, error) {
	out := new(DeleteMemberResponse)
	err := c.cc.Invoke(ctx, DirectoryService_DeleteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*UpdateMemberResponse, error) {
	out := new(UpdateMemberResponse)
	err := c.cc.Invoke(ctx, DirectoryService_UpdateMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberResponse, error) {
	out := new(GetMemberResponse)
	err := c.cc.Invoke(ctx, DirectoryService_GetMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error) {
	out := new(ListMembersResponse)
	err := c.cc.Invoke(ctx, DirectoryService_ListMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryServiceServer is the server API for DirectoryService service.
// All implementations must embed UnimplementedDirectoryServiceServer
// for forward compatibility
type DirectoryServiceServer interface {
	CreateMember(context.Context, *CreateMemberRequest) (*CreateMemberResponse, error)
	DeleteMember(context.Context, *DeleteMemberRequest) (*DeleteMemberResponse, error)
	UpdateMember(context.Context, *UpdateMemberRequest) (*UpdateMemberResponse, error)
	GetMember(context.Context, *GetMemberRequest) (*GetMemberResponse, error)
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
	mustEmbedUnimplementedDirectoryServiceServer()
}

// UnimplementedDirectoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryServiceServer struct {
}

func (UnimplementedDirectoryServiceServer) CreateMember(context.Context, *CreateMemberRequest) (*CreateMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (UnimplementedDirectoryServiceServer) DeleteMember(context.Context, *DeleteMemberRequest) (*DeleteMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedDirectoryServiceServer) UpdateMember(context.Context, *UpdateMemberRequest) (*UpdateMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMember not implemented")
}
func (UnimplementedDirectoryServiceServer) GetMember(context.Context, *GetMemberRequest) (*GetMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (UnimplementedDirectoryServiceServer) ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedDirectoryServiceServer) mustEmbedUnimplementedDirectoryServiceServer() {}

// UnsafeDirectoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryServiceServer will
// result in compilation errors.
type UnsafeDirectoryServiceServer interface {
	mustEmbedUnimplementedDirectoryServiceServer()
}

func RegisterDirectoryServiceServer(s grpc.ServiceRegistrar, srv DirectoryServiceServer) {
	s.RegisterService(&DirectoryService_ServiceDesc, srv)
}

func _DirectoryService_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectoryService_CreateMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).CreateMember(ctx, req.(*CreateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectoryService_DeleteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).DeleteMember(ctx, req.(*DeleteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_UpdateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).UpdateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectoryService_UpdateMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).UpdateMember(ctx, req.(*UpdateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectoryService_GetMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).GetMember(ctx, req.(*GetMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectoryService_ListMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).ListMembers(ctx, req.(*ListMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectoryService_ServiceDesc is the grpc.ServiceDesc for DirectoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DirectoryService",
	HandlerType: (*DirectoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMember",
			Handler:    _DirectoryService_CreateMember_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _DirectoryService_DeleteMember_Handler,
		},
		{
			MethodName: "UpdateMember",
			Handler:    _DirectoryService_UpdateMember_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _DirectoryService_GetMember_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _DirectoryService_ListMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directoryService.proto",
}
