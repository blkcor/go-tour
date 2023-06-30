package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayRoute(client, &pb.HelloRequest{Name: []string{"rainbow", "blkcor"}})
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Printf("resp from server: %v\n", resp)
	}

	_ = stream.CloseSend()

	return nil
}
