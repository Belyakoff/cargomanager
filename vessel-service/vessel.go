package main

import (
	"context"
	"os"
	"log"

	pb "github.com/Belyakoff/cargomanager/vessel-service/proto/vessel"
	"github.com/micro/go-micro/v2"
)


func main() {
	
	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("cargo").Collection("vessels")
	repository := &MongoRepository{vesselCollection}



	h := &handler{repository}

	if err := pb.RegisterVesselServiceHandler(srv.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}


}