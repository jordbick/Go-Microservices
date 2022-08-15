package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct contains the methods we want to expose to the public
// Convert the methods into this endpoint.Endpoint type = A function that takes in a context and an interface and returns an interface and an error
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

// Function which will allow us to create this endpoint struct
// Need to pass in the Service interface. Will pass it into two different functions which will take the methods and convert them into these endpoints
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

// Take in ther service and return an Endpoint which is a function
func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	// use anonymous functions as wrapper for our existing methods
	// cast tthe request interface{} as a CreateUserRequest struct (in reqresp.go file)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		// then call service.CreateUser on our context and the request email and password, which will return a success string or error
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		// can pass this into our CreateUserResponse struct
		return CreateUserResponse{Ok: ok}, err

	}
}
func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// Take the request interface{} and cast it as a GetUserRequest struct
		req := request.(GetUserRequest)
		// Call service.GetUser on the context and request ID
		email, err := s.GetUser(ctx, req.ID)

		return GetUserResponse{Email: email}, err
	}
}
