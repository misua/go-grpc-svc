package main

import (
	"log"

	"github.com/misua/go-grpc-svc/internal/db"
	"github.com/misua/go-grpc-svc/internal/rocket"
)

func Run() error {
	//TODO responsible for initializing and starting grpc
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}

}
