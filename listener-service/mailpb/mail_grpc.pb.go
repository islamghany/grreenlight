// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: mailpb/mail.proto

package mailpb

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

// MailSeviceClient is the client API for MailSevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailSeviceClient interface {
	SendMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error)
}

type mailSeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailSeviceClient(cc grpc.ClientConnInterface) MailSeviceClient {
	return &mailSeviceClient{cc}
}

func (c *mailSeviceClient) SendMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error) {
	out := new(MailResponse)
	err := c.cc.Invoke(ctx, "/mailpb.MailSevice/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailSeviceServer is the server API for MailSevice service.
// All implementations must embed UnimplementedMailSeviceServer
// for forward compatibility
type MailSeviceServer interface {
	SendMail(context.Context, *MailRequest) (*MailResponse, error)
	mustEmbedUnimplementedMailSeviceServer()
}

// UnimplementedMailSeviceServer must be embedded to have forward compatible implementations.
type UnimplementedMailSeviceServer struct {
}

func (UnimplementedMailSeviceServer) SendMail(context.Context, *MailRequest) (*MailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedMailSeviceServer) mustEmbedUnimplementedMailSeviceServer() {}

// UnsafeMailSeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailSeviceServer will
// result in compilation errors.
type UnsafeMailSeviceServer interface {
	mustEmbedUnimplementedMailSeviceServer()
}

func RegisterMailSeviceServer(s grpc.ServiceRegistrar, srv MailSeviceServer) {
	s.RegisterService(&MailSevice_ServiceDesc, srv)
}

func _MailSevice_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailSeviceServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailpb.MailSevice/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailSeviceServer).SendMail(ctx, req.(*MailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailSevice_ServiceDesc is the grpc.ServiceDesc for MailSevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailSevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mailpb.MailSevice",
	HandlerType: (*MailSeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _MailSevice_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailpb/mail.proto",
}
