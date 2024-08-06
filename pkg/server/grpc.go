package server

import (
	"SAI/pkg/log"
	"SAI/pkg/service"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

type Server struct {
	gatewayRegister GatewayRegister
	register        GrpcRegister

	httpServer *http.Server
	server     *grpc.Server

	logger log.Logger
}

type GrpcRegister func(server *grpc.Server)
type GatewayRegister func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

func New(register GrpcRegister, gwRegister GatewayRegister) *Server {
	return &Server{
		gatewayRegister: gwRegister,
		register:        register,
		logger:          log.GetLogger(),
	}
}

const (
	contentTypeName  = "content-type"
	contentTypeValue = "application/grpc"
)

func (s *Server) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", service.Port))
	if err != nil {
		s.logger.Fatal(func() string {
			return "Failed to listen: " + err.Error()
		})
	}

	if s.gatewayRegister != nil {
		m := cmux.New(l)
		httpL := m.Match(cmux.HTTP1Fast())
		grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings(contentTypeName, contentTypeValue))

		go func() {
			err := s.httpRun(httpL)
			if err != nil {
				return
			}
		}()
		go func() {
			err := s.grpcRun(grpcL)
			if err != nil {
				panic(err)
			}
		}()

		return m.Serve()
	}

	return s.grpcRun(l)
}

func (s *Server) grpcRun(l net.Listener) error {
	var unaryInterceptors = []grpc.UnaryServerInterceptor{
		requestInterceptor(),
	}
	var streamInterceptors []grpc.StreamServerInterceptor

	s.server = grpc.NewServer(grpc.ChainUnaryInterceptor(unaryInterceptors...), grpc.ChainStreamInterceptor(streamInterceptors...))
	s.register(s.server)

	s.introduce()
	return s.server.Serve(l)
}

func (s *Server) httpRun(l net.Listener) error {
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	//swagger
	mux := http.NewServeMux()
	mux.Handle("/", gwMux)

	fs := http.FileServer(http.Dir("/public/swagger-ui"))
	mux.Handle("/docs/", http.StripPrefix("/docs", fs))

	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger.json")
	})

	err := s.gatewayRegister(context.Background(), gwMux, fmt.Sprintf("localhost:%d", service.Port), opts)
	if err != nil {
		return err
	}

	s.httpServer = &http.Server{Handler: mux}
	return s.httpServer.Serve(l)
}

func (s *Server) introduce() {
	s.logger.Info(func() string {
		return "Started server, listening on " + fmt.Sprintf(":%d", service.Port)
	}, log.NewField("port", fmt.Sprintf("%v", service.Port)))
}
