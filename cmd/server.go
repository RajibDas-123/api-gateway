package cmd

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/controller"
)

func Execute() {
	router := controller.Initialize()
	if os.Getenv("BUILD") == "Prod" {
		log.Fatal(http.ListenAndServeTLS(
			os.Getenv("SERVER_ADDRESS"),
			filepath.Join("config", "tls.cert"),
			filepath.Join("config", "tls.key"),
			router,
		))
	} else {
		log.Fatal(http.ListenAndServe(
			os.Getenv("SERVER_ADDRESS"),
			router,
		))
	}
}
