package main

import (
	"context"
	pb "github.com/suboat/go-tensorflow-serving/tensorflow_serving/apis"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "127.0.0.1:8500"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	c := pb.NewModelServiceClient(conn)
	r, err := c.GetModelStatus(ctx, &pb.GetModelStatusRequest{
		ModelSpec: &pb.ModelSpec{
			Name: "some",
		},
	})
	if err != nil {
		panic(err)
	}
	//
	println(r.String())
}
