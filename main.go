package main

import (
	"context"
	"log"
	"net"

	"github.com/evandrolg/grpc-server/invoicer"
	"google.golang.org/grpc"
)

type invoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s invoicerServer) Create(ctx context.Context, request *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("PDF"),
		Docx: []byte("DOCX"),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalf("could not listen to :9090: %v", err)
	}

	serverRegister := grpc.NewServer()
	invoicer.RegisterInvoicerServer(serverRegister, &invoicerServer{})
	err = serverRegister.Serve(listen)

	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}
