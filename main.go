package main

import (
	"context"
	"github.com/RejwankabirHamim/hello-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	severRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(severRegistrar, service)
	if err := severRegistrar.Serve(lis); err != nil {
		log.Fatalf("cannot serve: %s", err)
	}
}
