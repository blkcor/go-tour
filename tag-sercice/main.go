package main

import (
	"flag"
	"fmt"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8001", "启动端口号")
	flag.Parse()
}
func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("net.Listen err: %v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("serve.Serve err: %v", err)
	}
}
