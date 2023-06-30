package main

import (
	"flag"
	"fmt"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"net"
	"strconv"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type GreeterServer struct{}

func (g *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("Receive from client: %v\n", resp)
		n++
		_ = stream.Send(&pb.HelloReply{Message: "message" + strconv.Itoa(n)})
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	err := server.Serve(lis)
	if err != nil {
		return
	}
}
