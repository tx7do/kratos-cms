// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             (unknown)
// source: front/service/v1/i_photo.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-blog/gen/api/go/common/pagination"
	v1 "kratos-blog/gen/api/go/content/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPhotoServiceCreatePhoto = "/front.service.v1.PhotoService/CreatePhoto"
const OperationPhotoServiceDeletePhoto = "/front.service.v1.PhotoService/DeletePhoto"
const OperationPhotoServiceGetPhoto = "/front.service.v1.PhotoService/GetPhoto"
const OperationPhotoServiceListPhoto = "/front.service.v1.PhotoService/ListPhoto"
const OperationPhotoServiceUpdatePhoto = "/front.service.v1.PhotoService/UpdatePhoto"

type PhotoServiceHTTPServer interface {
	CreatePhoto(context.Context, *v1.CreatePhotoRequest) (*v1.Photo, error)
	DeletePhoto(context.Context, *v1.DeletePhotoRequest) (*emptypb.Empty, error)
	GetPhoto(context.Context, *v1.GetPhotoRequest) (*v1.Photo, error)
	ListPhoto(context.Context, *pagination.PagingRequest) (*v1.ListPhotoResponse, error)
	UpdatePhoto(context.Context, *v1.UpdatePhotoRequest) (*v1.Photo, error)
}

func RegisterPhotoServiceHTTPServer(s *http.Server, srv PhotoServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/photos", _PhotoService_ListPhoto0_HTTP_Handler(srv))
	r.GET("/blog/v1/photos/{id}", _PhotoService_GetPhoto0_HTTP_Handler(srv))
	r.POST("/blog/v1/photos", _PhotoService_CreatePhoto0_HTTP_Handler(srv))
	r.PUT("/blog/v1/photos/{id}", _PhotoService_UpdatePhoto0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/photos/{id}", _PhotoService_DeletePhoto0_HTTP_Handler(srv))
}

func _PhotoService_ListPhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceListPhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPhoto(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListPhotoResponse)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_GetPhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetPhotoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceGetPhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPhoto(ctx, req.(*v1.GetPhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Photo)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_CreatePhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreatePhotoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceCreatePhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePhoto(ctx, req.(*v1.CreatePhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Photo)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_UpdatePhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdatePhotoRequest
		if err := ctx.Bind(&in.Photo); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceUpdatePhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePhoto(ctx, req.(*v1.UpdatePhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Photo)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_DeletePhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeletePhotoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceDeletePhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePhoto(ctx, req.(*v1.DeletePhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type PhotoServiceHTTPClient interface {
	CreatePhoto(ctx context.Context, req *v1.CreatePhotoRequest, opts ...http.CallOption) (rsp *v1.Photo, err error)
	DeletePhoto(ctx context.Context, req *v1.DeletePhotoRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetPhoto(ctx context.Context, req *v1.GetPhotoRequest, opts ...http.CallOption) (rsp *v1.Photo, err error)
	ListPhoto(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListPhotoResponse, err error)
	UpdatePhoto(ctx context.Context, req *v1.UpdatePhotoRequest, opts ...http.CallOption) (rsp *v1.Photo, err error)
}

type PhotoServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPhotoServiceHTTPClient(client *http.Client) PhotoServiceHTTPClient {
	return &PhotoServiceHTTPClientImpl{client}
}

func (c *PhotoServiceHTTPClientImpl) CreatePhoto(ctx context.Context, in *v1.CreatePhotoRequest, opts ...http.CallOption) (*v1.Photo, error) {
	var out v1.Photo
	pattern := "/blog/v1/photos"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPhotoServiceCreatePhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) DeletePhoto(ctx context.Context, in *v1.DeletePhotoRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/photos/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPhotoServiceDeletePhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) GetPhoto(ctx context.Context, in *v1.GetPhotoRequest, opts ...http.CallOption) (*v1.Photo, error) {
	var out v1.Photo
	pattern := "/blog/v1/photos/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPhotoServiceGetPhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) ListPhoto(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListPhotoResponse, error) {
	var out v1.ListPhotoResponse
	pattern := "/blog/v1/photos"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPhotoServiceListPhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) UpdatePhoto(ctx context.Context, in *v1.UpdatePhotoRequest, opts ...http.CallOption) (*v1.Photo, error) {
	var out v1.Photo
	pattern := "/blog/v1/photos/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPhotoServiceUpdatePhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Photo, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
