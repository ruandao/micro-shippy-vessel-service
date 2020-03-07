package main

import (
	"context"
	"fmt"
	"github.com/go-acme/lego/log"
	micro "github.com/micro/go-micro"
	"github.com/ruandao/micro-shippy-vessel-service/lib"
	pb "github.com/ruandao/micro-shippy-vessel-service/proto/vessel"
	"os"
)

const (
	defaultDBUri = "datastore:27017"
)
func main() {

	srv := micro.NewService(
		micro.Name(lib.CONST_SERVICE_NAME),
	)

	srv.Init()
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBUri
	}
	dbConn, err := lib.CreateConnect(context.Background(), uri, 0)
	if err != nil {
		log.Fatalf("create database connection err: %v", err)
	}
	defer dbConn.Disconnect(context.Background())

	collection := dbConn.Database("shippy").Collection("vessel")
	repository := &lib.MongoRepository{Collection:collection}
	h := &lib.Handler{Repository: repository}

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
