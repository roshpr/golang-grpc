package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"roshpr.net/example/cc-bootstrapper/bootstrap"
	"roshpr.net/example/cc-bootstrapper/controller"
	"time"
)

const repo = "https://pkg-repo.zscalerlabs.com/O123"

var (
	sleep  = flag.Duration("sleep", time.Second*30, "duration between changes in health")
	system = "" // empty string represents the health of the system
)

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
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(server, healthcheck)
	bsservice := &bsServer{}
	ctrlservice := &controllerServer{}
	log.Print("Register services")
	bootstrap.RegisterBootstrapServer(server, bsservice)
	controller.RegisterControllerRequestServer(server, ctrlservice)
	log.Print("Configure health check")
	go func() {
		// asynchronously inspect dependencies and toggle serving status as needed
		next := healthpb.HealthCheckResponse_SERVING
		for {
			healthcheck.SetServingStatus(system, next)
			log.Print("Serving health check")
			if next == healthpb.HealthCheckResponse_SERVING {
				next = healthpb.HealthCheckResponse_NOT_SERVING
			} else {
				next = healthpb.HealthCheckResponse_SERVING
			}
			time.Sleep(*sleep)
		}
	}()
	log.Print("Configure reflection")
	reflection.Register(server)
	log.Print("Listening on port " + serverPort)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve:", err)
	}
}
