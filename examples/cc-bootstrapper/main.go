package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"roshpr.net/example/cc-bootstrapper/bootstrap"
	"roshpr.net/example/cc-bootstrapper/controller"
)

const repo = "https://pkg-repo.zscalerlabs.com/O123"

type bsServer struct {
	bootstrap.UnimplementedBootstrapServer
}

func (s bsServer) Create(ctx context.Context, req *bootstrap.BootStrapConfig) (*bootstrap.BootStrapResponse, error) {
	return &bootstrap.BootStrapResponse{
		Type:  req.Type,
		Arch:  "amd64",
		Cores: 2,
	}, nil
}

type controllerServer struct {
	controller.UnimplementedControllerRequestServer
}

func (c controllerServer) Fetch(ctx context.Context, req *controller.ConnectorCtrlReq) (*controller.ConnectorCtrlResp, error) {
	log.Print("Received Controller request to Fetch service")
	return &controller.ConnectorCtrlResp{
		RepoServer: repo,
	}, nil
}
func main() {
	serverPort := ":8080"
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	server := grpc.NewServer()
	bsservice := &bsServer{}
	ctrlservice := &controllerServer{}
	log.Print("Register services")
	bootstrap.RegisterBootstrapServer(server, bsservice)
	controller.RegisterControllerRequestServer(server, ctrlservice)
	log.Print("Configure reflection")
	reflection.Register(server)
	log.Print("Listening on port " + serverPort)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve:", err)
	}
}
