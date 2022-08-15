package account

import "context"

// Interface with the methods you want to expose to the transports and use that interface to implement your business logic
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

// To implement the transports (HTTP etc) we need to takes our methods and turn them into endpoints
// To do this we need to create structs that represent the response and the requests from each of our method
