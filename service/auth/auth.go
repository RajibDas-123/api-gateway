package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/logging"
	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/service/auth/pb"
)

// Login : Calls the client
func Login(id string, pass string) (responseData map[string]string, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	request := &pb.LoginRequest{
		EmailId:  id,
		Password: pass,
	}
	if response, err := AuthClient.Login(ctx, request); err == nil {
		userJSON, err := json.Marshal(response)
		json.Unmarshal([]byte(userJSON), &responseData)
		return responseData, err
	}
	return responseData, err
}

// IsSessionExist : Calls the client
func IsSessionExist(session string) (responseData map[string]string, err error) {
	logging.CacheLogger.Infof("func:IsSessionExist: Checking whether Session token is in session or not ", session)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	request := &pb.IsSessionExistRequest{
		SessionToken: session,
	}
	if response, err := AuthClient.IsSessionExist(ctx, request); err == nil {
		responseJSON, err := json.Marshal(response)
		json.Unmarshal([]byte(responseJSON), &responseData)
		return responseData, err
	}
	return responseData, err
}
