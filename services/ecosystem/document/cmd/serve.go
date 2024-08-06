package cmd

import (
	"SAI/pkg/server"
	"SAI/proto/generated/ecosystem/document/document_pb"
	"SAI/services/ecosystem/document/handler"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Document Service",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	service := handler.New()

	grpcService := func(server *grpc.Server) {
		document_pb.RegisterServiceServer(server, service)
	}

	s := server.New(grpcService, document_pb.RegisterServiceHandlerFromEndpoint)
	s.Run()
}
