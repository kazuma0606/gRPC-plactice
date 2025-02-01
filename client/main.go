package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	certFile := "C:/Users/yoshi/AppData/Local/mkcert/rootCA.pem"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	//mainの修了時にconnを閉じる
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	// callListFIles(client)
	callDownLoad(client)
	// callUpload(client)
	// CallUploadAndNotifyProgress(client)
}

func callListFIles(client pb.FileServiceClient) {
	md := metadata.New(map[string]string{"authorization": "Bearer test-token"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.ListFiles(ctx, &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}

	fmt.Println(res.GetFilenames())
}

func callDownLoad(client pb.FileServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.DownloadRequest{Filename: "name.txt"}
	strem, err := client.DownLoad(ctx, req)
	if err != nil {
		log.Fatalf("Failed to download: %v", err)
	}

	for {
		res, err := strem.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			resErr, ok := status.FromError(err)
			if ok {
				if resErr.Code() == codes.NotFound {
					log.Fatalf("Error code: %v,Error message: %v", resErr.Code(), resErr.Message())
				} else if resErr.Code() == codes.DeadlineExceeded {
					log.Fatalf("Deadline exceeded: %v", resErr.Message())
				} else {
					log.Fatalf("Unknown gRPC Error: %v", err)
				}
			} else {
				log.Fatalf("Failed to receive: %v", err)
			}
		}

		log.Printf("Response from Download(bytes): %v", res.GetData())
		log.Printf("Response from Download(bytes): %v", string(res.GetData()))
	}
}

func callUpload(client pb.FileServiceClient) {
	filename := "sports.txt"
	path := "C:/Users/yoshi/golang-plactice/grpc-lesson/storage/" + filename

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatalf("Failed to upload: %v", err)
	}

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}
		req := &pb.UploadRequest{Data: buf[:n]}
		sendErr := stream.Send(req)
		if sendErr != nil {
			log.Fatalf("Failed to send: %v", sendErr)
		}

		time.Sleep(1 * time.Second)

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	log.Printf("Response from Upload: %v", res.GetSize())
}

func CallUploadAndNotifyProgress(client pb.FileServiceClient) {
	filename := "sports.txt"
	path := "C:/Users/yoshi/golang-plactice/grpc-lesson/storage/" + filename

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stream, err := client.UploadAndNotifyProgress(context.Background())
	if err != nil {
		log.Fatalf("Failed to upload: %v", err)
	}

	//request
	buf := make([]byte, 5)
	go func() {
		for {
			n, err := file.Read(buf)
			if n == 0 || err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to read file: %v", err)
			}

			req := &pb.UploadAndNotifyProgressRequest{Data: buf[:n]}
			sendErr := stream.Send(req)
			if sendErr != nil {
				log.Fatalf("Failed to send: %v", sendErr)
			}
			time.Sleep(1 * time.Second)
		}

		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Failed to close send: %v", err)
		}
	}()

	//response
	ch := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive: %v", err)
			}

			log.Printf("Recieced message %v", res.GetMessage())
		}
		close(ch)
	}()
	<-ch
}
