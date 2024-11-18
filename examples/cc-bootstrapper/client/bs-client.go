package main

import (
	"context"
	"crypto/tls"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	pb "roshpr.net/example/cc-bootstrapper/controller"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "https://repo.zscalerlabs.com:443", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	//conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	clientCert, err := tls.LoadX509KeyPair("/Users/roshpr/temp/cc_repo.crt", "/Users/roshpr/temp/cc_repo.key")
	if err != nil {
		log.Fatal(err)
	}
	// set config of tls credential
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
	}
	tlsCredential := credentials.NewTLS(config)
	conn, err := grpc.NewClient(
		"repo.zscalerlabs.com:443",
		grpc.WithTransportCredentials(tlsCredential),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)
	c := pb.NewControllerRequestClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Fetch(ctx, &pb.ConnectorCtrlReq{Type: "CCBC"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetRepoServer())
}
