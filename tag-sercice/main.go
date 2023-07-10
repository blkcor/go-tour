package main

import (
	"flag"
	"fmt"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var http_port string
var grpc_port string

func init() {
	flag.StringVar(&grpc_port, "grpc_port", "8001", "grpc服务启动端口号")
	flag.StringVar(&http_port, "http_port", "9001", "http服务启动端口号")
	flag.Parse()
}
func main() {
	errs := make(chan error)

	//rpc service
	go func() {
		err := RunRpcServer(grpc_port)
		if err != nil {
			errs <- err
			fmt.Printf("rpc service start err: %v", err)
		}
	}()

	//http service
	go func() {
		err := RunHttpServer(http_port)
		if err != nil {
			errs <- err
			fmt.Printf("http service start err: %v", err)
		}
	}()
	select {
	case err := <-errs:
		log.Fatalf("Run server err: %v", err)
	}
}

func RunHttpServer(port string) error {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})
	return http.ListenAndServe(":"+port, serverMux)
}

func RunRpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	//grpc service
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("net.Listen err: %v", err)
	}
	return s.Serve(lis)
}
