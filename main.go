package main

import (
	"log"

	"github.com/RajibDas-123/ms-grpc-auth/auth/cmd"
	"github.com/RajibDas-123/ms-grpc-auth/auth/logging"

	"github.com/subosito/gotenv"
)

func init() {
	if gotenv.Load(".env") != nil {
		log.Fatal("Failed to load the env file")
	}
	logging.Initialize()
}
func main() {
	cmd.Run()
}
