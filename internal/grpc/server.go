package grpc

import (
	"context"
	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kz2d/telegram-bot/internal/grpc/api"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func Serve(port string) {
	grpcSever := grpc.NewServer()
	api.RegisterSomeServiceServer(grpcSever, &GreeterServerImpl{})
	mux := runtime.NewServeMux()

	err := api.RegisterSomeServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+port, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: withLogger(mux),
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	m := cmux.New(l)

	httpL := m.Match(cmux.HTTP1Fast())

	grpcL := m.Match(cmux.HTTP2())

	go server.Serve(httpL)
	go grpcSever.Serve(grpcL)

	m.Serve()
}

func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
