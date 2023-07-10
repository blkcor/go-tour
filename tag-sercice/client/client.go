package main

import (
	"context"
	"fmt"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//create the background context
	ctx := context.Background()
	//get the client connection (block to wait for the connection established)
	clientConn, _ := GetClientCoon(ctx, "localhost:8004", []grpc.DialOption{grpc.WithBlock()})
	defer func() {
		err := clientConn.Close()
		if err != nil {
			fmt.Printf("the clientConnection fail to close")
			return
		}
	}()
	//get the tag service client
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	//call the service to get the response
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{
		Name:  "Go",
		State: 1,
	})
	log.Printf("resp is : %v", resp)
}

func GetClientCoon(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	//ban the secure examine
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(ctx, target, opts...)
}
