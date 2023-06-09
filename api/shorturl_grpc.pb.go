// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: api/shorturl.proto

package proto

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
	IShorturl_Post_FullMethodName        = "/proto.IShorturl/Post"
	IShorturl_LongLink_FullMethodName    = "/proto.IShorturl/LongLink"
	IShorturl_ShortLink_FullMethodName   = "/proto.IShorturl/ShortLink"
	IShorturl_UserAllLink_FullMethodName = "/proto.IShorturl/UserAllLink"
	IShorturl_AllLink_FullMethodName     = "/proto.IShorturl/AllLink"
	IShorturl_AllUsers_FullMethodName    = "/proto.IShorturl/AllUsers"
	IShorturl_UserDelLink_FullMethodName = "/proto.IShorturl/UserDelLink"
	IShorturl_ReadService_FullMethodName = "/proto.IShorturl/ReadService"
	IShorturl_SaveService_FullMethodName = "/proto.IShorturl/SaveService"
)

// IShorturlClient is the client API for IShorturl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IShorturlClient interface {
	// Post(context.Context, *entity.Shorturl) error
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	// LongLink(context.Context, *entity.Shorturl) (string, error)
	LongLink(ctx context.Context, in *LongLinkRequest, opts ...grpc.CallOption) (*LongLinkResponse, error)
	// ShortLink(context.Context, *entity.Shorturl) (*entity.Shorturl, error)
	ShortLink(ctx context.Context, in *ShortLinkRequest, opts ...grpc.CallOption) (*ShortLinkResponse, error)
	// UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error)
	UserAllLink(ctx context.Context, in *UserAllLinkRequest, opts ...grpc.CallOption) (*UserAllLinkResponse, error)
	// AllLink() (entity.CountURLs, error)
	AllLink(ctx context.Context, in *AllLinkRequest, opts ...grpc.CallOption) (*AllLinkResponse, error)
	// AllUsers() (entity.CountUsers, error)
	AllUsers(ctx context.Context, in *AllUsersRequest, opts ...grpc.CallOption) (*AllUsersResponse, error)
	// UserDelLink(ctx context.Context, u *entity.User) error
	UserDelLink(ctx context.Context, in *UserDelLinkRequest, opts ...grpc.CallOption) (*UserDelLinkResponse, error)
	// ReadService() error
	ReadService(ctx context.Context, in *ReadServiceRequest, opts ...grpc.CallOption) (*ReadServiceResponse, error)
	// SaveService() error
	SaveService(ctx context.Context, in *SaveServiceRequest, opts ...grpc.CallOption) (*SaveServiceResponse, error)
}

type iShorturlClient struct {
	cc grpc.ClientConnInterface
}

func NewIShorturlClient(cc grpc.ClientConnInterface) IShorturlClient {
	return &iShorturlClient{cc}
}

