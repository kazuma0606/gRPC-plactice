// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: proto/file.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileService_ListFiles_FullMethodName               = "/file.FileService/ListFiles"
	FileService_DownLoad_FullMethodName                = "/file.FileService/DownLoad"
	FileService_Upload_FullMethodName                  = "/file.FileService/Upload"
	FileService_UploadAndNotifyProgress_FullMethodName = "/file.FileService/UploadAndNotifyProgress"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	DownLoad(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadResponse], error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error)
	UploadAndNotifyProgress(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse], error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, FileService_ListFiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DownLoad(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], FileService_DownLoad_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DownloadRequest, DownloadResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_DownLoadClient = grpc.ServerStreamingClient[DownloadResponse]

func (c *fileServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[1], FileService_Upload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadRequest, UploadResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadClient = grpc.ClientStreamingClient[UploadRequest, UploadResponse]

func (c *fileServiceClient) UploadAndNotifyProgress(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[2], FileService_UploadAndNotifyProgress_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadAndNotifyProgressClient = grpc.BidiStreamingClient[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse]

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
type FileServiceServer interface {
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	DownLoad(*DownloadRequest, grpc.ServerStreamingServer[DownloadResponse]) error
	Upload(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error
	UploadAndNotifyProgress(grpc.BidiStreamingServer[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse]) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileServiceServer) DownLoad(*DownloadRequest, grpc.ServerStreamingServer[DownloadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method DownLoad not implemented")
}
func (UnimplementedFileServiceServer) Upload(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileServiceServer) UploadAndNotifyProgress(grpc.BidiStreamingServer[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadAndNotifyProgress not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_ListFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DownLoad_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownLoad(m, &grpc.GenericServerStream[DownloadRequest, DownloadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_DownLoadServer = grpc.ServerStreamingServer[DownloadResponse]

func _FileService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).Upload(&grpc.GenericServerStream[UploadRequest, UploadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadServer = grpc.ClientStreamingServer[UploadRequest, UploadResponse]

func _FileService_UploadAndNotifyProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadAndNotifyProgress(&grpc.GenericServerStream[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_UploadAndNotifyProgressServer = grpc.BidiStreamingServer[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse]

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FileService_ListFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownLoad",
			Handler:       _FileService_DownLoad_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upload",
			Handler:       _FileService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadAndNotifyProgress",
			Handler:       _FileService_UploadAndNotifyProgress_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/file.proto",
}
