package auth

import (
	"log"

	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/service/auth/pb"

	"google.golang.org/grpc"
)

// AuthClient : be used to connect auth service
var AuthClient pb.AuthenticationClient

// Initialize : Initialize function
func Initialize() {
	Conn, err := grpc.Dial("auth-service:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	} else {
		log.Println("Established connection.", Conn.GetState())
	}
	AuthClient = pb.NewAuthenticationClient(Conn)
}
