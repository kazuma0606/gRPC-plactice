package main

import (
	"bytes"
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"
	"net"
	"os"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// path: "C:/Users/yoshi/golang-plactice/grpc-lesson/storage/"

type server struct {
	pb.UnimplementedFileServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println(("ListFiles function was invoked"))

	dir := "C:/Users/yoshi/golang-plactice/grpc-lesson/storage"

	paths, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}

	return res, nil
}

func (*server) DownLoad(req *pb.DownloadRequest, stream pb.FileService_DownLoadServer) error {
	fmt.Println("Download function was invoked")

	filename := req.GetFilename()
	path := "C:/Users/yoshi/golang-plactice/grpc-lesson/storage/" + filename

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return status.Error(codes.NotFound, "File not found")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		res := &pb.DownloadResponse{Data: buf[:n]}
		sendErr := stream.Send(res)
		if sendErr != nil {
			return sendErr
		}

		time.Sleep((1 * time.Second)) // 1秒待つ
	}

	return nil
}

func (*server) Upload(stream pb.FileService_UploadServer) error {
	fmt.Println("Upload function was invoked")

	var buf bytes.Buffer
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &pb.UploadResponse{Size: int32(buf.Len())}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}

		data := req.GetData()
		log.Printf("Received data size:(bytes) %v", data)
		log.Printf("Received data size:(string) %v", string(data))
		buf.Write(data)
	}
}

func (*server) UploadAndNotifyProgress(stream pb.FileService_UploadAndNotifyProgressServer) error {
	fmt.Println("UploadAndNotifyProgress function was invoked")

	size := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		data := req.GetData()
		log.Printf("Received data size:(bytes) %v", data)
		size += len(data)

		res := &pb.UploadAndNotifyProgressResponse{
			Message: fmt.Sprintf("Received data size: %v bytes", size),
		}
		err = stream.Send(res)
		if err != nil {
			return err
		}
	}

}

func myLogging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		log.Printf("Request data: %v", req)

		resp, err = handler(ctx, req)
		if err != nil {
			return nil, err
		}
		log.Printf("Response data: %v", resp)

		return resp, nil
	}
}

func authorize(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}

	// トークンの検証
	if token != "test-token" {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	return ctx, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(
		"ssl/localhost.pem",
		"ssl/localhost-key.pem")

	if err != nil {
		log.Fatalf("Failed to generate credentials: %v", err)
	}

	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				myLogging(),
				grpc_auth.UnaryServerInterceptor(authorize))))
	pb.RegisterFileServiceServer(s, &server{})

	fmt.Println("Server is running on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
