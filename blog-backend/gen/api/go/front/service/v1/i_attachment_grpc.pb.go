// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: front/service/v1/i_attachment.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-blog/gen/api/go/common/pagination"
	v1 "kratos-blog/gen/api/go/file/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AttachmentServiceClient is the client API for AttachmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttachmentServiceClient interface {
	// 获取附件列表
	ListAttachment(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*v1.ListAttachmentResponse, error)
	// 获取附件数据
	GetAttachment(ctx context.Context, in *v1.GetAttachmentRequest, opts ...grpc.CallOption) (*v1.Attachment, error)
	// 创建附件
	CreateAttachment(ctx context.Context, in *v1.CreateAttachmentRequest, opts ...grpc.CallOption) (*v1.Attachment, error)
	// 更新附件
	UpdateAttachment(ctx context.Context, in *v1.UpdateAttachmentRequest, opts ...grpc.CallOption) (*v1.Attachment, error)
	// 删除附件
	DeleteAttachment(ctx context.Context, in *v1.DeleteAttachmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type attachmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttachmentServiceClient(cc grpc.ClientConnInterface) AttachmentServiceClient {
	return &attachmentServiceClient{cc}
}

func (c *attachmentServiceClient) ListAttachment(ctx context.Context, in *pagination.PagingRequest, opts ...grpc.CallOption) (*v1.ListAttachmentResponse, error) {
	out := new(v1.ListAttachmentResponse)
	err := c.cc.Invoke(ctx, "/front.service.v1.AttachmentService/ListAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) GetAttachment(ctx context.Context, in *v1.GetAttachmentRequest, opts ...grpc.CallOption) (*v1.Attachment, error) {
	out := new(v1.Attachment)
	err := c.cc.Invoke(ctx, "/front.service.v1.AttachmentService/GetAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) CreateAttachment(ctx context.Context, in *v1.CreateAttachmentRequest, opts ...grpc.CallOption) (*v1.Attachment, error) {
	out := new(v1.Attachment)
	err := c.cc.Invoke(ctx, "/front.service.v1.AttachmentService/CreateAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) UpdateAttachment(ctx context.Context, in *v1.UpdateAttachmentRequest, opts ...grpc.CallOption) (*v1.Attachment, error) {
	out := new(v1.Attachment)
	err := c.cc.Invoke(ctx, "/front.service.v1.AttachmentService/UpdateAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentServiceClient) DeleteAttachment(ctx context.Context, in *v1.DeleteAttachmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/front.service.v1.AttachmentService/DeleteAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachmentServiceServer is the server API for AttachmentService service.
// All implementations must embed UnimplementedAttachmentServiceServer
// for forward compatibility
type AttachmentServiceServer interface {
	// 获取附件列表
	ListAttachment(context.Context, *pagination.PagingRequest) (*v1.ListAttachmentResponse, error)
	// 获取附件数据
	GetAttachment(context.Context, *v1.GetAttachmentRequest) (*v1.Attachment, error)
	// 创建附件
	CreateAttachment(context.Context, *v1.CreateAttachmentRequest) (*v1.Attachment, error)
	// 更新附件
	UpdateAttachment(context.Context, *v1.UpdateAttachmentRequest) (*v1.Attachment, error)
	// 删除附件
	DeleteAttachment(context.Context, *v1.DeleteAttachmentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAttachmentServiceServer()
}

// UnimplementedAttachmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAttachmentServiceServer struct {
}

func (UnimplementedAttachmentServiceServer) ListAttachment(context.Context, *pagination.PagingRequest) (*v1.ListAttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) GetAttachment(context.Context, *v1.GetAttachmentRequest) (*v1.Attachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) CreateAttachment(context.Context, *v1.CreateAttachmentRequest) (*v1.Attachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) UpdateAttachment(context.Context, *v1.UpdateAttachmentRequest) (*v1.Attachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) DeleteAttachment(context.Context, *v1.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttachment not implemented")
}
func (UnimplementedAttachmentServiceServer) mustEmbedUnimplementedAttachmentServiceServer() {}

// UnsafeAttachmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttachmentServiceServer will
// result in compilation errors.
type UnsafeAttachmentServiceServer interface {
	mustEmbedUnimplementedAttachmentServiceServer()
}

func RegisterAttachmentServiceServer(s grpc.ServiceRegistrar, srv AttachmentServiceServer) {
	s.RegisterService(&AttachmentService_ServiceDesc, srv)
}

func _AttachmentService_ListAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).ListAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/front.service.v1.AttachmentService/ListAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).ListAttachment(ctx, req.(*pagination.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_GetAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).GetAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/front.service.v1.AttachmentService/GetAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).GetAttachment(ctx, req.(*v1.GetAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_CreateAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).CreateAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/front.service.v1.AttachmentService/CreateAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).CreateAttachment(ctx, req.(*v1.CreateAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_UpdateAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).UpdateAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/front.service.v1.AttachmentService/UpdateAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).UpdateAttachment(ctx, req.(*v1.UpdateAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttachmentService_DeleteAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentServiceServer).DeleteAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/front.service.v1.AttachmentService/DeleteAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentServiceServer).DeleteAttachment(ctx, req.(*v1.DeleteAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttachmentService_ServiceDesc is the grpc.ServiceDesc for AttachmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttachmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "front.service.v1.AttachmentService",
	HandlerType: (*AttachmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAttachment",
			Handler:    _AttachmentService_ListAttachment_Handler,
		},
		{
			MethodName: "GetAttachment",
			Handler:    _AttachmentService_GetAttachment_Handler,
		},
		{
			MethodName: "CreateAttachment",
			Handler:    _AttachmentService_CreateAttachment_Handler,
		},
		{
			MethodName: "UpdateAttachment",
			Handler:    _AttachmentService_UpdateAttachment_Handler,
		},
		{
			MethodName: "DeleteAttachment",
			Handler:    _AttachmentService_DeleteAttachment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "front/service/v1/i_attachment.proto",
}
