package main

import (
	"flag"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8003", "grpc服务启动端口号")
	flag.Parse()
}
func main() {
	lis, err := RunTcpServer(port)
	if err != nil {
		log.Fatalf("Run tcp server err: %v", err)
	}
	m := cmux.New(lis)
	//listener
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())
	//server
	grpcS := RunGrpcServer(port)
	httpS := RunHttpServer(port)

	go func() {
		_ = grpcS.Serve(grpcL)
	}()
	go func() {
		_ = httpS.Serve(httpL)
	}()

	//handle error
	err = m.Serve()
	if err != nil {
		log.Fatalf("Run cmux server err: %v", err)
	}
}

func RunHttpServer(port string) *http.Server {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})
	return &http.Server{
		Addr:    ":" + port,
		Handler: serverMux,
	}
}

func RunTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}
func RunGrpcServer(port string) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	return s
}
