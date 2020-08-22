package controller

import (
	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/controller/auth"

	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Initialize : Initialize function
func Initialize() *mux.Router {
	auth.Initialize()
	router := mux.NewRouter()
	router.Use(cors.AllowAll().Handler)

	router.Use(middleware.Auth)
	router.HandleFunc("/api/v1/login", auth.Login)
	return router
}
