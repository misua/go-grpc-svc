package main

import "log"

func Run() error {
	//TODO responsible for initializing and starting grpc
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}

}
