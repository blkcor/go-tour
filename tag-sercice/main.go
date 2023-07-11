package main

import (
	"context"
	"encoding/json"
	"flag"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strings"
)

var port string

type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func init() {
	flag.StringVar(&port, "port", "8004", "grpc服务启动端口号")
	flag.Parse()
}
func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Server err: %v", err)
	}
}

func RunServer(port string) error {
	httpMux := RunHttpServer()
	grpcS := RunGrpcServer()
	gatewayMux := RunGrpcGatewayServer()
	httpMux.Handle("/", gatewayMux)
	httpS := &http.Server{
		Addr:    ":" + port,
		Handler: GrpcHandlerFunc(grpcS, httpMux),
	}
	return httpS.ListenAndServe()
}

func RunHttpServer() *http.ServeMux {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})
	return serverMux
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	return s
}

func RunGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	//registry grpcGatewayError as http error
	runtime.HTTPError = grpcGatewayError
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)
	return gwmux
}

func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//判断是否为grpc请求
		if request.ProtoMajor == 2 && strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(writer, request)
		} else {
			otherHandler.ServeHTTP(writer, request)
		}
	}), &http2.Server{})
}

func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	//from error get the status
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}
	//get http error
	httpError := httpError{
		Code:    int32(s.Code()),
		Message: s.Message(),
	}
	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}
	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-Type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}
