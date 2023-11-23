// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/file.proto

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

// FileUploadClient is the client API for FileUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileUploadClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileUpload_UploadClient, error)
}

type fileUploadClient struct {
	cc grpc.ClientConnInterface
}

func NewFileUploadClient(cc grpc.ClientConnInterface) FileUploadClient {
	return &fileUploadClient{cc}
}

func (c *fileUploadClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileUpload_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileUpload_ServiceDesc.Streams[0], "/file_upload.FileUpload/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadUploadClient{stream}
	return x, nil
}

type FileUpload_UploadClient interface {
	Send(*FileChunks) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type fileUploadUploadClient struct {
	grpc.ClientStream
}

func (x *fileUploadUploadClient) Send(m *FileChunks) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileUploadUploadClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileUploadServer is the server API for FileUpload service.
// All implementations must embed UnimplementedFileUploadServer
// for forward compatibility
type FileUploadServer interface {
	Upload(FileUpload_UploadServer) error
	mustEmbedUnimplementedFileUploadServer()
}

// UnimplementedFileUploadServer must be embedded to have forward compatible implementations.
type UnimplementedFileUploadServer struct {
}

func (UnimplementedFileUploadServer) Upload(FileUpload_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileUploadServer) mustEmbedUnimplementedFileUploadServer() {}

// UnsafeFileUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileUploadServer will
// result in compilation errors.
type UnsafeFileUploadServer interface {
	mustEmbedUnimplementedFileUploadServer()
}

func RegisterFileUploadServer(s grpc.ServiceRegistrar, srv FileUploadServer) {
	s.RegisterService(&FileUpload_ServiceDesc, srv)
}

func _FileUpload_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileUploadServer).Upload(&fileUploadUploadServer{stream})
}

type FileUpload_UploadServer interface {
	SendAndClose(*Status) error
	Recv() (*FileChunks, error)
	grpc.ServerStream
}

type fileUploadUploadServer struct {
	grpc.ServerStream
}

func (x *fileUploadUploadServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileUploadUploadServer) Recv() (*FileChunks, error) {
	m := new(FileChunks)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileUpload_ServiceDesc is the grpc.ServiceDesc for FileUpload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileUpload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file_upload.FileUpload",
	HandlerType: (*FileUploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileUpload_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/file.proto",
}