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

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("failed to run migrations")
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
