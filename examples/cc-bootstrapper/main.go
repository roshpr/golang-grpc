package main

import (
	"context"
	"google.golang.org/grpc"
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
	return &controller.ConnectorCtrlResp{
		RepoServer: repo,
	}, nil
}
func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	server := grpc.NewServer()
	bsservice := &bsServer{}
	ctrlservice := &controllerServer{}
	bootstrap.RegisterBootstrapServer(server, bsservice)
	controller.RegisterControllerRequestServer(server, ctrlservice)

	err = server.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve:", err)
	}
}
