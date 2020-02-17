package main

import (
	"context"
	"fmt"
	"github.com/go-acme/lego/log"
	micro "github.com/micro/go-micro"
	pb "github.com/ruandao/micro-shippy-vessel-service/proto/vessel"
	"os"
)

const (
	defaultDBUri = "datastore:27017"
)
func main() {

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
	)

	srv.Init()
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBUri
	}
	dbConn, err := CreateConnect(context.Background(), uri, 0)
	if err != nil {
		log.Fatalf("create database connection err: %v", err)
	}
	defer dbConn.Disconnect(context.Background())

	collection := dbConn.Database("shippy").Collection("vessel")
	repository := &MongoRepository{collection}
	h := &handler{repository}

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
