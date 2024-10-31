// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: protos/teacher.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TeacherService_CreateNewTeacher_FullMethodName = "/auth.TeacherService/CreateNewTeacher"
	TeacherService_GetAllTeacher_FullMethodName    = "/auth.TeacherService/GetAllTeacher"
	TeacherService_UpdateTeacher_FullMethodName    = "/auth.TeacherService/UpdateTeacher"
	TeacherService_DeleteTeacher_FullMethodName    = "/auth.TeacherService/DeleteTeacher"
)

// TeacherServiceClient is the client API for TeacherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeacherServiceClient interface {
	CreateNewTeacher(ctx context.Context, in *CreateTeacherRequest, opts ...grpc.CallOption) (*CreateTeacherResponse, error)
	GetAllTeacher(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllTeacherResponse, error)
	UpdateTeacher(ctx context.Context, in *UpdateTeacherRequest, opts ...grpc.CallOption) (*UpdateTeacherResponse, error)
	DeleteTeacher(ctx context.Context, in *DeleteTeacherRequest, opts ...grpc.CallOption) (*DeleteTeacherResponse, error)
}

type teacherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeacherServiceClient(cc grpc.ClientConnInterface) TeacherServiceClient {
	return &teacherServiceClient{cc}
}

func (c *teacherServiceClient) CreateNewTeacher(ctx context.Context, in *CreateTeacherRequest, opts ...grpc.CallOption) (*CreateTeacherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_CreateNewTeacher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetAllTeacher(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllTeacherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_GetAllTeacher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) UpdateTeacher(ctx context.Context, in *UpdateTeacherRequest, opts ...grpc.CallOption) (*UpdateTeacherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_UpdateTeacher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) DeleteTeacher(ctx context.Context, in *DeleteTeacherRequest, opts ...grpc.CallOption) (*DeleteTeacherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTeacherResponse)
	err := c.cc.Invoke(ctx, TeacherService_DeleteTeacher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeacherServiceServer is the server API for TeacherService service.
// All implementations must embed UnimplementedTeacherServiceServer
// for forward compatibility.
type TeacherServiceServer interface {
	CreateNewTeacher(context.Context, *CreateTeacherRequest) (*CreateTeacherResponse, error)
	GetAllTeacher(context.Context, *emptypb.Empty) (*GetAllTeacherResponse, error)
	UpdateTeacher(context.Context, *UpdateTeacherRequest) (*UpdateTeacherResponse, error)
	DeleteTeacher(context.Context, *DeleteTeacherRequest) (*DeleteTeacherResponse, error)
	mustEmbedUnimplementedTeacherServiceServer()
}

// UnimplementedTeacherServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTeacherServiceServer struct{}

func (UnimplementedTeacherServiceServer) CreateNewTeacher(context.Context, *CreateTeacherRequest) (*CreateTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) GetAllTeacher(context.Context, *emptypb.Empty) (*GetAllTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) UpdateTeacher(context.Context, *UpdateTeacherRequest) (*UpdateTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) DeleteTeacher(context.Context, *DeleteTeacherRequest) (*DeleteTeacherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) mustEmbedUnimplementedTeacherServiceServer() {}
func (UnimplementedTeacherServiceServer) testEmbeddedByValue()                        {}

// UnsafeTeacherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeacherServiceServer will
// result in compilation errors.
type UnsafeTeacherServiceServer interface {
	mustEmbedUnimplementedTeacherServiceServer()
}

func RegisterTeacherServiceServer(s grpc.ServiceRegistrar, srv TeacherServiceServer) {
	// If the following call pancis, it indicates UnimplementedTeacherServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TeacherService_ServiceDesc, srv)
}

func _TeacherService_CreateNewTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).CreateNewTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_CreateNewTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).CreateNewTeacher(ctx, req.(*CreateTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetAllTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetAllTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_GetAllTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetAllTeacher(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_UpdateTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).UpdateTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_UpdateTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).UpdateTeacher(ctx, req.(*UpdateTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_DeleteTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).DeleteTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeacherService_DeleteTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).DeleteTeacher(ctx, req.(*DeleteTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeacherService_ServiceDesc is the grpc.ServiceDesc for TeacherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeacherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.TeacherService",
	HandlerType: (*TeacherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewTeacher",
			Handler:    _TeacherService_CreateNewTeacher_Handler,
		},
		{
			MethodName: "GetAllTeacher",
			Handler:    _TeacherService_GetAllTeacher_Handler,
		},
		{
			MethodName: "UpdateTeacher",
			Handler:    _TeacherService_UpdateTeacher_Handler,
		},
		{
			MethodName: "DeleteTeacher",
			Handler:    _TeacherService_DeleteTeacher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/teacher.proto",
}
