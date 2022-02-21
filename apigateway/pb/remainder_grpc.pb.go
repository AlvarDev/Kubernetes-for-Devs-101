// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// ReminderServiceClient is the client API for ReminderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReminderServiceClient interface {
	GetNewId(ctx context.Context, in *GetNewIdRequest, opts ...grpc.CallOption) (*GetNewIdResponse, error)
}

type reminderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReminderServiceClient(cc grpc.ClientConnInterface) ReminderServiceClient {
	return &reminderServiceClient{cc}
}

func (c *reminderServiceClient) GetNewId(ctx context.Context, in *GetNewIdRequest, opts ...grpc.CallOption) (*GetNewIdResponse, error) {
	out := new(GetNewIdResponse)
	err := c.cc.Invoke(ctx, "/remainderprotos.ReminderService/GetNewId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReminderServiceServer is the server API for ReminderService service.
// All implementations must embed UnimplementedReminderServiceServer
// for forward compatibility
type ReminderServiceServer interface {
	GetNewId(context.Context, *GetNewIdRequest) (*GetNewIdResponse, error)
	mustEmbedUnimplementedReminderServiceServer()
}

// UnimplementedReminderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReminderServiceServer struct {
}

func (UnimplementedReminderServiceServer) GetNewId(context.Context, *GetNewIdRequest) (*GetNewIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewId not implemented")
}
func (UnimplementedReminderServiceServer) mustEmbedUnimplementedReminderServiceServer() {}

// UnsafeReminderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReminderServiceServer will
// result in compilation errors.
type UnsafeReminderServiceServer interface {
	mustEmbedUnimplementedReminderServiceServer()
}

func RegisterReminderServiceServer(s grpc.ServiceRegistrar, srv ReminderServiceServer) {
	s.RegisterService(&ReminderService_ServiceDesc, srv)
}

func _ReminderService_GetNewId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).GetNewId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remainderprotos.ReminderService/GetNewId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).GetNewId(ctx, req.(*GetNewIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReminderService_ServiceDesc is the grpc.ServiceDesc for ReminderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReminderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remainderprotos.ReminderService",
	HandlerType: (*ReminderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewId",
			Handler:    _ReminderService_GetNewId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remainder.proto",
}
