package account

import "context"

// struct that will represent a user inside of out business logic
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Interface to be able to put our users into a DB
// These methods help us interface with our database - wehereas the methods in service will be the ones we expose from our microservice
type Repository interface {
	// Takes in a user struct
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}
