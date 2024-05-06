// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: front/service/v1/i_tag.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-cms/gen/api/go/content/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TagService_ListTag_FullMethodName   = "/front.service.v1.TagService/ListTag"
	TagService_GetTag_FullMethodName    = "/front.service.v1.TagService/GetTag"
	TagService_CreateTag_FullMethodName = "/front.service.v1.TagService/CreateTag"
	TagService_UpdateTag_FullMethodName = "/front.service.v1.TagService/UpdateTag"
	TagService_DeleteTag_FullMethodName = "/front.service.v1.TagService/DeleteTag"
)

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	// 获取标签列表
	ListTag(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListTagResponse, error)
	// 获取标签数据
	GetTag(ctx context.Context, in *v11.GetTagRequest, opts ...grpc.CallOption) (*v11.Tag, error)
	// 创建标签
	CreateTag(ctx context.Context, in *v11.CreateTagRequest, opts ...grpc.CallOption) (*v11.Tag, error)
	// 更新标签
	UpdateTag(ctx context.Context, in *v11.UpdateTagRequest, opts ...grpc.CallOption) (*v11.Tag, error)
	// 删除标签
	DeleteTag(ctx context.Context, in *v11.DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) ListTag(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListTagResponse, error) {
	out := new(v11.ListTagResponse)
	err := c.cc.Invoke(ctx, TagService_ListTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) GetTag(ctx context.Context, in *v11.GetTagRequest, opts ...grpc.CallOption) (*v11.Tag, error) {
	out := new(v11.Tag)
	err := c.cc.Invoke(ctx, TagService_GetTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) CreateTag(ctx context.Context, in *v11.CreateTagRequest, opts ...grpc.CallOption) (*v11.Tag, error) {
	out := new(v11.Tag)
	err := c.cc.Invoke(ctx, TagService_CreateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) UpdateTag(ctx context.Context, in *v11.UpdateTagRequest, opts ...grpc.CallOption) (*v11.Tag, error) {
	out := new(v11.Tag)
	err := c.cc.Invoke(ctx, TagService_UpdateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) DeleteTag(ctx context.Context, in *v11.DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TagService_DeleteTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	// 获取标签列表
	ListTag(context.Context, *v1.PagingRequest) (*v11.ListTagResponse, error)
	// 获取标签数据
	GetTag(context.Context, *v11.GetTagRequest) (*v11.Tag, error)
	// 创建标签
	CreateTag(context.Context, *v11.CreateTagRequest) (*v11.Tag, error)
	// 更新标签
	UpdateTag(context.Context, *v11.UpdateTagRequest) (*v11.Tag, error)
	// 删除标签
	DeleteTag(context.Context, *v11.DeleteTagRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) ListTag(context.Context, *v1.PagingRequest) (*v11.ListTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTag not implemented")
}
func (UnimplementedTagServiceServer) GetTag(context.Context, *v11.GetTagRequest) (*v11.Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedTagServiceServer) CreateTag(context.Context, *v11.CreateTagRequest) (*v11.Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagServiceServer) UpdateTag(context.Context, *v11.UpdateTagRequest) (*v11.Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedTagServiceServer) DeleteTag(context.Context, *v11.DeleteTagRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_ListTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).ListTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_ListTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).ListTag(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_GetTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTag(ctx, req.(*v11.GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_CreateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).CreateTag(ctx, req.(*v11.CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).UpdateTag(ctx, req.(*v11.UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_DeleteTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).DeleteTag(ctx, req.(*v11.DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "front.service.v1.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTag",
			Handler:    _TagService_ListTag_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _TagService_GetTag_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _TagService_CreateTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _TagService_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _TagService_DeleteTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "front/service/v1/i_tag.proto",
}