func (c *iShorturlClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, IShorturl_Post_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) LongLink(ctx context.Context, in *LongLinkRequest, opts ...grpc.CallOption) (*LongLinkResponse, error) {
	out := new(LongLinkResponse)
	err := c.cc.Invoke(ctx, IShorturl_LongLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) ShortLink(ctx context.Context, in *ShortLinkRequest, opts ...grpc.CallOption) (*ShortLinkResponse, error) {
	out := new(ShortLinkResponse)
	err := c.cc.Invoke(ctx, IShorturl_ShortLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) UserAllLink(ctx context.Context, in *UserAllLinkRequest, opts ...grpc.CallOption) (*UserAllLinkResponse, error) {
	out := new(UserAllLinkResponse)
	err := c.cc.Invoke(ctx, IShorturl_UserAllLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) AllLink(ctx context.Context, in *AllLinkRequest, opts ...grpc.CallOption) (*AllLinkResponse, error) {
	out := new(AllLinkResponse)
	err := c.cc.Invoke(ctx, IShorturl_AllLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) AllUsers(ctx context.Context, in *AllUsersRequest, opts ...grpc.CallOption) (*AllUsersResponse, error) {
	out := new(AllUsersResponse)
	err := c.cc.Invoke(ctx, IShorturl_AllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) UserDelLink(ctx context.Context, in *UserDelLinkRequest, opts ...grpc.CallOption) (*UserDelLinkResponse, error) {
	out := new(UserDelLinkResponse)
	err := c.cc.Invoke(ctx, IShorturl_UserDelLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) ReadService(ctx context.Context, in *ReadServiceRequest, opts ...grpc.CallOption) (*ReadServiceResponse, error) {
	out := new(ReadServiceResponse)
	err := c.cc.Invoke(ctx, IShorturl_ReadService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iShorturlClient) SaveService(ctx context.Context, in *SaveServiceRequest, opts ...grpc.CallOption) (*SaveServiceResponse, error) {
	out := new(SaveServiceResponse)
	err := c.cc.Invoke(ctx, IShorturl_SaveService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IShorturlServer is the server API for IShorturl service.
// All implementations must embed UnimplementedIShorturlServer
// for forward compatibility
type IShorturlServer interface {
	// Post(context.Context, *entity.Shorturl) error
	Post(context.Context, *PostRequest) (*PostResponse, error)
	// LongLink(context.Context, *entity.Shorturl) (string, error)
	LongLink(context.Context, *LongLinkRequest) (*LongLinkResponse, error)
	// ShortLink(context.Context, *entity.Shorturl) (*entity.Shorturl, error)
	ShortLink(context.Context, *ShortLinkRequest) (*ShortLinkResponse, error)
	// UserAllLink(ctx context.Context, u *entity.User) (*entity.User, error)
	UserAllLink(context.Context, *UserAllLinkRequest) (*UserAllLinkResponse, error)
	// AllLink() (entity.CountURLs, error)
	AllLink(context.Context, *AllLinkRequest) (*AllLinkResponse, error)
	// AllUsers() (entity.CountUsers, error)
	AllUsers(context.Context, *AllUsersRequest) (*AllUsersResponse, error)
	// UserDelLink(ctx context.Context, u *entity.User) error
	UserDelLink(context.Context, *UserDelLinkRequest) (*UserDelLinkResponse, error)
	// ReadService() error
	ReadService(context.Context, *ReadServiceRequest) (*ReadServiceResponse, error)
	// SaveService() error
	SaveService(context.Context, *SaveServiceRequest) (*SaveServiceResponse, error)
	mustEmbedUnimplementedIShorturlServer()
}

// UnimplementedIShorturlServer must be embedded to have forward compatible implementations.
type UnimplementedIShorturlServer struct {
}

func (UnimplementedIShorturlServer) Post(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedIShorturlServer) LongLink(context.Context, *LongLinkRequest) (*LongLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LongLink not implemented")
}
func (UnimplementedIShorturlServer) ShortLink(context.Context, *ShortLinkRequest) (*ShortLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortLink not implemented")
}
func (UnimplementedIShorturlServer) UserAllLink(context.Context, *UserAllLinkRequest) (*UserAllLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAllLink not implemented")
}
func (UnimplementedIShorturlServer) AllLink(context.Context, *AllLinkRequest) (*AllLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllLink not implemented")
}
func (UnimplementedIShorturlServer) AllUsers(context.Context, *AllUsersRequest) (*AllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllUsers not implemented")
}
func (UnimplementedIShorturlServer) UserDelLink(context.Context, *UserDelLinkRequest) (*UserDelLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelLink not implemented")
}
func (UnimplementedIShorturlServer) ReadService(context.Context, *ReadServiceRequest) (*ReadServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadService not implemented")
}
func (UnimplementedIShorturlServer) SaveService(context.Context, *SaveServiceRequest) (*SaveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveService not implemented")
}
func (UnimplementedIShorturlServer) mustEmbedUnimplementedIShorturlServer() {}

// UnsafeIShorturlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IShorturlServer will
// result in compilation errors.
type UnsafeIShorturlServer interface {
	mustEmbedUnimplementedIShorturlServer()
}

func RegisterIShorturlServer(s grpc.ServiceRegistrar, srv IShorturlServer) {
	s.RegisterService(&IShorturl_ServiceDesc, srv)
}

func _IShorturl_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_Post_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_LongLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).LongLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_LongLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).LongLink(ctx, req.(*LongLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_ShortLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).ShortLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_ShortLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).ShortLink(ctx, req.(*ShortLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_UserAllLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAllLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).UserAllLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_UserAllLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).UserAllLink(ctx, req.(*UserAllLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_AllLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).AllLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_AllLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).AllLink(ctx, req.(*AllLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_AllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).AllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_AllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).AllUsers(ctx, req.(*AllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_UserDelLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDelLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).UserDelLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_UserDelLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).UserDelLink(ctx, req.(*UserDelLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_ReadService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).ReadService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_ReadService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).ReadService(ctx, req.(*ReadServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IShorturl_SaveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IShorturlServer).SaveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IShorturl_SaveService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IShorturlServer).SaveService(ctx, req.(*SaveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IShorturl_ServiceDesc is the grpc.ServiceDesc for IShorturl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IShorturl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.IShorturl",
	HandlerType: (*IShorturlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _IShorturl_Post_Handler,
		},
		{
			MethodName: "LongLink",
			Handler:    _IShorturl_LongLink_Handler,
		},
		{
			MethodName: "ShortLink",
			Handler:    _IShorturl_ShortLink_Handler,
		},
		{
			MethodName: "UserAllLink",
			Handler:    _IShorturl_UserAllLink_Handler,
		},
		{
			MethodName: "AllLink",
			Handler:    _IShorturl_AllLink_Handler,
		},
		{
			MethodName: "AllUsers",
			Handler:    _IShorturl_AllUsers_Handler,
		},
		{
			MethodName: "UserDelLink",
			Handler:    _IShorturl_UserDelLink_Handler,
		},
		{
			MethodName: "ReadService",
			Handler:    _IShorturl_ReadService_Handler,
		},
		{
			MethodName: "SaveService",
			Handler:    _IShorturl_SaveService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/shorturl.proto",
}
