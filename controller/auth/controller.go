package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/logging"
	"github.com/RajibDas-123/ms-grpc-auth/api-gateway/response"
	authservice "github.com/RajibDas-123/ms-grpc-auth/api-gateway/service/auth"
)

// Login : Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	responseCodes := make(map[string]string)
	file, _ := ioutil.ReadFile("../responseCode.json")
	_ = json.Unmarshal([]byte(file), &responseCodes)

	loginRequest := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		logging.AppLogger.Errorf("func : Login : Unable to decode to request body%v", err)
		response.BadRequest(w, "407", "Bad request body")
		return
	}

	sessionToken := r.Header.Get("session_token")
	// var data map[string]string //nolint:gosimple
	// data = make(map[string]string)
	sessionprocresp, err := authservice.IsSessionExist(sessionToken)
	success := sessionprocresp["success"]
	if err != nil {
		logging.AppLogger.Errorf("func: Login : Unable to get the session data", err)
		response.InternalServerError(w, "501", "Unable to fethch the session token")
	} else {
		if success == "false" {
			loginprocresp, err := authservice.Login(loginRequest["email"], loginRequest["password"])
			success = loginprocresp["success"]
			if err != nil {
				logging.AppLogger.Errorf("func: Login : Unable to login the user", err)
				response.InternalServerError(w, "502", "Unable to login")
			} else {
				if success == "true" {
					logging.AppLogger.Infof("func: Login : login success", err)
					response.Success(w, "201", "login success", nil)
				} else {
					logging.AppLogger.Infof("func: Login : login failed")
					response.StatusNotfound(w, "401", "Invalid credentials")
				}
			}

		} else {
			logging.AppLogger.Infof("func: Login : session exist", err)
			response.Success(w, "201", "login success", nil)
		}
	}

}
