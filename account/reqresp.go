package account

import (
	"context"
	"encoding/json"
	"net/http"
)

// structs that represent the response and the requests from each of our method
// All structs and fields need to be Public
type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		ID string `json:"id"`
	}

	GetUserResponse struct {
		Email string `json:"email"`
	}
)

// Take in context, ResponseWriter and the response interface and output an error
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	// Decode the request into the variable type CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	// Use gorilla mux.Vars helper function to get back a map that will contain all of the variables isnide of our path

	vars := mux.Vars(r)

	req = GetUserRequest{
		// so we can get the variable by just passing in the key "id" and put it into our GetUserRequest struct
		ID: vars["id"],
	}
	return req, nil
}
